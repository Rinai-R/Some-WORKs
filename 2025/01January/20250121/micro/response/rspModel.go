package response

const (
	//ok
	Success       = 10000
	BindError     = 40001
	InternalError = 50001
	RegisterError = 40002
	PasswordErr   = 40003
	TokenErr      = 40004
)

type ResponseWithData struct {
	Code int         `json:"code"`
	Info string      `json:"info"`
	Data interface{} `json:"data"`
}

type Response struct {
	Code int    `json:"code"`
	Info string `json:"info"`
}
