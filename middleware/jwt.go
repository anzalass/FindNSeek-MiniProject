package middleware

import (
	"fmt"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/sirupsen/logrus"
)

func CreateToken(id string, name string, email string) (string, error) {
	claims := jwt.MapClaims{}
	claims["id"] = id
	claims["name"] = name
	claims["email"] = email
	claims["exp"] = time.Now().Add(time.Hour * 100).Unix()

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte("anzalasganteng"))

}

func ExtractToken(tokenString string) (map[string]any, error) {

	claims := jwt.MapClaims{}
	_, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
		return []byte("anzalasganteng"), nil
	})
	if err != nil {
		logrus.Error("Error extracting")
	}
	for key, val := range claims {
		fmt.Printf("Key: %v, value: %v\n", key, val)
	}
	return claims, nil
}
