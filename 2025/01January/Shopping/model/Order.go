package model

type Order struct {
	Id      string `json:"order_id"`
	User_id string
	Goods   []Order_Goods
	Sum     float64
}
type Order_Goods struct {
	Goods_Id   string
	Goods_Name string
	Number     int
	Price      float64
	Order_id   string
}
