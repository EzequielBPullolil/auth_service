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
