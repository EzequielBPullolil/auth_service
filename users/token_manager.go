package users

import (
	"github.com/golang-jwt/jwt/v5"
)

func CreateToken(id string) string {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"id": id,
	})
	tokenString, _ := token.SignedString([]byte("secret"))
	return tokenString
}

func GetTokenId(tokenString string) string {
	t, _ := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return []byte("secret"), nil // Aquí debes usar la clave secreta con la que se firmó el token
	})
	if claims, ok := t.Claims.(jwt.MapClaims); ok && t.Valid {
		id := claims["id"].(string)
		return id
	}
	return ""
}
