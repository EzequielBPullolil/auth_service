package auth

import (
	"fmt"
	"net/http"

	"github.com/EzequielBPullolil/auth_service/common"
)

type AuthController struct {
	repo common.Repository
	common.Controller
}

func NewAuthController(db_repository common.Repository) AuthController {
	return AuthController{
		repo: db_repository,
	}
}

func (uc AuthController) SignupUser(res http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" || r.Method == "post" {
		u := uc.GetUserData(r)
		entity, _ := uc.repo.Create(u)

		response := fmt.Sprintf(`{
			"status": "Successful user registration",
			"data":{
				%s
			}
		}`, entity.ToJson())
		uc.ResponseWithStatus(response, 201, res)
	}
}
func (uc AuthController) LoginUser(res http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" || r.Method == "post" {
		u := uc.GetUserData(r)
		user, _ := uc.repo.Read(u.GetId())
		response := fmt.Sprintf(`{
			"status": "Successful user login",
			"data":{
				"token": "%s",
				"user": %s
			}
		}`, CreateToken(user.GetId()), user.ToJson())
		uc.ResponseWithStatus(response, 201, res)
	}
}

func (uc AuthController) ValidateUserToken(res http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" || r.Method == "post" {
		uc.ResponseWithStatus(`{
			"status": "Valid auth token",
		}`, 200, res)
	}
}
