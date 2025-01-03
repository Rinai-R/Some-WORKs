package model

type Star struct {
	User_id  int `json:"user_id"`
	Goods_id int `json:"goods_id"`
}

type Browse struct {
	User_id  int `json:"user_id"`
	Goods_id int `json:"goods_id"`
}

type Praise struct {
	User_id    int `json:"user_id"`
	Message_id int `json:"message_id"`
}

type Search struct {
	Id      int    `json:"id"`
	Content string `json:"content"`
}

type Association struct {
	Search_id  int    `json:"search_id"`
	Goods_id   int    `json:"goods_id"`
	Goods_name string `json:"goods_name"`
	Avatar     string `json:"avatar"`
	Value      int    `json:"value"`
}
