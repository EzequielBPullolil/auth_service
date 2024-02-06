package auth

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type User struct {
	Name, Email, Password string
}

func HandleAuthRoutes(s *http.ServeMux) {
	s.Handle("/auth/singup", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
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

	s.Handle("/auth/login", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method == "POST" || r.Method == "post" {
			var u User
			json.NewDecoder(r.Body).Decode(&u)
			response := fmt.Sprintf(`{
				"status": "Successful user login",
				"data":{
					"token": "fake_token",
					"user":{
						"name":"%s",
						"email":"%s",
					}
					
				}
			}`, u.Name, u.Email)
			w.WriteHeader(201)
			w.Write([]byte(response))
		}
	}))
}
