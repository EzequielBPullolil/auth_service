package auth

import (
	"net/http"

	"github.com/EzequielBPullolil/auth_service/common"
)

func HandleAuthRoutes(s *http.ServeMux, db_inyection common.Repository) {
	user_controller := NewUserController(db_inyection)
	s.Handle("/auth/singup", http.HandlerFunc(user_controller.SignupUser))
	s.Handle("/auth/login", http.HandlerFunc(user_controller.LoginUser))
	s.Handle("/auth/validate", http.HandlerFunc(user_controller.ValidateUserToken))
}
