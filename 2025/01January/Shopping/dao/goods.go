package dao

import (
	"Golang/2025/01January/Shopping/model"
	"database/sql"
	"encoding/json"
	"errors"
	"log"
	"strconv"
	"time"
)

func GetGoodInfo(goods_id string) (int, float64) {
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

func GetShopAndGoodsInfo(shop *model.Shop) bool {
	query := `SELECT id, shop_name, profit FROM shop WHERE shop_name = ?`
	err := db.QueryRow(query, shop.Shop_name).Scan(&shop.Id, &shop.Shop_name, &shop.Profit)
	if err != nil {
		log.Println(err)
		return false
	}
	var goods []model.DisplayGoods//利用切片存储多个商品内容，方便返回
	query = `SELECT goods_name, type, price, star, avatar FROM goods WHERE shop_id = ?`
	Rows, err0 := db.Query(query, shop.Id)
	if err0 != nil {
		if errors.Is(err0, sql.ErrNoRows) {
			return true
		}
		log.Println(err)
		return false
	}
	for Rows.Next() {
		var tmp model.DisplayGoods
		err = Rows.Scan(&tmp.Goods_name, &tmp.Type, &tmp.Price, &tmp.Star, &tmp.Avatar)
		if err != nil {
			log.Println(err)
			return false
		}
		goods = append(goods, tmp)
	}
	shop.Goods = goods
	return true
}

func GetGoodsName(id string) string {
	query := `SELECT goods_name FROM goods WHERE id = ?`
	var name string
	err := db.QueryRow(query, id).Scan(&name)
	if err != nil {
		log.Println(err)
		return ""
	}
	return name
}

// AddGoods 向购物车中添加商品
func AddGoods(username string, goods model.Goods) (string, bool) {
	id := GetId(username)
	num, price := GetGoodInfo(goods.Id)
	if num < goods.Number {
		return "lack", false
	}
	goods.Goods_name = GetGoodsName(goods.Id)
	query := `insert into cart_goods (user_id, goods_id, goods_name, number, price) values (?, ?, ?, ?, ?)`
	_, err1 := db.Exec(query, id, goods.Id, goods.Goods_name, goods.Number, price)
	if err1 != nil {
		log.Println(err1)
		return err1.Error(), false
	}
	//更新购物车中的总价格
	query = `update shopping_cart set sum = sum + ? where user_id = ?`
	sum := float64(goods.Number) * price
	_, err2 := db.Exec(query, sum, id)
	if err2 != nil {
		log.Println(err2)
		return err2.Error(), false
	}
	return "", true
}

// GetCartGoodsInfo 查询购物车中某个物品的价格和数量
func GetCartGoodsInfo(cart_goods model.Cart_Goods) (int, float64) {
	query := `SELECT number, price FROM cart_goods WHERE user_id = ? AND goods_id = ?`
	var num int
	var price float64
	err := db.QueryRow(query, cart_goods.User_Id, cart_goods.Goods_Id).Scan(&num, &price)
	if err != nil {
		log.Println(err)
		return -1, -1.0
	}
	return num, price
}

// GetCartInfo 传入指针，无需再创建变量，直接使用传入的参数即可
func GetCartInfo(cart *model.Shopping_Cart) bool {
	query := `SELECT sum FROM shopping_cart WHERE user_id = ?`
	err := db.QueryRow(query, cart.Id).Scan(&cart.Sum)
	if err != nil {
		log.Println(err)
		return false
	}
	query = `SELECT goods_id, goods_name ,number, price FROM cart_goods WHERE user_id = ?`
	Rows, err0 := db.Query(query, cart.Id)
	if err0 != nil {
		if errors.Is(err0, sql.ErrNoRows) {
			return false
		}
		log.Println(err0)
		return false
	}
	for Rows.Next() {
		var cart_goods model.Cart_Goods
		err = Rows.Scan(&cart_goods.Goods_Id, &cart_goods.Goods_Name, &cart_goods.Number, &cart_goods.Price)
		if err != nil {
			log.Println(err)
			return false
		}
		//切片存储商品信息
		cart.Goods = append(cart.Goods, cart_goods)
	}
	return true
}

func DelCartGoods(cart_goods model.Cart_Goods) bool {
	num, price := GetCartGoodsInfo(cart_goods)
	if num == -1 {
		return false
	}
	total := float64(num) * price
	//更新购物车的总价格
	query := `UPDATE shopping_cart SET sum = sum - ? WHERE user_id = ?`
	_, err := db.Exec(query, total, cart_goods.User_Id)
	if err != nil {
		log.Println(err)
		return false
	}
	//删
	query = `DELETE FROM cart_goods WHERE user_id = ? AND goods_id = ?`
	_, err = db.Exec(query, cart_goods.User_Id, cart_goods.Goods_Id)
	if err != nil {
		log.Println(err)
		return false
	}
	return true
}

// BrowseGoods 浏览商品详细信息
func BrowseGoods(goods *model.Goods, Browse model.Browse) bool {
	//先查询redis缓存的数据
	CacheKey := "goods:" + Browse.Goods_id
	CacheData, err := rdb.Get(ctx, CacheKey).Result()
	if err == nil {
		if err = json.Unmarshal([]byte(CacheData), &goods); err == nil {
			query := `INSERT INTO browse_records (user_id, goods_id, goods_name, avatar) values(?, ?, ?, ?) `
			_, err = db.Exec(query, Browse.User_id, Browse.Goods_id, goods.Goods_name, goods.Avatar)
			if err == nil {
				log.Println("成功缓存")
				return true
			}
		}
	}
	//没查询到，查mysql数据库
	query := `SELECT id, goods_name, shop_id, type, number, price, star, content, avatar FROM goods WHERE id = ?`
	err = db.QueryRow(query, Browse.Goods_id).Scan(&goods.Id, &goods.Goods_name, &goods.Shop_id, &goods.Type, &goods.Number, &goods.Price, &goods.Star, &goods.Content, &goods.Avatar)
	if err != nil {
		log.Println(err)
		return false
	}
	CachedData, err := json.Marshal(goods)
	if err != nil {
		log.Println(err)
		return false
	}
	//将查询的拘束存入redis，并设置过期时间
	rdb.Set(ctx, CacheKey, CachedData, time.Hour)
	//注册浏览记录，便于查询
	query = `INSERT INTO browse_records (user_id, goods_id, goods_name, avatar) values(?, ?, ?, ?) `
	_, err = db.Exec(query, Browse.User_id, Browse.Goods_id, goods.Goods_name, goods.Avatar)
	if err != nil {
		log.Println(err)
		return false
	}
	log.Println("通过MySQL得到数据")
	return true
}
//浏览记录
func BrowseRecords(Browse model.Browse) ([]model.Browse, bool) {
	query := `SELECT id, user_id, goods_id, goods_name, avatar, browse_time FROM browse_records WHERE user_id = ? ORDER BY browse_time DESC`
	rows, err := db.Query(query, Browse.User_id)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, true
		}
		log.Println(err)
		return nil, false
	}
	defer func(rows *sql.Rows) {
		err := rows.Close()
		if err != nil {
			log.Println(err)
		}
	}(rows)
	var records []model.Browse
	for rows.Next() {
		var browse model.Browse
		err = rows.Scan(&browse.Id, &browse.User_id, &browse.Goods_id, &browse.Goods_name, &browse.Avatar, &browse.BrowseTime)
		if err != nil {
			log.Println(err)
			return nil, false
		}
		//切片存储多条记录
		records = append(records, browse)
	}
	return records, true

}

