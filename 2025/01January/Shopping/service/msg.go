package service

import (
	"Golang/2025/01January/Shopping/dao"
	"Golang/2025/01January/Shopping/model"
	"errors"
)

// Publish 给商品评论
func Publish(username string, msg model.Msg) error {
	msg.User_id = dao.GetId(username)
	if !dao.PubMsg(msg) {
		return errors.New("发布评论失败")
	}
	return nil
}

// Response 回复评论
func Response(username string, msg model.Msg) error {
	msg.User_id = dao.GetId(username)
	if !dao.Response(msg) {
		return errors.New("回复评论失败")
	}
	return nil
}

// Praise 点赞评论
func Praise(username string, praise model.Praise) error {
	praise.User_id = dao.GetId(username)
	if !dao.Praise(praise) {
		return errors.New("点赞失败")
	}
	return nil
}

// GetGoodsMsg 获取关于商品的所有评论
func GetGoodsMsg(goods model.Goods) ([]model.Msg, error) {
	if data := dao.GetGoodsMsg(goods); data != nil {
		return data, nil
	}
	return nil, errors.New("获取评论失败")
}

// AlterMsg 修改评论内容
func AlterMsg(username string, msg model.Msg) error {
	msg.User_id = dao.GetId(username)
	if !dao.AlterMsg(msg) {
		return errors.New("修改评论失败")
	}
	return nil
}

// DeleteMsg 删除评论
func DeleteMsg(username string, msg model.Msg) error {
	msg.User_id = dao.GetId(username)
	if !dao.DelMsg(msg) {
		return errors.New("删除评论失败")
	}
	return nil
}
