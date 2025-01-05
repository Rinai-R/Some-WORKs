package model

type Star struct {
	User_id  string `json:"user_id"`
	Goods_id string `json:"goods_id"`
}

type Browse struct {
	User_id  string `json:"user_id"` //此处id使用string类型是为了处理方便
	Goods_id string `json:"goods_id"`
}

type Praise struct {
	User_id    string `json:"user_id"`
	Message_id string `json:"message_id"`
}

type Search struct {
	Id      string `json:"id"`
	Content string `json:"content"`
}

type Association struct {
	Search_id  string `json:"search_id"`
	Goods_id   string `json:"goods_id"`
	Goods_name string `json:"goods_name"`
	Avatar     string `json:"avatar"`
	Value      int    `json:"value"`
}
