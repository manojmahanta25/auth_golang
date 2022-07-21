package services

import (
	"fmt"
	"github.com/golang-jwt/jwt"
	"os"
	"time"
)

var signKey = []byte(os.Getenv("JWT_SECRET"))

func GenerateJWTToken(data any) (string, error) {
	token := jwt.New(jwt.SigningMethodHS256)
	claims := token.Claims.(jwt.MapClaims)
	claims["authorized"] = true
	claims["iss"] = "AuthServer"
	claims["sub"] = data
	claims["iat"] = time.Now().Unix()
	claims["exp"] = time.Now().Add(time.Minute * 60).Unix()
	tokenString, err := token.SignedString(signKey)
	if err != nil {
		fmt.Errorf("something went wrong: %s", err.Error())
		return "", err
	}
	return tokenString, nil
}
func JwtVerify(tokenString string) (string, error) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (i interface{}, err error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}
		return signKey, nil
	})
	if err != nil {
		return "", err
	}
	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		var sub = fmt.Sprintf("%v", claims["sub"])
		return sub, nil

	} else {
		return "", err
	}
}
