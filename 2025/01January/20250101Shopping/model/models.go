package model

import (
	"time"
)

type User struct {
	Id       string  `json:"id"`
	Username string  `json:"username"`
	Password string  `json:"password"`
	Balance  float64 `json:"balance"`
	Avatar   string  `json:"avatar"`
	Nickname string  `json:"nickname"`
	Bio      string  `json:"bio"`
}

type Msg struct {
	Id          string `json:"id"`
	Parent_id   string `json:"parent_id"`
	Content     string `json:"content"`
	User_id     string `json:"user_id"`
	Goods_id    string `json:"goods_id"`
	Praised_num int
	Create_at   time.Time
	Updated_at  time.Time
	Response    []Msg
}

type Goods struct {
	Id         string  `json:"id"`
	Avatar     string  `json:"avatar"`
	Goods_name string  `json:"goods_name"`
	Shop_id    string  `json:"shop_id"`
	Content    string  `json:"content"`
	Type       string  `json:"type"`
	Number     int     `json:"number"`
	Price      float64 `json:"price"`
	Star       int     `json:"star"`
}

type Cart_Goods struct {
	Goods_Id   string `json:"id"`
	Goods_Name string
	Number     int     `json:"number"`
	Price      float64 `json:"price"`
	User_Id    string  `json:"user_id"`
}

type Shop struct {
	Id        string  `json:"id"`
	Shop_name string  `json:"shop_name"`
	Password  string  `json:"password"`
	Profit    float64 `json:"profit"`
	Goods     []DisplayGoods
}

type DisplayGoods struct {
	Avatar     string
	Goods_name string
	Type       string `json:"type"`
	Price      float64
	Star       int
}

type Shopping_Cart struct {
	Id    string `json:"id"`
	Sum   float64
	Goods []Cart_Goods
}

type Lack_Msg struct {
	Goods_id    string
	Current_Num int
	Query_Num   int
}
