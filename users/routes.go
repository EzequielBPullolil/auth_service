package auth

import (
	"net/http"

	"github.com/EzequielBPullolil/auth_service/common"
)

func HandleUserRoute(s *http.ServeMux, db_inyection common.Repository) {
	s.Handle("/users", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case "GET":
			user, _ := db_inyection.Read("as")
			response := user.ToJson()

			w.WriteHeader(200)
			w.Write([]byte(response))
		}
	}))
}
