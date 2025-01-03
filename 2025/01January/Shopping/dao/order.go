package dao

import (
	"Golang/2025/01January/Shopping/model"
	"log"
)

// SubmitOrder 提交订单，清空购物车
func SubmitOrder(order *model.Order) bool {
	//初步向数据库注册订单
	query := `SELECT sum FROM shopping_cart WHERE user_id = ?`
	err := db.QueryRow(query, order.User_id).Scan(&order.Sum)
	if err != nil {
		log.Println(err)
		return false
	}
	query = `INSERT INTO orders (user_id, sum) values (?, ?)`
	IdGet, err0 := db.Exec(query, order.User_id, order.Sum)
	if err0 != nil {
		log.Println(err0)
		return false
	}
	id, err1 := IdGet.LastInsertId()
	if err1 != nil {
		log.Println(err1)
		return false
	}
	order.Id = int(id)
	//给订单添加商品
	query = `SELECT goods_id, number, price FROM cart_goods WHERE user_id = ?`
	Rows, err2 := db.Query(query, order.User_id)
	if err2 != nil {
		log.Println(err2)
		return false
	}
	for Rows.Next() {
		var goods model.Order_Goods
		goods.Order_id = order.Id
		err = Rows.Scan(&goods.Goods_Id, &goods.Number, &goods.Price)
		if err != nil {
			log.Println(err)
			return false
		}
		query = `INSERT INTO order_goods (id, price, number, order_id) values (?, ?, ?, ?)`
		_, err = db.Exec(query, &goods.Goods_Id, goods.Price, goods.Number, order.Id)
		if err != nil {
			log.Println(err)
			return false
		}
		order.Goods = append(order.Goods, goods)
	}
	return true
}

func ConfirmOrder(order model.Order) string {
	query := `SELECT is_deleted FROM orders WHERE id = ?`
	var st int
	err := db.QueryRow(query, order.Id).Scan(&st)
	if err != nil {
		log.Println(err)
		return "error"
	}
	if st == 1 {
		return "deleted"
	}
	query = `UPDATE shopping_cart SET sum = 0 WHERE user_id = ?`
	_, err = db.Exec(query, order.User_id)
	if err != nil {
		log.Println(err)
		return "error"
	}
	query = `DELETE FROM cart_goods WHERE user_id = ?`
	_, err = db.Exec(query, order.User_id)
	if err != nil {
		log.Println(err)
		return "error"
	}
	query = `SELECT sum FROM orders WHERE id = ?`
	var sum float64
	err = db.QueryRow(query, order.Id).Scan(&sum)
	if err != nil {
		log.Println(err)
		return "error"
	}
	query = `SELECT balance FROM user WHERE id = ?`
	var balance float64
	err = db.QueryRow(query, order.User_id).Scan(&balance)
	if balance < sum {
		return "lack"
	}
	query = `UPDATE user SET balance = balance - ? WHERE id = ?`
	_, err = db.Exec(query, sum, order.User_id)
	if err != nil {
		log.Println(err)
		return "error"
	}
	return "ok"
}

func CancelOrder(order model.Order) bool {
	query := `UPDATE orders SET is_deleted = 1 WHERE id = ?`
	_, err := db.Exec(query, order.Id)
	if err != nil {
		log.Println(err)
		return false
	}
	return true
}
