package auth

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type User struct {
	Name, Email, Password string
}

func HandleAuthRoutes(r *http.ServeMux) {
	r.Handle("/auth/singup", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method == "POST" || r.Method == "post" {
			var u User
			json.NewDecoder(r.Body).Decode(&u)
			response := fmt.Sprintf(`{
				"status": "Successful user registration",
				"data":{
					"name":"%s",
					"email":"%s",
					
				}
			}`, u.Name, u.Email)
			w.WriteHeader(201)
			w.Write([]byte(response))
		}
	}))
}
