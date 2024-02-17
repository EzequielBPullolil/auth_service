package user

import (
	"context"
	"errors"
	"fmt"
	"log"
	"strings"

	"github.com/EzequielBPullolil/auth_service/src/types"
	"github.com/google/uuid"
	"github.com/jackc/pgx/v5/pgxpool"
)

type UserRepository struct {
	connectionPool *pgxpool.Pool
}

func NewUserRepository(pool *pgxpool.Pool) UserRepository {

	return UserRepository{
		connectionPool: pool,
	}
}

func (r UserRepository) ConnectionPool() *pgxpool.Pool {
	return r.connectionPool
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

func (r UserRepository) Create(userFields types.User) (types.User, error) {
	id, _ := uuid.NewUUID()
	userFields.Id = id.String()
	userFields.HashPassword()
	query := fmt.Sprintf("INSERT INTO users (id, name, password, email) VALUES('%s','%s','%s','%s');", userFields.Id, userFields.Name, userFields.Password, userFields.Email)
	log.Println(query)
	_, err := r.connectionPool.Exec(context.Background(), query)
	return userFields, err
}

func (r UserRepository) Read(email string) (*types.User, error) {
	var user types.User
	query := fmt.Sprintf("SELECT * FROM users WHERE email='%s';", email)

	r.connectionPool.QueryRow(context.Background(), query).Scan(&user.Id, &user.Name, &user.Email, &user.Password)
	if user.Name == "" {
		return nil, errors.New("unregistered user")
	}
	return &user, nil
}
func (r UserRepository) FindById(id string) (*types.User, error) {
	var user types.User
	query := fmt.Sprintf("SELECT * FROM users WHERE id='%s';", id)

	r.connectionPool.QueryRow(context.Background(), query).Scan(&user.Id, &user.Name, &user.Email, &user.Password)
	if user.Name == "" {
		return nil, errors.New("unregistered user")
	}
	return &user, nil
}

func (r UserRepository) Update(user_id string, new_fields types.User) (*types.User, error) {
	var user types.User
	if new_fields.Id != "" {
		return nil, errors.New("Can't update user ID")
	}
	if err, ok := r.validateFields(new_fields); !ok {
		return nil, err
	}

	if new_fields.Password != "" {
		new_fields.HashPassword()
	}
	err := r.connectionPool.QueryRow(context.Background(), `UPDATE users
	SET name = CASE
				WHEN $1 <> '' THEN $1
				ELSE name
			   END,
		password = CASE
				   WHEN $2 <> '' THEN $2
				   ELSE password
				  END,
		email = CASE
				WHEN $3 <> '' THEN $3
				ELSE email
			   END
	WHERE id = $4 RETURNING id, name, password, email;`, new_fields.Name, new_fields.Password, new_fields.Email, user_id).Scan(&user.Id, &user.Name, &user.Password, &user.Email)
	if err != nil {
		if strings.Contains(err.Error(), "SQLSTATE 23505") {
			return nil, errors.New("Cannot update user: Email already in use")
		}
		return nil, err
	}

	return &user, nil
}

func (r UserRepository) Delete(user_id string) error {
	if result, err := r.connectionPool.Exec(context.Background(), "DELETE FROM users WHERE id=$1;", user_id); err != nil {
		return err
	} else if result.RowsAffected() == 0 {
		return errors.New(fmt.Sprintf("There is no user with the ID: '%s'", user_id))
	}
	return nil
}

func (r UserRepository) validateFields(fields types.User) (error, bool) {
	if fields.Name != "" && !fields.ValidateName() {
		return errors.New("The field passed to update `name` is invalid"), false
	}

	if fields.Email != "" && !fields.ValidateEmaiL() {
		return errors.New("The field passed to update `email` is invalid"), false
	}
	if fields.Password != "" && !fields.ValidatePassword() {
		return errors.New("The field passed to update `password` is invalid"), false
	}

	return nil, true
}
