package auth

import (
	"github.com/golang-jwt/jwt/v5"
)

func CreateToken(id string) string {
	token := jwt.NewWithClaims(jwt.SigningMethodES256, jwt.MapClaims{
		"id": id,
	})
	tokenString, _ := token.SignedString(token)
	return tokenString
}
