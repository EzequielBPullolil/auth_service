package tokenmanager

import (
	"os"
	"testing"

	"github.com/EzequielBPullolil/auth_service/src/types"
	"github.com/golang-jwt/jwt/v5"
	"github.com/stretchr/testify/assert"
)

var secret string

func init() {
	secret = "test_secret_2024"
	os.Setenv("JWT_secret", secret)
}

func getClaims(token string) jwt.MapClaims {
	parsedToken, _ := jwt.Parse(token, func(t *jwt.Token) (interface{}, error) {
		return []byte(secret), nil
	})
	claims, ok := parsedToken.Claims.(jwt.MapClaims)

	if !ok {
		panic("error obtaining the claims")
	}
	return claims
}
func TestCreateToken(t *testing.T) {
	t.Run("A token generated should have expected fields", func(t *testing.T) {
		user_suject := types.User{
			Id:    "FakeId",
			Name:  "Ezequiel",
			Email: "ezequiel@test.com",
		}
		token, err := CreateToken(user_suject)
		assert.NoError(t, err)
		assert.NotEmpty(t, token)

		claims := getClaims(token)

		assert.Equal(t, "FakeId", claims["id"])
		assert.Equal(t, "ezequiel@test.com", claims["email"])
		assert.Nil(t, claims["name"])
	})
}
