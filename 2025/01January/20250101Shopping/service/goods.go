package service

import (
	"Golang/2025/01January/20250101Shopping/dao"
	"Golang/2025/01January/20250101Shopping/model"
	"errors"
)

// GetGoodsInfo 获取商品信息
func GetGoodsInfo(username string, browse model.Browse) (model.Goods, error) {
	var goods model.Goods
	browse.User_id = dao.GetId(username)

	if !dao.BrowseGoods(&goods, browse) {
		return model.Goods{}, errors.New("浏览商品失败")
	}
	return goods, nil
}

// BrowseRecords 获取浏览记录
func BrowseRecords(username string) ([]model.Browse, error) {
	var browse model.Browse
	browse.User_id = dao.GetId(username)

	if records, ok := dao.BrowseRecords(browse); ok {
		return records, nil
	}
	return nil, errors.New("获取浏览记录失败")
}

// AddGoodsToCart 增加商品到购物车
func AddGoodsToCart(username string, goods model.Goods) error {
	if mes, ok := dao.AddGoods(username, goods); !ok {
		return errors.New(mes)
	}
	return nil
}

// DelGoodsFromCart 从购物车中删除商品
func DelGoodsFromCart(username string, cart_goods model.Cart_Goods) error {
	cart_goods.User_Id = dao.GetId(username)
	if !dao.DelCartGoods(cart_goods) {
		return errors.New("删除购物车商品失败")
	}
	return nil
}

// GetCartInfo 获取购物车中的商品信息
func GetCartInfo(username string) (model.Shopping_Cart, error) {
	var cart model.Shopping_Cart
	cart.Id = dao.GetId(username)
	if cart.Id == "" || !dao.GetCartInfo(&cart) {
		return cart, errors.New("获取购物车信息失败")
	}
	return cart, nil
}

// SearchTypeGoods 根据类型搜索商品
func SearchTypeGoods(goods model.DisplayGoods) ([]model.DisplayGoods, error) {
	if data, ok := dao.SearchTypeGoods(&goods); ok {
		return data, nil
	}
	return nil, errors.New("搜索商品失败")
}

// StarGoods 收藏商品
func StarGoods(username string, star model.Star) error {
	star.User_id = dao.GetId(username)
	if !dao.StarGoods(star) {
		return errors.New("收藏商品失败")
	}
	return nil
}

// GetAllStar 获取所有收藏的商品
func GetAllStar(username string) ([]model.DisplayGoods, error) {
	var user model.User
	user.Id = dao.GetId(username)
	if goods, ok := dao.GetAllStar(user); ok {
		return goods, nil
	}
	return nil, errors.New("获取收藏商品失败")
}

// SearchGoods 搜索商品
func SearchGoods(search model.Search) ([]model.Association, error) {
	if lists := dao.SearchGoods(search); lists != nil {
		return lists, nil
	}
	return nil, errors.New("搜索商品失败")
}
