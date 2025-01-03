package model

type Order struct {
	Id      int           `json:"id"`
	User_id int           `json:"user_id"`
	Shop_id int           `json:"shop_id"`
	Goods   []Order_Goods `json:"goods"`
	Sum     float64       `json:"sum"`
}
type Order_Goods struct {
	Goods_Id int     `json:"id"`
	Number   int     `json:"number"`
	Price    float64 `json:"price"`
	Order_id int     `json:"order_id"`
}
