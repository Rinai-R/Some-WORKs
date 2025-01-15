package service

import (
	"Golang/2025/01January/20250101Shopping/dao"
	"Golang/2025/01January/20250101Shopping/model"
	"errors"
)

// SubmitOrder 提交订单，生成功能
func SubmitOrder(username string) (model.Order, error) {
	var order model.Order
	order.User_id = dao.GetId(username)

	if !dao.SubmitOrder(&order) {
		return order, errors.New("提交订单失败")
	}
	return order, nil
}

// Confirm 确认订单
func ConfirmOrder(username string, order model.Order) (interface{}, string) {
	order.User_id = dao.GetId(username)
	return dao.ConfirmOrder(order)
}

// CancelOrder 取消订单
func CancelOrder(username string, order model.Order) error {
	order.User_id = dao.GetId(username)
	if !dao.CancelOrder(order) {
		return errors.New("取消订单失败")
	}
	return nil
}
