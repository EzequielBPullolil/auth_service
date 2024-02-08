package users

import (
	"context"
	"fmt"
	"log"
	"testing"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/stretchr/testify/assert"
)

var repo UserRepository
var pool *pgxpool.Pool

func init() {
	var err error
	pool, err = pgxpool.New(context.Background(), "postgresql://ezequiel-k:ezequiel_dev_pass@localhost:5432/auth_systemtest")
	if err != nil {
		log.Fatal(err)
	}

	if err = pool.Ping(context.Background()); err != nil {
		log.Fatal(err)
	}

	defer pool.Close()
	repo = NewUserRepository(pool)

	repo.CreateTables()
}
func TestCreateUser(t *testing.T) {
	var user = User{
		Name:     "test_user",
		Password: "Test_password",
		Email:    "Test_email",
	}
	t.Run("Should be have id", func(t *testing.T) {
		assert.Equal(t, user.GetId(), "")
		persistedUser, err := repo.Create(user)
		assert.NoError(t, err)
		assert.NotEqual(t, persistedUser.GetId(), "")
	})
	t.Run("Should persist an user", func(t *testing.T) {
		var response string
		persistedUser, err := repo.Create(user)
		query := fmt.Sprintf("SELECT id FROM users WHERE name='%s'", persistedUser.Name)
		assert.NoError(t, err)
		assert.NoError(t, pool.QueryRow(context.Background(), query).Scan(&response))
		assert.Equal(t, persistedUser.GetId(), response)
	})
}
