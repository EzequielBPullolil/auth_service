package user_test

import (
	"context"
	"fmt"
	"log"
	"os"
	"testing"

	"github.com/EzequielBPullolil/auth_service/src/types"
	"github.com/EzequielBPullolil/auth_service/src/user"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/stretchr/testify/assert"
)

var repo user.UserRepository
var connection *pgx.Conn

func init() {
	var pool *pgxpool.Pool
	var err error
	if pool, err = pgxpool.New(context.Background(), os.Getenv("DB_TEST_URI")); err != nil {
		log.Fatal(err)
	}

	if connection, err = pgx.Connect(context.Background(), os.Getenv("DB_TEST_URI")); err != nil {
		log.Fatal(err)
	}

	if err = pool.Ping(context.Background()); err != nil {
		log.Fatal(err)
	}
	pool.QueryRow(context.Background(), "delete from users;")
	repo = user.NewUserRepository(pool)

	if err = repo.CreateTables(); err != nil {
		log.Fatal(err)
	}
}
func TestCreate(t *testing.T) {
	var user = types.User{
		Name:     "test_user",
		Password: "Test_password4",
		Email:    "email@create_test.com",
	}
	assert.Equal(t, user.GetId(), "")
	persistedUser, err := repo.Create(user)
	assert.NoError(t, err)
	t.Run("Should be have id", func(t *testing.T) {
		assert.NotEqual(t, persistedUser.GetId(), "")
	})
	t.Run("Should persist an user", func(t *testing.T) {
		var response_id string
		assert.NoError(t, connection.QueryRow(context.Background(), "SELECT id FROM users WHERE email='"+persistedUser.Email+"'").Scan(&response_id))
		assert.Equal(t, persistedUser.GetId(), response_id)
	})
	t.Run("Should not create an user with the same email", func(t *testing.T) {
		_, err := repo.Create(user)
		assert.Error(t, err)
		var count int
		connection.QueryRow(context.Background(), "SELECT COUNT(email) FROM users WHERE email='"+persistedUser.Email+"'").Scan(&count)
		assert.Equal(t, count, 1)
	})
	t.Run("Should persist a user with encrypted password", func(t *testing.T) {
		var newPassword string
		connection.QueryRow(context.Background(), "SELECT password FROM users WHERE email='"+persistedUser.Email+"'").Scan(&newPassword)

		assert.NotEmpty(t, newPassword)
		assert.NotEqual(t, user.Password, newPassword)

		t.Run("Password don't loss data", func(t *testing.T) {

			assert.False(t, user.ComparePassword(user.Password)) // Compare before hash password instance
			user.HashPassword()
			assert.True(t, user.ComparePassword(user.Password))
		})
	})
}

func TestRead(t *testing.T) {
	var userSuject = types.User{
		Id:       "some_id",
		Name:     "Analia",
		Password: "Password#34",
		Email:    "email@read_test.com",
	}
	query := fmt.Sprintf("INSERT INTO users (id, name, password, email) VALUES('%s','%s','%s','%s');",
		userSuject.Id,
		userSuject.Name,
		userSuject.Password,
		userSuject.Email)
	connection.Exec(context.Background(), query)

	t.Run("Return nil if the user non Exist", func(t *testing.T) {
		geted_user, err := repo.Read("not_registeredEmail@email.test.com")
		assert.ErrorContains(t, err, "unregistered user")
		assert.Nil(t, geted_user)
	})
	t.Run("Should return an User with the expected fiels", func(t *testing.T) {
		geted_user, err := repo.Read(userSuject.Email)
		assert.NoError(t, err)
		assert.NotNil(t, geted_user)
		assert.Equal(t, userSuject.Name, geted_user.Name)
		assert.Equal(t, userSuject.Email, geted_user.Email)
		assert.Equal(t, userSuject.Id, geted_user.Id)
	})
}

func TestUpdate(t *testing.T) {
	var userSuject = types.User{
		Id:       "a-user-to-update",
		Name:     "Granata",
		Password: "PasSWord#32",
		Email:    "email@update_test.com",
	}
	query := fmt.Sprintf("INSERT INTO users (id, name, password, email) VALUES('%s','%s','%s','%s');",
		userSuject.Id,
		userSuject.Name,
		userSuject.Password,
		userSuject.Email)
	_, err := connection.Exec(context.Background(), query)
	assert.NoError(t, err)
	var newFields = types.User{
		Name:     "Ezequiel",
		Password: "PasSWord#2",
		Email:    "newemail@test.com",
	}
	t.Run("Should return error if user non exist", func(t *testing.T) {
		u, err := repo.Update("nonExistingID", newFields)
		assert.Nil(t, u)
		assert.Error(t, err)
	})
	t.Run("Should update user in db", func(t *testing.T) {
		u, err := repo.Update(userSuject.Id, newFields)
		assert.NotNil(t, u)
		assert.NoError(t, err)

		assert.Equal(t, u.Name, newFields.Name)
		assert.Equal(t, u.Email, newFields.Email)
	})

	t.Run("Cant update ID", func(t *testing.T) {
		u, err := repo.Update(userSuject.Id, types.User{Id: "hola"})

		assert.Nil(t, u)
		assert.ErrorContains(t, err, "Can't update user ID")
	})
	t.Run("Should return error if try to update email with already registered email", func(t *testing.T) {
		_, err := connection.Exec(context.Background(), "INSERT INTO users (id, name, password, email) VALUES('##sds##', '','','test@registeredEmail.com')")
		assert.NoError(t, err)
		u, err := repo.Update(userSuject.Id, types.User{
			Email: "test@registeredEmail.com",
		})

		assert.ErrorContains(t, err, "Cannot update user: Email already in use")

		assert.Nil(t, u)
	})
}

func TestDelete(t *testing.T) {
	var userSuject = types.User{
		Id:       "a-user-to-delete",
		Name:     "Amalia",
		Password: "Password#45",
		Email:    "test@detelet_test.com",
	}
	query := fmt.Sprintf("INSERT INTO users (id, name, password, email) VALUES('%s','%s','%s','%s');",
		userSuject.Id,
		userSuject.Name,
		userSuject.Password,
		userSuject.Email)
	connection.Exec(context.Background(), query)

	t.Run("Should return error if User dont exist", func(t *testing.T) {
		err := repo.Delete("fake_id")
		assert.ErrorContains(t, err, "There is no user with the ID: 'fake_id'")
	})

	t.Run("Should delete User in DB if erro == nil ", func(t *testing.T) {
		assert.Nil(t, repo.Delete(userSuject.Id))
		var count int
		connection.QueryRow(context.Background(), "SELECT COUNT(id) FROM users WHERE id='"+userSuject.Id+"'").Scan(&count)
		assert.Equal(t, count, 0)
	})
}