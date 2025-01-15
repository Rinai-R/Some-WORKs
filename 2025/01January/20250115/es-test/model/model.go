package model

type Goods struct {
	Id         string  `json:"id"`
	Avatar     string  `json:"avatar"`
	Goods_name string  `json:"goods_name"`
	Shop_id    string  `json:"shop_id"`
	Content    string  `json:"content"`
	Type       string  `json:"Type"`
	Number     int     `json:"number"`
	Price      float64 `json:"price"`
	Star       int     `json:"star"`
}
