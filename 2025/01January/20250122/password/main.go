package main

import (
	"Golang/2025/01January/20250122/password/utils"
	"fmt"
)

func main() {
	password := "123456"
	encryptedPassword1 := utils.EncryptPassword(password)
	fmt.Println(encryptedPassword1)
	encryptedPassword2 := utils.EncryptPassword(password)
	fmt.Println(encryptedPassword2)
	fmt.Println(utils.ComparePasswords(encryptedPassword1, password))
	fmt.Println(utils.ComparePasswords(encryptedPassword2, password))
	fmt.Println(utils.ComparePasswords(encryptedPassword2, "123446"))
	fmt.Println(utils.ComparePasswords(encryptedPassword1, encryptedPassword2))
}
