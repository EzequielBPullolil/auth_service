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
