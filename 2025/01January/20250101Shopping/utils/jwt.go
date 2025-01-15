package utils

import (
	"errors"
	"fmt"
	"github.com/golang-jwt/jwt/v4"
	"time"
)

var UserSigningKey = []byte("Wishing")
var ShopSigningKey = []byte("Avid")

func GenerateShopJWT(Message string) (string, error) {
	token := jwt.New(jwt.SigningMethodHS256)
	claims := token.Claims.(jwt.MapClaims)
	claims["GetName"] = Message
	claims["exp"] = time.Now().Add(60 * time.Hour).Unix()
	TokenString, err := token.SignedString(ShopSigningKey)
	if err != nil {
		return "", err
	}
	return TokenString, nil
}

func GenerateUserJWT(Message string) (string, error) {
	token := jwt.New(jwt.SigningMethodHS256)
	claims := token.Claims.(jwt.MapClaims)
	claims["GetName"] = Message
	claims["exp"] = time.Now().Add(60 * time.Hour).Unix()
	TokenString, err := token.SignedString(UserSigningKey)
	if err != nil {
		return "", err
	}
	return TokenString, nil
}

func VerifyUserJWT(TokenString string) (string, error) {
	token, err := jwt.Parse(TokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}
		return UserSigningKey, nil
	})
	if err != nil {
		return "", err
	}
	if claims, ok1 := token.Claims.(jwt.MapClaims); ok1 {
		message := claims["GetName"].(string)
		return message, nil
	}
	return "", errors.New("token invalid")
}

func VerifyShopJWT(TokenString string) (string, error) {
	token, err := jwt.Parse(TokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}
		return ShopSigningKey, nil
	})
	if err != nil {
		return "", err
	}
	if claims, ok1 := token.Claims.(jwt.MapClaims); ok1 {
		message := claims["GetName"].(string)
		return message, nil
	}
	return "", errors.New("token invalid")
}
