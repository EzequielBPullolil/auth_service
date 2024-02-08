package users

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"runtime"
	"strings"

	"github.com/EzequielBPullolil/auth_service/common"
)

func HandleUserRoute(s *http.ServeMux, db_inyection common.Repository) {
	s.Handle("/users", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodGet:
			user, _ := db_inyection.Read("as")
			w.WriteHeader(200)
			if _, err := w.Write([]byte(user.ToJson())); err != nil {
				_, _, line, _ := runtime.Caller(0)
				log.Printf("Error en la línea %d: %s\n", line, err.Error())
			}
		case http.MethodPut:
			var u common.User
			if err := json.NewDecoder(r.Body).Decode(&u); err != nil {
				_, _, line, _ := runtime.Caller(0)
				log.Fatalf("Error en la línea %d: %s\n", line, err.Error())
			}
			updated_user, _ := db_inyection.Update("fake_id", u)

			response := fmt.Sprintf(`{
				"status": "Successful user update",
				"data": %s
			}`, updated_user.ToJson())

			if _, err := w.Write([]byte(response)); err != nil {
				_, _, line, _ := runtime.Caller(0)
				log.Printf("Error en la línea %d: %s\n", line, err.Error())
			}
		case http.MethodDelete:
			db_inyection.Delete("fake_id")

			response := fmt.Sprintf(`{
				"status": "Successful user delete",
			}`)
			if _, err := w.Write([]byte(response)); err != nil {
				_, _, line, _ := runtime.Caller(0)
				log.Printf("Error en la línea %d: %s\n", line, err.Error())
			}
		}
	}))
	s.Handle("/users/", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		id := strings.TrimPrefix(r.URL.Path, "/users/")
		user, _ := db_inyection.Read(id)
		w.WriteHeader(200)
		if _, err := w.Write([]byte(user.ToJson())); err != nil {
			_, _, line, _ := runtime.Caller(0)
			log.Printf("Error en la línea %d: %s\n", line, err.Error())
		}
	}))
}
