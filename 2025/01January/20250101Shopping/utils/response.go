package utils

import "Golang/2025/01January/20250101Shopping/model"

const (
	//ok
	statusOK = 10000
	//未授权或者授权错误
	statusUnauthorized = 40001
	//内部出现问题，包括但不限于绑定错误，未找到用户，未找到商品
	statusInternalError = 50000

	statusLackGoods = 40002

	statusBalanceLack = 40003

	statusOrderDeleted = 40004
)

// ErrRsp 绑定会出现的内部错误
func ErrRsp(err error) model.ResponseOnlyInfo {
	return model.ResponseOnlyInfo{
		Status: statusInternalError,
		Info:   "error " + err.Error(),
	}
}

// UnAuthorized token不正确或不存在引发的错误
func UnAuthorized() model.ResponseData {
	return model.ResponseData{
		Status: statusUnauthorized,
		Info:   "UnAuthorized",
	}
}

// OK 请求成功
func OK() model.ResponseData {
	return model.ResponseData{
		Status: statusOK,
		Info:   "ok",
	}
}

// Refused 由于请求的字段不正确（比如无法找到请求id对应的用户或商品）引发的错误
func Refused(info string) model.ResponseData {
	return model.ResponseData{
		Status: statusInternalError,
		Info:   info,
	}
}

// OkWithData 请求成功并返回数据
func OkWithData(data interface{}) model.ResponseData {
	return model.ResponseData{
		Status: statusOK,
		Info:   "ok",
		Data:   data,
	}
}

func LackGoods(lack interface{}) model.ResponseData {
	return model.ResponseData{
		Status: statusLackGoods,
		Info:   "LackGoods",
		Data:   lack,
	}
}

func BalanceLack() model.ResponseData {
	return model.ResponseData{
		Status: statusBalanceLack,
		Info:   "balance lack",
	}
}

func OrderDeleted() model.ResponseData {
	return model.ResponseData{
		Status: statusOrderDeleted,
		Info:   "OrderDeleted",
	}
}
