package dao

import (
	"Golang/2025/01January/20250121/micro/app/model"
	"Golang/2025/01January/20250121/micro/server/user/conf"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var db *gorm.DB

func init() {
	var err error
	dsn := conf.Conf.Username + ":" + conf.Conf.Password + "@tcp(" + conf.Conf.Addr + ")/" + conf.Conf.DB
	db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	err = db.AutoMigrate(model.User{})
	if err != nil {
		return
	}
}
