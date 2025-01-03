package api

import (
	"errors"
	"fmt"
	"time"
)

var MySigningKey = []byte("KeyBoard!-RinRin")

func GenerateJWT(username string) (string, error) {
	token := jwt.New(jwt.SigningMethodHS256)
	claims := token.Claims.(jwt.MapClaims)
	claims["username"] = username
	claims["exp"] = time.Now().Add(time.Hour * 72).Unix()
	Sign, err := token.SignedString(MySigningKey)
	if err != nil {
		return "", err
	}
	return Sign, nil
}

func VerifyJWT(tokenString string) (string, error) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}
		return MySigningKey, nil
	})
	if err != nil {
		return "", err
	}

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		if time.Now().After(claims["exp"].(time.Time)) {
			return claims["username"].(string), nil
		}
		return "", errors.New("token expired")
	}
	return "", errors.New("token invalid")
}
