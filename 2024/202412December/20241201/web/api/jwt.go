package api

//关于jwt的所有操作都在这里
import (
	"errors"
	"fmt"
	"time"
)

var mySigningKey = []byte("Bang dream Girls Band Part!114514Girls Band Cry!114514,It's My Go!!!!!")

func GenerateJWT(username string) (string, error) {
	token := jwt.New(jwt.SigningMethodHS256)
	claims := token.Claims.(jwt.MapClaims)
	claims["username"] = username
	claims["exp"] = time.Now().Add(time.Second * 30).Unix()
	ToKenSting, err := token.SignedString(mySigningKey)
	if err != nil {
		return "", err
	}
	return ToKenSting, nil
}

func VerifyJWT(tokenString string) (string, error) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}
		return mySigningKey, nil
	})
	if err != nil {
		return "", err
	}
	if claims, ok1 := token.Claims.(jwt.MapClaims); ok1 && token.Valid {
		username := claims["username"].(string)
		return username, nil
	}
	return "", errors.New("token invalid")
}
