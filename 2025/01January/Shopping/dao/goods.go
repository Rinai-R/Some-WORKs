package dao

import (
	"Golang/2025/01January/Shopping/model"
	"log"
)

func GetGoodInfo(goods_id int) (int, float64) {
	query := `select number, price from goods where id = ?`
	var num int
	var price float64
	err := db.QueryRow(query, goods_id).Scan(&num, &price)
	if err != nil {
		log.Println(err)
		return 0, 0.0
	}
	return num, price

}

// AddGoods 向购物车中添加商品
func AddGoods(user model.User, goods model.Goods) bool {
	id := GetId(user.Username)
	num, price := GetGoodInfo(goods.Id)
	if num < goods.Number {
		return false
	}
	query := `insert into cart_goods (user_id, goods_id, number, price) values (?, ?, ?, ?)`
	_, err1 := db.Exec(query, id, goods.Id, goods.Number, price)
	if err1 != nil {
		log.Println(err1)
		return false
	}
	query = `update shopping_cart set sum = sum + ? where user_id = ?`
	sum := float64(goods.Number) * price
	_, err2 := db.Exec(query, sum, user.Username)
	if err2 != nil {
		log.Println(err2)
		return false
	}
	return true
}

func GetCartGoodsInfo(cart_goods model.Cart_Goods) (int, float64) {
	query := `SELECT number, price FROM cart_goods WHERE user_id = ? AND goods_id = ?`
	var num int
	var price float64
	err := db.QueryRow(query, cart_goods.User_Id, cart_goods.Goods_Id).Scan(&num, &price)
	if err != nil {
		log.Println(err)
		return 0, 0.0
	}
	return num, price
}

func DelCartGoods(cart_goods model.Cart_Goods) bool {
	num, price := GetCartGoodsInfo(cart_goods)
	total := float64(num) * price
	query := `UPDATE shopping_cart SET sum = sum - ? WHERE user_id = ?`
	_, err := db.Exec(query, total)
	if err != nil {
		log.Println(err)
		return false
	}
	query = `DELETE FROM cart_goods WHERE user_id = ? AND goods_id = ?`
	_, err = db.Exec(query, cart_goods.User_Id, cart_goods.Goods_Id)
	if err != nil {
		log.Println(err)
		return false
	}
	return true
}

func BrowseGoods(goods *model.Goods, Browse model.Browse) bool {
	query0 := `SELECT id, goods_name, shop_id, type, number, price, star, content, avatar FROM goods WHERE id = ?`
	err0 := db.QueryRow(query0, Browse.Goods_id).Scan(&goods.Id, &goods.Goods_name, &goods.Shop_id, &goods.Type, &goods.Number, &goods.Price, &goods.Star, &goods.Content, &goods.Avatar)
	if err0 != nil {
		log.Println(err0)
		return false
	}
	query := `INSERT INTO browse_records (user_id, goods_id) values(?, ?) `
	_, err := db.Exec(query, Browse.User_id, Browse.Goods_id)
	if err != nil {
		log.Println(err)
		return false
	}
	return true
}

func StarGoods(star model.Star) bool {
	query := `INSERT INTO star (user_id, goods_id) values(?, ?)`
	_, err := db.Exec(query, star.User_id, star.Goods_id)
	if err != nil {
		log.Println(err)
		return false
	}
	return true
}

func SearchTypeGoods(goods *model.Goods) bool {
	query := `SELECT id, goods_name, shop_id, type, number, price, star FROM goods WHERE type = ? ORDER BY star DESC `
	err := db.QueryRow(query, goods.Type).Scan(&goods.Id, &goods.Goods_name, &goods.Type, &goods.Number, &goods.Price, &goods.Price, &goods.Star)
	if err != nil {
		log.Println(err)
		return false
	}
	return true
}

func SearchGoods(search model.Search) []model.Association {
	query := `INSERT INTO search (content) values (?)`
	IdGet, err := db.Exec(query, search.Content)
	if err != nil {
		log.Println(err)
		return nil
	}
	Id, err1 := IdGet.LastInsertId()
	if err1 != nil {
		log.Println(err1)
		return nil
	}
	search.Id = int(Id)
	return AssociationCount(search)
}

func AssociationCount(search model.Search) []model.Association {
	query := `SELECT id, goods_name, content, avatar FROM goods `
	rows, err := db.Query(query)
	if err != nil {
		log.Println(err)
		return nil
	}
	var ans []model.Association
	for rows.Next() {
		var tmp model.Association
		tmp.Search_id = search.Id
		var goods_content string
		err = rows.Scan(&tmp.Goods_id, &tmp.Goods_name, &goods_content, &tmp.Avatar)
		if err != nil {
			log.Println(err)
			return nil
		}
		tmp.Value = ComPare(search.Content, goods_content) + ComPare(search.Content, tmp.Goods_name)
		query = `INSERT INTO association (search_id, goods_id, value, goods_name, avatar) values(?, ?, ?, ?, ?)`
		_, err = db.Exec(query, tmp.Search_id, tmp.Goods_id, tmp.Value, tmp.Goods_name, tmp.Avatar)
		if err != nil {
			log.Println(err)
			return nil
		}
	}
	query = `SELECT search_id, goods_id, value, goods_name, avatar FROM association ORDER BY value DESC`
	ROWS, err1 := db.Query(query)
	if err1 != nil {
		log.Println(err1)
		return nil
	}
	for ROWS.Next() {
		var res model.Association
		err1 = ROWS.Scan(&res.Search_id, &res.Goods_id, &res.Value, &res.Goods_name, &res.Avatar)
		if err1 != nil {
			log.Println(err1)
			return nil
		}
		ans = append(ans, res)
	}
	return ans
}

func ComPare(x string, y string) int {
	charCount := make(map[rune]int)
	for _, char := range x {
		charCount[char]++
	}
	commonCount := 0
	for _, char := range y {
		if charCount[char] > 0 {
			commonCount++
			charCount[char]--
		}
	}

	return commonCount
}
