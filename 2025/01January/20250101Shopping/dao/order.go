package dao

import (
	"Golang/2025/01January/20250101Shopping/model"
	"log"
	"strconv"
)

// SubmitOrder 提交订单，清空购物车
func SubmitOrder(order *model.Order) bool {
	//初步向数据库注册订单
	query := `SELECT sum FROM shopping_cart WHERE user_id = ?`
	err := db.QueryRow(query, order.User_id).Scan(&order.Sum)
	if err != nil || order.Sum == 0 {
		log.Println(err)
		return false
	}
	tx, _ := db.Begin()
	query = `INSERT INTO orders (user_id, sum) values (?, ?)`
	IdGet, err0 := tx.Exec(query, order.User_id, order.Sum)
	if err0 != nil {
		err := tx.Rollback()
		if err != nil {
			return false
		}
		log.Println(err0)
		return false
	}
	id, err1 := IdGet.LastInsertId()
	if err1 != nil {
		log.Println(err1)
		return false
	}
	order.Id = strconv.FormatInt(id, 10)
	//给订单添加商品
	query = `SELECT goods_id, goods_name, number, price FROM cart_goods WHERE user_id = ?`
	Rows, err2 := db.Query(query, order.User_id)
	if err2 != nil {
		err := tx.Rollback()
		if err != nil {
			return false
		}
		log.Println(err2)
		return false
	}
	for Rows.Next() {
		var goods model.Order_Goods
		goods.Order_id = order.Id
		err = Rows.Scan(&goods.Goods_Id, &goods.Goods_Name, &goods.Number, &goods.Price)
		if err != nil {
			log.Println(err)
			return false
		}
		query = `INSERT INTO order_goods (id, goods_name, number, price, order_id) values (?, ?, ?, ?, ?)`
		_, err = tx.Exec(query, goods.Goods_Id, goods.Goods_Name, goods.Number, goods.Price, order.Id)
		if err != nil {
			err := tx.Rollback()
			if err != nil {
				return false
			}
			log.Println(err)
			return false
		}
		order.Goods = append(order.Goods, goods)
	}
	err = tx.Commit()
	if err != nil {
		return false
	}
	return true
}

// ConfirmOrder 确认订单，直接缴费
func ConfirmOrder(order model.Order) ([]model.Lack_Msg, string) {
	query := `SELECT is_deleted FROM orders WHERE id = ? AND user_id = ?`
	var st int
	err := db.QueryRow(query, order.Id, order.User_id).Scan(&st)
	if err != nil {
		log.Println(err)
		return nil, "error"
	}
	if st == 1 {
		return nil, "deleted"
	}
	//查看用户余额是否足够
	query = `SELECT sum FROM orders WHERE id = ?`
	var sum float64
	err = db.QueryRow(query, order.Id).Scan(&sum)
	if err != nil {
		log.Println(err)
		return nil, "error"
	}
	query = `SELECT balance FROM user WHERE id = ?`
	var balance float64
	err = db.QueryRow(query, order.User_id).Scan(&balance)
	if balance < sum {
		return nil, "lack"
	}
	//检查商品库存
	query = `SELECT id, number FROM order_goods WHERE order_id = ?`
	Rows, err0 := db.Query(query, order.Id)
	if err0 != nil {
		log.Println(err0)
		return nil, "error"
	}
	var LackGoods []model.Lack_Msg
	for Rows.Next() {
		var id string
		var num int
		err = Rows.Scan(&id, &num)
		if err != nil {
			log.Println(err)
			return nil, "error"
		}
		query = `SELECT number FROM goods WHERE id = ?`
		var shop_num int
		err = db.QueryRow(query, id).Scan(&shop_num)
		if shop_num < num {
			LackGoods = append(LackGoods, model.Lack_Msg{
				Goods_id:    id,
				Current_Num: shop_num,
				Query_Num:   num,
			})
		}
	}
	//如果有缺货的商品，返回缺货商品组成的字符串
	if LackGoods != nil {
		return LackGoods, "LackGoods"
	}

	//最后开始执行交易
	//此处减少商店库存
	query = `SELECT id, number, price FROM order_goods WHERE order_id = ?`
	Rows, err = db.Query(query, order.Id)
	if err != nil {
		log.Println(err)
		return nil, "error"
	}
	tx, _ := db.Begin()
	for Rows.Next() {
		var id int
		var num int
		var price float64
		err = Rows.Scan(&id, &num, &price)
		if err != nil {
			log.Println(err)
			return nil, "error"
		}
		query = `UPDATE goods SET number = number - ? WHERE id = ?`
		_, err = tx.Exec(query, num, id)
		if err != nil {
			err := tx.Rollback()
			if err != nil {
				return nil, ""
			}
			log.Println(err)
			return nil, "error"
		}
		query = `SELECT shop_id FROM goods WHERE id = ?`
		var shop_id int
		err = db.QueryRow(query, id).Scan(&shop_id)
		if err != nil {
			err := tx.Rollback()
			if err != nil {
				return nil, ""
			}
			log.Println(err)
			return nil, "error"
		}
		all := float64(num) * price
		query = `UPDATE shop SET profit = profit + ? WHERE id = ?`
		_, err = tx.Exec(query, all, shop_id)
		if err != nil {
			err := tx.Rollback()
			if err != nil {
				return nil, ""
			}
			log.Println(err)
			return nil, "error"
		}
	}
	//清空购物车
	query = `UPDATE shopping_cart SET sum = 0 WHERE user_id = ?`
	_, err = tx.Exec(query, order.User_id)
	if err != nil {
		err := tx.Rollback()
		if err != nil {
			return nil, ""
		}
		log.Println(err)
		return nil, "error"
	}
	query = `DELETE FROM cart_goods WHERE user_id = ?`
	_, err = tx.Exec(query, order.User_id)
	if err != nil {
		err := tx.Rollback()
		if err != nil {
			return nil, ""
		}
		log.Println(err)
		return nil, "error"
	}
	query = `UPDATE user SET balance = balance - ? WHERE id = ?`
	_, err = tx.Exec(query, sum, order.User_id)
	if err != nil {
		err := tx.Rollback()
		if err != nil {
			return nil, ""
		}
		log.Println(err)
		return nil, "error"
	}
	query = `UPDATE orders SET is_deleted = 1 WHERE id = ?`
	_, err = tx.Exec(query, order.Id)
	if err != nil {
		err := tx.Rollback()
		if err != nil {
			return nil, ""
		}
		log.Println(err)
		return nil, "error"
	}
	err = tx.Commit()
	if err != nil {
		return nil, ""
	}
	return nil, "ok"
}

// CancelOrder 取消订单，设置delete字段，数据库中不删除，只是无法再操作该订单。
func CancelOrder(order model.Order) bool {
	query := `UPDATE orders SET is_deleted = 1 WHERE id = ? AND user_id = ?`
	_, err := db.Exec(query, order.Id, order.User_id)
	if err != nil {
		log.Println(err)
		return false
	}
	return true
}
