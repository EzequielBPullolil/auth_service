package auth

import (
	"fmt"
	"log"
	"net/http"

	"github.com/EzequielBPullolil/auth_service/users"
)

type AuthController struct {
	repo users.Repository
	users.Controller
}

func NewAuthController(db_repository users.Repository) AuthController {
	return AuthController{
		repo: db_repository,
	}
}

func (uc AuthController) SignupUser(res http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {
		u := uc.GetUserData(r)
		if entity, err := uc.repo.Create(u); err == nil {
			response := fmt.Sprintf(`{
				"status": "Successful user registration",
				"data":{
					%s
				}
			}`, entity.ToJson())
			uc.ResponseWithStatus(response, http.StatusCreated, res)
			log.Println(u, entity, err)
		} else {
			uc.ResponseWithStatus(fmt.Sprintf(`{
				"error": "%s"
			}`, err.Error()), http.StatusCreated, res)
			log.Println(u, entity, err)
		}
	}
}
func (uc AuthController) LoginUser(res http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {
		u := uc.GetUserData(r)
		if user, err := uc.repo.GetBy("email", u.GetEmail()); err == nil {
			response := fmt.Sprintf(`{
				"status": "Successful user login",
				"data":{
					"token": "%s",
					"user": %s
				}
			}`, users.CreateToken(user.GetId()), user.ToJson())
			uc.ResponseWithStatus(response, http.StatusOK, res)
		} else {
			uc.ResponseWithStatus(fmt.Sprintf(`{
				"error": "%s"
			}`, err.Error()), http.StatusCreated, res)
		}
	}
}

func (uc AuthController) ValidateUserToken(res http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {
		c, err := r.Cookie("auth_token")

		if err != nil {
			uc.ResponseWithStatus(`{
				"status": "Invalid auth token",
			}`, http.StatusBadRequest, res)
			return
		}
		if id, ok := users.ValidateToken(c.Value); ok {
			log.Println(c, err, id, ok)
			if id == "" {
				uc.ResponseWithStatus(`{
					"status": "Invalid auth token",
				}`, http.StatusBadRequest, res)
			} else {
				uc.ResponseWithStatus(`{
					"status": "Valid auth token",
				}`, http.StatusOK, res)
			}
		} else {
			uc.ResponseWithStatus(`{
				"status": "Invalid auth token",
			}`, http.StatusBadRequest, res)
		}
		// Si la cookie no está presente o el token no es válido

	}
}
