package model

type ResponseData struct {
	Status int         `json:"status"`
	Info   string      `json:"info"`
	Data   interface{} `json:"data"`
}

type ResponseOnlyInfo struct {
	Status int    `json:"status"`
	Info   string `json:"info"`
}
