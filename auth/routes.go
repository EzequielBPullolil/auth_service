package auth

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/EzequielBPullolil/auth_service/common"
)

func HandleAuthRoutes(s *http.ServeMux, db_inyection common.Repository) {
	s.Handle("/auth/singup", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method == "POST" || r.Method == "post" {
			var u common.User
			json.NewDecoder(r.Body).Decode(&u)
			entity, _ := db_inyection.Create(u)

			response := fmt.Sprintf(`{
				"status": "Successful user registration",
				"data":{
					%s
				}
			}`, entity.ToJson())
			w.WriteHeader(201)
			w.Write([]byte(response))
		}
	}))

	s.Handle("/auth/login", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method == "POST" || r.Method == "post" {
			var u common.User
			json.NewDecoder(r.Body).Decode(&u)
			user, _ := db_inyection.Read(u.GetId())
			response := fmt.Sprintf(`{
				"status": "Successful user login",
				"data":{
					"token": "%s",
					"user": %s
				}
			}`, CreateToken(user.GetId()), user.ToJson())
			w.WriteHeader(201)
			w.Write([]byte(response))
		}
	}))

	s.Handle("/auth/validate", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method == "POST" || r.Method == "post" {
			w.WriteHeader(200)
			w.Write([]byte(`{
				"status": "Valid auth token",
			}`))
		}
	}))
}