// StarGoods 收藏商品
func StarGoods(star model.Star) bool {
	query := `SELECT 1 FROM star WHERE user_id = ? AND goods_id = ?`
	var exist bool
	err := db.QueryRow(query, star.User_id, star.Goods_id).Scan(&exist)
	if err != nil && !errors.Is(err, sql.ErrNoRows) {
		log.Println(err)
		return false
	}
	if !exist {
		query = `INSERT INTO star (user_id, goods_id) values(?, ?)`
		_, err = db.Exec(query, star.User_id, star.Goods_id)
		if err != nil {
			log.Println(err)
			return false
		}
		query = `UPDATE goods SET star = star + 1 WHERE id = ?`
		_, err = db.Exec(query, star.Goods_id)
		if err != nil {
			log.Println(err)
			return false
		}
	} else {
		query = `DELETE FROM star WHERE goods_id = ? AND user_id = ?`
		_, err = db.Exec(query, star.Goods_id, star.User_id)
		if err != nil {
			log.Println(err)
			return false
		}
		query = `UPDATE goods SET star = star - 1 WHERE id = ?`
		_, err = db.Exec(query, star.Goods_id)
		if err != nil {
			log.Println(err)
			return false
		}
	}
	return true
}

// GetAllStar 查询收藏的商品
func GetAllStar(user model.User) ([]model.DisplayGoods, bool) {
	var ans []model.DisplayGoods
	query := `SELECT goods_id FROM star WHERE user_id = ?`
	Rows, err := db.Query(query, user.Id)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, true
		}
		log.Println(err)
		return nil, false
	}
	for Rows.Next() {
		var id string
		err = Rows.Scan(&id)
		if err != nil {
			log.Println(err)
			return nil, false
		}
		var goods model.DisplayGoods
		query = `SELECT goods_name, type, price, star, avatar FROM goods WHERE id = ?`
		err = db.QueryRow(query, id).Scan(&goods.Goods_name, &goods.Type, &goods.Price, &goods.Star, &goods.Avatar)
		if err != nil {
			log.Println(err)
			return nil, false
		}
		ans = append(ans, goods)
	}
	return ans, true
}

