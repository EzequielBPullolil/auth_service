package tokenmanager

import (
	"os"

	"github.com/EzequielBPullolil/auth_service/src/types"
	"github.com/golang-jwt/jwt/v5"
)

func CreateToken(user types.User) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"id":    user.Id,
		"email": user.Email,
	})
	tokenString, _ := token.SignedString([]byte(os.Getenv("JWT_secret")))
	return tokenString, nil
}

func GetTokenId(tokenString string) string {
	t, _ := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return []byte("secret"), nil
	})
	if claims, ok := t.Claims.(jwt.MapClaims); ok && t.Valid {
		id := claims["id"].(string)
		return id
	}
	return ""
}

func ValidateToken(tokenString string) bool {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return []byte("secret"), nil
	})
	if err != nil || !token.Valid {
		return false
	}
	return true
}
