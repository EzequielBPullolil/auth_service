package user

import (
	"log"
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

	c, err := req.Cookie("auth_token")
	if err != nil {
		uc.ResponseError("Error finding cookie", err, res)
		return
	}

	_, email := tokenmanager.GetUserData(c.Value)
	user, err := uc.repo.Read(email)
	if err != nil {
		uc.ResponseError("unregistered user", err, res)
		return
	}
	uc.ResponseWithData("Successful user find", types.UserDAO{
		User: *user,
	}, res)

}
func (uc UserController) UpdateAuthenticatedUser(res http.ResponseWriter, req *http.Request) {
	c, err := req.Cookie("auth_token")
	if err != nil {
		uc.ResponseError("Error finding cookie", err, res)
		return
	}
	id := tokenmanager.GetTokenId(c.Value)
	u := uc.GetUserData(req)
	updated_user, err := uc.repo.Update(id, u)

	if err != nil {
		log.Println(u, updated_user, id, err)
		uc.ResponseError("error updating user", err, res)
		return
	}
	uc.ResponseWithData("Successful user update", types.UserDAO{
		User: *updated_user,
	}, res)
}
func (uc UserController) DeleteAuthenticatedUser(res http.ResponseWriter, req *http.Request) {
	c, err := req.Cookie("auth_token")
	if err != nil {
		uc.ResponseError("Error finding cookie", err, res)
		return
	}
	id := tokenmanager.GetTokenId(c.Value)
	if err := uc.repo.Delete(id); err != nil {
		log.Println(id, err, res.Header())
		uc.ResponseError("error deleting user", err, res)
		return
	}
	uc.ResponseWithData("Successful user delete", struct{}{}, res)
}

func (uc UserController) GetUserById(res http.ResponseWriter, req *http.Request) {
	id := strings.TrimPrefix(req.URL.Path, "/users/")
	user, err := uc.repo.Read(id)

	if err != nil {
		uc.ResponseError("error finding user by id", err, res)
		return
	}

	uc.ResponseWithData("finded user", user, res)
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
