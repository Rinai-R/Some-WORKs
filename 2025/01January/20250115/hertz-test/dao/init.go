package dao

import (
	"Golang/2025/01January/20250115/hertz-test/model"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var db *gorm.DB

func init() {
	var err error
	dsn := "root:~Cy710822@tcp(127.0.0.1:3306)/test"
	db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	err = db.AutoMigrate(model.User{})
	if err != nil {
		return
	}
}
