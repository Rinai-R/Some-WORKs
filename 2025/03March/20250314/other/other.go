package main

import (
	"fmt"
	"github.com/golang-jwt/jwt/v5"
)

func main() {
	for i := 0; i < 5; i++ {
		token := jwt.New(jwt.SigningMethodHS256)
		claims := token.Claims.(jwt.MapClaims)
		claims["GetName"] = "Some-WORK"
		tokenString, err := token.SignedString([]byte("secret"))
		if err != nil {
			fmt.Println(err)
			return
		}
		fmt.Println(tokenString)
	}
}
