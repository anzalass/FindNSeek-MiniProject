package middleware

import (
	"fmt"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/sirupsen/logrus"
)

func CreateToken(id string, name string) (string, error) {
	claims := jwt.MapClaims{}
	claims["id"] = id
	claims["name"] = name
	claims["exp"] = time.Now().Add(time.Hour * 100).Unix()

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte("anzalasganteng"))

}

func ExtractToken(tokenString string) (map[string]any, error) {

	claims := jwt.MapClaims{}
	_, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
		return []byte("anzalasganteng"), nil
	})
	// ... error handling
	if err != nil {
		logrus.Error("Error extracting")
	}

	// do something with decoded claims
	for key, val := range claims {
		fmt.Printf("Key: %v, value: %v\n", key, val)
	}
	return claims, nil
}
