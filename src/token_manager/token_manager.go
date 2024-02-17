package tokenmanager

import (
	"errors"
	"os"

	"github.com/EzequielBPullolil/auth_service/src/types"
	"github.com/golang-jwt/jwt/v5"
)

func CreateToken(user types.User) (string, error) {
	if user.Email == "" {
		return "", errors.New("The 'email' field is empty")
	}
	if user.Id == "" {
		return "", errors.New("The 'id' field is empty")
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"id":    user.Id,
		"email": user.Email,
	})
	tokenString, _ := token.SignedString([]byte(os.Getenv("JWT_secret")))
	return tokenString, nil
}

func GetTokenId(tokenString string) string {
	t, _ := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return []byte(os.Getenv("JWT_secret")), nil
	})
	if claims, ok := t.Claims.(jwt.MapClaims); ok && t.Valid {
		id := claims["id"].(string)
		return id
	}
	return ""
}

func ValidateToken(tokenString string) bool {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return []byte(os.Getenv("JWT_secret")), nil
	})
	if err != nil {
		return false
	}
	return token.Valid
}

// Describes the user's email and the ID that the token has
// The token should be valid
func GetUserData(tokenString string) (string, string) {
	t, _ := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return []byte(os.Getenv("JWT_secret")), nil
	})
	claims, _ := t.Claims.(jwt.MapClaims)

	return claims["id"].(string), claims["email"].(string)
}
