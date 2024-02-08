package common

import (
	"log"
	"net/http"
)

type Controller struct {
}

func (c Controller) ResponseWithStatus(data string, statusCode int, res http.ResponseWriter) {
	res.WriteHeader(statusCode)

	if _, err := res.Write([]byte(data)); err != nil {
		log.Println("error in response" + err.Error())
	}
}
