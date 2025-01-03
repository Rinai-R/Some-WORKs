package model

import "time"

type User struct {
	Id       int     `json:"id"`
	Username string  `json:"username"`
	Password string  `json:"password"`
	Balance  float64 `json:"balance"`
	Avatar   string  `json:"avatar"`
}

type Msg struct {
	Id          int       `json:"id"`
	Parent_id   int       `json:"parent_id"`
	Content     string    `json:"content"`
	User_id     int       `json:"user_id"`
	Goods_id    string    `json:"goods_id"`
	Praised_num int       `json:"praised_num"`
	Create_at   time.Time `json:"create_at"`
	Updated_at  time.Time `json:"updated_at"`
}

type Goods struct {
	Id         int     `json:"id"`
	Avatar     string  `json:"avatar"`
	Goods_name string  `json:"goods_name"`
	Shop_id    int     `json:"shop_id"`
	Content    string  `json:"content"`
	Type       string  `json:"type"`
	Number     int     `json:"number"`
	Price      float64 `json:"price"`
	Star       int     `json:"star"`
}

type Cart_Goods struct {
	Goods_Id int     `json:"id"`
	Number   int     `json:"number"`
	Price    float64 `json:"price"`
	User_Id  int     `json:"user_id"`
}

type Shop struct {
	Id        int     `json:"id"`
	Shop_name string  `json:"shop_name"`
	Password  string  `json:"password"`
	Profit    float64 `json:"profit"`
}
