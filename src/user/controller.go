package user

import (
	"fmt"
	"net/http"
	"strings"

	tokenmanager "github.com/EzequielBPullolil/auth_service/src/token_manager"
	"github.com/EzequielBPullolil/auth_service/src/types"
)

type UserController struct {
	repo types.Repository
	types.Controller
}

func NewUserController(db_repository types.Repository) UserController {
	return UserController{
		repo: db_repository,
	}
}

func (uc UserController) GetAuthenticatedUser(res http.ResponseWriter, req *http.Request) {

	if c, err := req.Cookie("auth_token"); err == nil {
		id := tokenmanager.GetTokenId(c.Value)

		if user, err := uc.repo.Read(id); err == nil {

			uc.ResponseJsonWithStatus(user.ToJson(), http.StatusOK, res)
			return
		} else {
			uc.ResponseError("unregistered user", err, res)
			return
		}
	} else {
		uc.ResponseError("Error finding cookie", err, res)
	}
}
func (uc UserController) UpdateAuthenticatedUser(res http.ResponseWriter, req *http.Request) {
	c, _ := req.Cookie("auth_token")
	id := tokenmanager.GetTokenId(c.Value)
	u := uc.GetUserData(req)
	updated_user, _ := uc.repo.Update(id, u)

	response := fmt.Sprintf(`{
				"status": "Successful user update",
				"data": %s
			}`, updated_user.ToJson())

	uc.ResponseWithStatus(response, 200, res)
}
func (uc UserController) DeleteAuthenticatedUser(res http.ResponseWriter, req *http.Request) {
	c, _ := req.Cookie("auth_token")
	id := tokenmanager.GetTokenId(c.Value)
	uc.repo.Delete(id)

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
