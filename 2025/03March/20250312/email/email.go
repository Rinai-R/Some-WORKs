package main

import (
	"fmt"
	"gopkg.in/gomail.v2"
)

func SendMail(To string, verificationCode string) {
	M := gomail.NewMessage()
	M.SetHeader("From", "rinai@g-rinai.cn")
	M.SetHeader("To", To)
	M.SetHeader("Subject", "Hello")
	M.SetBody("text/html", "Hello! your VerifyCode is "+verificationCode)

	d := gomail.NewDialer("smtp.exmail.qq.com", 465, "rinai@g-rinai.cn", "123456")

	if err := d.DialAndSend(M); err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("Email Sent!")
}

func main() {
	SendMail("xxx@qq.com", "123456")
}
