package dao

import (
	"Golang/2025/01January/Shopping/model"
	"database/sql"
	"errors"
	"log"
)

func ShopExist(shop model.Shop) bool {
	var Exist bool
	query := `select 1 from shop where shop_name = ?`
	err = db.QueryRow(query, shop.Shop_name).Scan(&Exist)
	if err != nil {
		return false
	}
	return Exist
}

func GetShopId(shop_name string) string {
	var Id string
	query := `SELECT id FROM shop WHERE shop_name = ?`
	err := db.QueryRow(query, shop_name).Scan(&Id)
	if err != nil {
		log.Println(err)
		return ""
	}
	return Id
}

func RegisterMall(shop model.Shop) bool {
	query := `insert into shop (shop_name, password) values(?, ?)`
	_, err := db.Exec(query, shop.Shop_name, shop.Password)
	if err != nil {
		log.Println(err)
		return false
	}
	return true
}

func LoginMall(shop model.Shop) bool {
	var exist bool
	query := `select 1 from shop where shop_name = ? and password = ?`
	err := db.QueryRow(query, shop.Shop_name, shop.Password).Scan(&exist)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return false
		}
		log.Println(err)
		return false
	}
	return exist
}

func RegisterGoods(goods model.Goods) bool {
	query := `insert into goods (goods_name, shop_id, type, number, price, content, avatar) values (?, ?, ?, ?, ?, ?, ?)`
	_, err := db.Exec(query, goods.Goods_name, goods.Shop_id, goods.Type, goods.Number, goods.Price, goods.Content, goods.Avatar)
	if err != nil {
		log.Println(err)
		return false
	}
	return true
}

func AlterGoodsInfo(goods model.Goods) bool {
	if goods.Goods_name != "" {
		query := `UPDATE goods SET goods_name = ? WHERE id = ? AND shop_id = ?`
		res, err := db.Exec(query, goods.Goods_name, goods.Id, goods.Shop_id)
		if err != nil {
			log.Println(err)
			return false
		}
		aff, err0 := res.RowsAffected()
		if err0 != nil || aff == 0 {
			log.Println(err0)
			return false
		}
	}
	if goods.Content != "" {
		query := `UPDATE goods SET content = ? WHERE id = ? AND shop_id = ?`
		res, err := db.Exec(query, goods.Content, goods.Id, goods.Shop_id)
		if err != nil {
			log.Println(err)
			return false
		}
		aff, err0 := res.RowsAffected()
		if err0 != nil || aff == 0 {
			log.Println(err0)
			return false
		}
	}
	if goods.Price != 0.0 {
		query := `update goods set price = ? where shop_id = ? AND id = ?`
		res, err := db.Exec(query, goods.Price, goods.Shop_id, goods.Id)
		if err != nil {
			log.Println(err)
			return false
		}
		aff, err0 := res.RowsAffected()
		if err0 != nil || aff == 0 {
			log.Println(err0)
			return false
		}
	}
	if goods.Number != 0 {
		query := `update goods set number = ? where shop_id = ? AND id = ?`
		res, err := db.Exec(query, goods.Number, goods.Shop_id, goods.Id)
		if err != nil {
			log.Println(err)
			return false
		}
		aff, err0 := res.RowsAffected()
		if err0 != nil || aff == 0 {
			log.Println(err0)
			return false
		}
	}

	if goods.Avatar != "" {
		query := `UPDATE goods SET avatar = ? WHERE shop_id = ? AND id = ?`
		res, err := db.Exec(query, goods.Avatar, goods.Shop_id, goods.Id)
		if err != nil {
			log.Println(err)
			return false
		}
		aff, err0 := res.RowsAffected()
		if err0 != nil || aff == 0 {
			log.Println(err0)
			return false
		}
	}
	return true
}

func DeleteGoods(goods model.Goods) bool {
	query := `delete from goods where id = ? AND shop_id = ?`
	res, err := db.Exec(query, goods.Id, goods.Shop_id)
	if err != nil {
		log.Println(err)
		return false
	}
	aff, err0 := res.RowsAffected()
	if err0 != nil {
		log.Println(err0)
		return false
	}
	if aff == 0 {
		return false
	}
	return true
}
