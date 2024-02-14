package users

import (
	"context"
	"errors"
	"fmt"
	"log"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5/pgxpool"
)

type Repository interface {
	Create(User) (User, error)
	Read(string) (*User, error)
	Delete(string) error
	Update(string, User) (*User, error)
	CreateTables() error
}

type UserRepository struct {
	connectionPool *pgxpool.Pool
}

func NewUserRepository(pool *pgxpool.Pool) UserRepository {

	return UserRepository{
		connectionPool: pool,
	}
}

func (r UserRepository) CreateTables() error {
	_, err := r.connectionPool.Exec(context.Background(), `CREATE TABLE IF NOT EXISTS users (
		id VARCHAR PRIMARY KEY,
		name VARCHAR NOT NULL,
		email VARCHAR UNIQUE NOT NULL,
		password VARCHAR NOT NULL
	);`)
	return err
}

func (r UserRepository) Create(userFields User) (User, error) {
	id, _ := uuid.NewUUID()
	userFields.Id = id.String()
	query := fmt.Sprintf("INSERT INTO users (id, name, password, email) VALUES('%s','%s','%s','%s');", userFields.Id, userFields.Name, userFields.Password, userFields.Email)
	log.Println(query)
	_, err := r.connectionPool.Exec(context.Background(), query)
	return userFields, err
}

func (r UserRepository) Read(email string) (*User, error) {
	var user User
	query := fmt.Sprintf("SELECT * FROM users WHERE email='%s';", email)

	r.connectionPool.QueryRow(context.Background(), query).Scan(&user.Id, &user.Name, &user.Email, &user.Password)
	if user.Name == "" {
		return nil, errors.New("unregistered user")
	}
	return &user, nil
}

func (r UserRepository) Update(user_id string, new_fields User) (*User, error) {
	var user User

	if new_fields.Id != "" {
		return nil, errors.New("Can't update user ID")
	}
	err := r.connectionPool.QueryRow(
		context.Background(),
		"UPDATE users SET name=$1, password=$2, email=$3 WHERE id=$4 RETURNING id, name, password, email;",
		new_fields.Name, new_fields.Password, new_fields.Email, user_id).Scan(&user.Id, &user.Name, &user.Email, &user.Password)
	if err != nil {
		return nil, err
	}

	return &user, err
}

func (r UserRepository) Delete(user_id string) error {
	if result, err := r.connectionPool.Exec(context.Background(), "DELETE FROM users WHERE id=$1;", user_id); err != nil {
		return err
	} else if result.RowsAffected() == 0 {
		return errors.New(fmt.Sprintf("There is no user with the ID: '%s'", user_id))
	}
	return nil
}
