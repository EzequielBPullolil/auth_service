package usermodule

import (
	"net/http"

	"github.com/EzequielBPullolil/auth_service/src/types"
)

func HandleUserRoute(s *http.ServeMux, db_inyection types.Repository) {
	userController := NewUserController(db_inyection)
	s.Handle("/users", http.HandlerFunc(userController.HandleMethod))
	s.Handle("/users/", http.HandlerFunc(userController.GetUserById))
}