// SearchTypeGoods 根据类型查询商品
func SearchTypeGoods(goods *model.DisplayGoods) ([]model.DisplayGoods, bool) {
	query := `SELECT goods_name, type, price, star, avatar  FROM goods WHERE type = ? ORDER BY star DESC `
	rows, err := db.Query(query, goods.Type)
	if err != nil {
		log.Println(err)
		return nil, false
	}
	var ans []model.DisplayGoods
	for rows.Next() {
		var res model.DisplayGoods
		err = rows.Scan(&res.Goods_name, &res.Type, &res.Price, &res.Star, &res.Avatar)
		if err != nil {
			log.Println(err)
			return nil, false
		}
		ans = append(ans, res)
	}
	return ans, true
}

// SearchGoods 搜索商品
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
	search.Id = strconv.FormatInt(Id, 10)
	return AssociationCount(search)
}

// AssociationCount 计算搜索内容与商品信息的相关性。依照相关性排序
func AssociationCount(search model.Search) []model.Association {
	query := `SELECT id, goods_name, type, price, star, avatar, content FROM goods `
	rows, err := db.Query(query)
	if err != nil {
		log.Println(err)
		return nil
	}
	var ans []model.Association
	for rows.Next() {
		tmp := struct {
			model.Association
			content string
		}{}
		tmp.Search_id = search.Id
		err = rows.Scan(&tmp.Goods_id, &tmp.Goods_name, &tmp.Type, &tmp.Price, &tmp.Star, &tmp.Avatar, &tmp.content)
		if err != nil {
			log.Println(err)
			return nil
		}
		tmp.Value = ComPare(search.Content, tmp.content) + ComPare(search.Content, tmp.Goods_name)
		query = `INSERT INTO association (search_id, goods_id, value, goods_name, avatar, price, type, star) values(?, ?, ?, ?, ?, ?, ?, ?)`
		_, err = db.Exec(query, tmp.Search_id, tmp.Goods_id, tmp.Value, tmp.Goods_name, tmp.Avatar, tmp.Price, tmp.Type, tmp.Star)
		if err != nil {
			log.Println(err)
			return nil
		}
	}
	query = `SELECT search_id, goods_id, value, goods_name, avatar, price, type, star FROM association WHERE search_id = ? ORDER BY value DESC `
	ROWS, err1 := db.Query(query, search.Id)
	if err1 != nil {
		log.Println(err1)
		return nil
	}
	for ROWS.Next() {
		var res model.Association
		err1 = ROWS.Scan(&res.Search_id, &res.Goods_id, &res.Value, &res.Goods_name, &res.Avatar, &res.Price, &res.Type, &res.Star)
		if err1 != nil {
			log.Println(err1)
			return nil
		}
		ans = append(ans, res)
	}
	return ans
}

// ComPare 工具函数，用于输出两个字符串的相同部分的数量
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
