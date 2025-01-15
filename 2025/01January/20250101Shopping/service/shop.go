package service

import (
	"Golang/2025/01January/20250101Shopping/dao"
	"Golang/2025/01January/20250101Shopping/model"
	"Golang/2025/01January/20250101Shopping/utils"
	"errors"
)

// RegisterMall 注册店铺
func RegisterMall(shop model.Shop) error {
	if dao.ShopExist(shop) {
		return errors.New("店铺已存在")
	}
	if !dao.RegisterMall(shop) {
		return errors.New("注册店铺失败")
	}
	return nil
}

// LoginMall 登录店铺
func LoginMall(shop model.Shop) (string, error) {
	if !dao.LoginMall(shop) {
		return "", errors.New("登录失败")
	}
	token, err := utils.GenerateShopJWT(shop.Shop_name)
	if err != nil {
		return "", err
	}
	return token, nil
}

// RegisterGoods 注册商品
func RegisterGoods(goods model.Goods, shop_name string) error {
	if !dao.ShopExist(model.Shop{Shop_name: shop_name}) {
		return errors.New("店铺不存在")
	}
	goods.Shop_id = dao.GetShopId(shop_name)

	if !dao.RegisterGoods(goods) {
		return errors.New("注册商品失败")
	}
	return nil
}

// GetShopAndGoodsInfo 获取店铺和商品信息
func GetShopAndGoodsInfo(shop *model.Shop) error {
	if !dao.GetShopAndGoodsInfo(shop) {
		return errors.New("获取店铺和商品信息失败")
	}
	return nil
}

// AlterGoodsInfo 修改商品信息
func AlterGoodsInfo(goods model.Goods) error {
	if !dao.AlterGoodsInfo(goods) {
		return errors.New("修改商品信息失败")
	}
	return nil
}

// DeleteGoods 删除商品
func DeleteGoods(goods model.Goods) error {
	if !dao.DeleteGoods(goods) {
		return errors.New("删除商品失败")
	}
	return nil
}
