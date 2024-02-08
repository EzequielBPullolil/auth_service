package users

import "fmt"

type User struct {
	Id, Name, Email, Password string
}

func (u User) GetId() string {
	return u.Id
}

func (u User) ToJson() string {
	return fmt.Sprintf(`{
		"id": "%s",
		"name": "%s",
		"email": "%s",
	}`, u.Id, u.Name, u.Email)
}

type Repository interface {
	Create(User) (User, error)
	Read(string) (User, error)
	Delete(string) error
	Update(string, User) (User, error)
}
