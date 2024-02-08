package users

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var repo UserRepository

func init() {
	repo = NewUserRepository()

	repo.CreateTables()
}
func TestCreateUser(t *testing.T) {
	user := User{
		Name:     "test_user",
		Password: "Test_password",
		Email:    "Test_email",
	}
	t.Run("Should be have id", func(t *testing.T) {
		user, err := repo.Create(user)
		assert.NoError(t, err)
		assert.NotEqual(t, user.GetId(), "")
	})
}
