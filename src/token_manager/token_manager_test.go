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

		parsedToken, _ := jwt.Parse(token, func(t *jwt.Token) (interface{}, error) {
			return []byte(secret), nil
		})

		claims, ok := parsedToken.Claims.(jwt.MapClaims)

		assert.True(t, ok)
		assert.Equal(t, "FakeId", claims["id"])
		assert.Equal(t, "ezequiel@test.com", claims["email"])
		assert.Equal(t, "", claims["name"])
	})
}
