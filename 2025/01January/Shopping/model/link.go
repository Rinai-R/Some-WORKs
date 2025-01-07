package model

import "time"

type Star struct {
	User_id  string `json:"user_id"`
	Goods_id string `json:"goods_id"`
}

type Browse struct {
	Id         string
	User_id    string `json:"user_id"` //此处id使用string类型是为了处理方便
	Goods_id   string
	Goods_name string
	Avatar     string
	BrowseTime time.Time
}

type Praise struct {
	User_id    string
	Message_id string `json:"message_id"`
}

type Search struct {
	Id      string
	Content string `json:"content"`
}

type Association struct {
	Search_id  string
	Goods_id   string
	Goods_name string
	Avatar     string
	Value      int
	Star       int
	Price      float64
	Type       string
}
