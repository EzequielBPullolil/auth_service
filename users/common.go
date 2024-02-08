package users

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

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

type Controller struct {
}

func (c Controller) ResponseWithStatus(data string, statusCode int, res http.ResponseWriter) {
	res.WriteHeader(statusCode)

	if _, err := res.Write([]byte(data)); err != nil {
		log.Println("error in response" + err.Error())
	}
}

func (c Controller) GetUserData(req *http.Request) User {
	var u User

	json.NewDecoder(req.Body).Decode(&u)

	return u
}
