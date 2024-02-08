package users

import (
	"context"
	"fmt"
	"log"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5/pgxpool"
)

type Repository interface {
	Create(User) (User, error)
	Read(string) (*User, error)
	Delete(string) error
	Update(string, User) (User, error)
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

func (r UserRepository) CreateTables() {
	_, err := r.connectionPool.Exec(context.Background(), `CREATE TABLE IF NOT EXISTS users (
		id VARCHAR PRIMARY KEY,
		name VARCHAR NOT NULL,
		email VARCHAR UNIQUE NOT NULL,
		password VARCHAR NOT NULL
	);`)
	if err != nil {
		log.Println("Error creating tables")
	}
}

func (r UserRepository) Create(userFields User) (User, error) {
	id, _ := uuid.NewUUID()
	userFields.Id = id.String()
	query := fmt.Sprintf("INSERT INTO users (id, name, password, email) VALUES('%s','%s','%s','%s');", userFields.Id, userFields.Name, userFields.Password, userFields.Email)
	log.Println(query)
	_, err := r.connectionPool.Exec(context.Background(), query)
	return userFields, err
}

func (r UserRepository) Read(user_id string) (*User, error) {
	var user User
	query := fmt.Sprintf("SELECT * FROM users WHERE id='%s';", user_id)

	r.connectionPool.QueryRow(context.Background(), query).Scan(&user.Id, &user.Name, &user.Email, &user.Password)
	if user.Name == "" {
		return nil, nil
	}
	return &user, nil
}
