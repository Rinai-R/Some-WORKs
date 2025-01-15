package model

type User struct {
	Id       int    `json:"id" gorm:"primary_key,auto_increment"`
	Name     string `json:"name" gorm:"type:varchar(25);unique"`
	Password string `json:"password" gorm:"type:varchar(255)"`
	Bio      string `json:"bio" gorm:"type:varchar(255)"`
}
