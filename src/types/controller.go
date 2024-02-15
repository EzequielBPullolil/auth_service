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

func (c Controller) ResponseError(status string, err error, res http.ResponseWriter) {
	res.Header().Add("Content-Type", "application/json")
	res.WriteHeader(400)

	data := ResponseError{
		Status: status,
		Error:  err.Error(),
	}
	if err := json.NewEncoder(res).Encode(data); err != nil {
		log.Println("error in response: " + err.Error())
	}
}

func (c Controller) ResponseWithData(status string, data any, res http.ResponseWriter) {
	res.Header().Add("Content-Type", "application/json")

	d := ResponseWithData{
		Status: status,
		Data:   data,
	}

	if err := json.NewEncoder(res).Encode(d); err != nil {
		log.Println("error in response: " + err.Error())
	}
}
