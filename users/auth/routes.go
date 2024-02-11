package auth

import (
	"net/http"

	"github.com/EzequielBPullolil/auth_service/users"
)

func HandleAuthRoutes(s *http.ServeMux, db_inyection users.Repository) {
	auth_controller := NewAuthController(db_inyection)
	s.Handle("/auth/signup", http.HandlerFunc(auth_controller.SignupUser))
	s.Handle("/auth/login", http.HandlerFunc(auth_controller.LoginUser))
	s.Handle("/auth/validate", http.HandlerFunc(auth_controller.ValidateUserToken))
}
