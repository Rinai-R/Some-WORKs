package utils

import "Golang/2025/01January/Shopping/model"

const (
	//ok
	statusOK            = 200
	//未授权或者授权错误
	statusUnauthorized  = 40001
	//内部出现问题，包括但不限于绑定错误，未找到用户，未找到商品
	statusInternalError = 50000
)

// ErrRsp 绑定会出现的内部错误
func ErrRsp(err error) model.ResponseOnlyInfo {
	return model.ResponseOnlyInfo{
		Status: statusInternalError,
		Info:   "error " + err.Error(),
	}
}

func UnAuthorized() model.ResponseData {
	return model.ResponseData{
		Status: statusUnauthorized,
		Info:   "UnAuthorized",
	}
}

func OK() model.ResponseData {
	return model.ResponseData{
		Status: statusOK,
		Info:   "ok",
	}
}

func Refused() model.ResponseData {
	return model.ResponseData{
		Status: statusInternalError,
		Info:   "Refused",
	}
}