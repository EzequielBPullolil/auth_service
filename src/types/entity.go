package types

import (
	"crypto/sha256"
	"errors"
	"fmt"
	"regexp"
)

var (
	InvalidName     error = errors.New("Invalid Name")
	InvalidEmail    error = errors.New("Invalid Email")
	InvalidPassword error = errors.New("Invalid Password")
)

type User struct {
	Id       string `json:"id"`
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`

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

func (u User) ValidateFields() (error, bool) {
	if !u.ValidateName() {
		return InvalidName, false
	}

	if !u.ValidateEmaiL() {
		return InvalidEmail, false
	}

	if !u.ValidatePassword() {
		return InvalidPassword, false
	}
	return nil, true
}

func (u User) ValidateName() bool {
	ok, _ := regexp.MatchString(`[^0-9\W_ ]+$`, u.Name)
	return len(u.Name) > 5 && ok
}

func (u User) ValidateEmaiL() bool {
	ok, _ := regexp.MatchString(`^[^@]+@[^@]+\.[a-zA-Z]{2,}$`, u.Email)
	return ok
}
func (u User) ValidatePassword() bool {
	ok, _ := regexp.MatchString(`[A-Z]+.*[a-z]+.*\d+.*[\W_]+|.*[a-z]+.*[A-Z]+.*\d+.*[\W_]+|.*\d+.*[a-z]+.*[A-Z]+.*[\W_]+|.*[\W_]+.*[a-z]+.*[A-Z]+.*\d+`, u.Password)
	return len(u.Password) > 7 && ok
}
