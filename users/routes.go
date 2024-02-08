package users

import (
	"net/http"
)

func HandleUserRoute(s *http.ServeMux, db_inyection Repository) {
	userController := NewUserController(db_inyection)
	s.Handle("/users", http.HandlerFunc(userController.HandleMethod))
	s.Handle("/users/", http.HandlerFunc(userController.GetUserById))
}
