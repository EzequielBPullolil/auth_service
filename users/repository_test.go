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
	pool.QueryRow(context.Background(), "delete from users;")
	repo = NewUserRepository(pool)

	repo.CreateTables()
}
func TestCreate(t *testing.T) {
	var user = User{
		Name:     "test_user",
		Password: "Test_password",
		Email:    "Test_email",
	}
	assert.Equal(t, user.GetId(), "")
	persistedUser, err := repo.Create(user)
	assert.NoError(t, err)
	query := fmt.Sprintf("SELECT id FROM users WHERE name='%s'", persistedUser.Name)
	t.Run("Should be have id", func(t *testing.T) {
		assert.NoError(t, err)
		assert.NotEqual(t, persistedUser.GetId(), "")
	})
	t.Run("Should persist an user", func(t *testing.T) {
		var response string
		assert.NoError(t, err)
		assert.NoError(t, pool.QueryRow(context.Background(), query).Scan(&response))
		assert.Equal(t, persistedUser.GetId(), response)
	})
	t.Run("Shoul not create an user with the same email", func(t *testing.T) {
		_, err := repo.Create(user)
		assert.Error(t, err)
		var count int
		repo.connectionPool.QueryRow(context.Background(), "SELECT COUNT(email) FROM users WHERE email='"+persistedUser.Email+"'").Scan(&count)
		assert.Equal(t, count, 1)
	})
}

func TestRead(t *testing.T) {
	var userSuject = User{
		Id:       "some_id",
		Name:     "an natural name",
		Password: "Some password",
		Email:    "An email",
	}
	query := fmt.Sprintf("INSERT INTO users (id, name, password, email) VALUES('%s','%s','%s','%s');",
		userSuject.Id,
		userSuject.Name,
		userSuject.Password,
		userSuject.Email)
	repo.connectionPool.Exec(context.Background(), query)

	t.Run("Return nil if the user non Exist", func(t *testing.T) {
		geted_user, err := repo.Read("NotRegisteredId")
		assert.NoError(t, err)
		assert.Nil(t, geted_user)
	})
	t.Run("Should return an User with the expected fiels", func(t *testing.T) {
		geted_user, err := repo.Read(userSuject.Id)
		assert.NoError(t, err)
		assert.NotNil(t, geted_user)
		assert.Equal(t, userSuject.Name, geted_user.Name)
		assert.Equal(t, userSuject.Email, geted_user.Email)
		assert.Equal(t, userSuject.Id, geted_user.Id)
	})
}
