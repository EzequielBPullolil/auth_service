package auth

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/EzequielBPullolil/auth_service/common"
)

type UserController struct {
	repo common.Repository
}

func NewUserController(db_repository common.Repository) UserController {
	return UserController{
		repo: db_repository,
	}
}

func (uc UserController) SignupUser(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" || r.Method == "post" {
		var u common.User
		json.NewDecoder(r.Body).Decode(&u)
		entity, _ := uc.repo.Create(u)

		response := fmt.Sprintf(`{
			"status": "Successful user registration",
			"data":{
				%s
			}
		}`, entity.ToJson())
		w.WriteHeader(201)
		w.Write([]byte(response))
	}
}
func (uc UserController) LoginUser(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" || r.Method == "post" {
		var u common.User
		json.NewDecoder(r.Body).Decode(&u)
		user, _ := uc.repo.Read(u.GetId())
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
}

func (uc UserController) ValidateUserToken(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" || r.Method == "post" {
		w.WriteHeader(200)
		w.Write([]byte(`{
			"status": "Valid auth token",
		}`))
	}
}
