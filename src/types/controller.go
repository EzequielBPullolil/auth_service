package types

import (
	"encoding/json"
	"log"
	"net/http"
)

type Controller struct {
	repo Repository
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
