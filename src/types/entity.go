package types

import (
	"crypto/sha256"
	"fmt"
)

type User struct {
	Id, Name, Email, Password string

	hashedPassword bool
}

func (u User) ToJson() string {
	return fmt.Sprintf(`{
		"id": "%s",
		"name": "%s",
		"email": "%s",
		}`, u.Id, u.Name, u.Email)
}

func (u User) GetId() string    { return u.Id }
func (u User) GetEmail() string { return u.Email }

// Hash the password if it is not already hashed
func (u *User) HashPassword() {
	if !u.hashedPassword {
		u.Password = HashPassword(u.Password)
	}
	u.hashedPassword = true
}

func HashPassword(plainPassword string) string {
	h := sha256.New()
	h.Write([]byte(plainPassword))
	return fmt.Sprintf("%x", h.Sum(nil))
}
func (u User) ComparePassword(plainPassword string) bool {
	return u.hashedPassword || HashPassword(plainPassword) == u.Password
}
