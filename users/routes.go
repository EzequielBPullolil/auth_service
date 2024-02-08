package users

import (
	"net/http"

	"github.com/EzequielBPullolil/auth_service/common"
)

func HandleUserRoute(s *http.ServeMux, db_inyection common.Repository) {
	userController := NewUserController(db_inyection)
	s.Handle("/users", http.HandlerFunc(userController.HandleMethod))
	s.Handle("/users/", http.HandlerFunc(userController.GetUserById))
}
