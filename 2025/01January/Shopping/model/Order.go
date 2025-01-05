package model

type Order struct {
	Id      string        `json:"id"`
	User_id string        `json:"user_id"`
	Shop_id string        `json:"shop_id"`
	Goods   []Order_Goods `json:"goods"`
	Sum     float64       `json:"sum"`
}
type Order_Goods struct {
	Goods_Id string  `json:"id"`
	Number   int     `json:"number"`
	Price    float64 `json:"price"`
	Order_id string  `json:"order_id"`
}
