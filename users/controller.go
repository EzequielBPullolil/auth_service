package users

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/EzequielBPullolil/auth_service/common"
)

type UserController struct {
	repo common.Repository
	common.Controller
}

func NewUserController(db_repository common.Repository) UserController {
	return UserController{
		repo: db_repository,
	}
}

func (uc UserController) GetAuthenticatedUser(res http.ResponseWriter, req *http.Request) {
	user, _ := uc.repo.Read("as")
	uc.ResponseWithStatus(user.ToJson(), 200, res)
}
func (uc UserController) UpdateAuthenticatedUser(res http.ResponseWriter, req *http.Request) {
	u := uc.GetUserData(req)
	updated_user, _ := uc.repo.Update("fake_id", u)

	response := fmt.Sprintf(`{
				"status": "Successful user update",
				"data": %s
			}`, updated_user.ToJson())

	uc.ResponseWithStatus(response, 200, res)
}
func (uc UserController) DeleteAuthenticatedUser(res http.ResponseWriter, req *http.Request) {
	uc.repo.Delete("fake_id")

	response := fmt.Sprintf(`{
			"status": "Successful user delete",
		}`)

	uc.ResponseWithStatus(response, 200, res)
}

func (uc UserController) GetUserById(res http.ResponseWriter, req *http.Request) {
	id := strings.TrimPrefix(req.URL.Path, "/users/")
	user, _ := uc.repo.Read(id)
	uc.ResponseWithStatus(user.ToJson(), 200, res)
}

func (uc UserController) HandleMethod(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		uc.GetAuthenticatedUser(w, r)
	case http.MethodPut:
		uc.UpdateAuthenticatedUser(w, r)
	case http.MethodDelete:
		uc.DeleteAuthenticatedUser(w, r)
	}
}
