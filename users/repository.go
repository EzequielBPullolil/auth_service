package users

import (
	"context"
	"log"

	"github.com/jackc/pgx/v5/pgxpool"
)

type Repository interface {
	Create(User) (User, error)
	Read(string) (User, error)
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
		id SERIAL PRIMARY KEY,
		name VARCHAR NOT NULL,
		email VARCHAR UNIQUE NOT NULL,
		password VARCHAR NOT NULL
	);`)
	if err != nil {
		log.Println("Error creating tables")
	}
}

func (r UserRepository) Create(userFields User) (User, error) {
	userFields.Id = "an id"
	return userFields, nil
}
