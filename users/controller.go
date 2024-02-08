package users

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"runtime"
	"strings"

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

func (uc UserController) GetAuthenticatedUser(res http.ResponseWriter, req *http.Request) {
	user, _ := uc.repo.Read("as")
	res.WriteHeader(200)
	if _, err := res.Write([]byte(user.ToJson())); err != nil {
		_, _, line, _ := runtime.Caller(0)
		log.Printf("Error en la línea %d: %s\n", line, err.Error())
	}
}
func (uc UserController) UpdateAuthenticatedUser(res http.ResponseWriter, req *http.Request) {
	var u common.User
	if err := json.NewDecoder(req.Body).Decode(&u); err != nil {
		_, _, line, _ := runtime.Caller(0)
		log.Fatalf("Error en la línea %d: %s\n", line, err.Error())
	}
	updated_user, _ := uc.repo.Update("fake_id", u)

	response := fmt.Sprintf(`{
				"status": "Successful user update",
				"data": %s
			}`, updated_user.ToJson())

	if _, err := res.Write([]byte(response)); err != nil {
		_, _, line, _ := runtime.Caller(0)
		log.Printf("Error en la línea %d: %s\n", line, err.Error())
	}
}
func (uc UserController) DeleteAuthenticatedUser(res http.ResponseWriter, req *http.Request) {
	uc.repo.Delete("fake_id")

	response := fmt.Sprintf(`{
			"status": "Successful user delete",
		}`)

	if _, err := res.Write([]byte(response)); err != nil {
		_, _, line, _ := runtime.Caller(0)
		log.Printf("Error en la línea %d: %s\n", line, err.Error())
	}
}

func (uc UserController) GetUserById(res http.ResponseWriter, req *http.Request) {
	id := strings.TrimPrefix(req.URL.Path, "/users/")
	user, _ := uc.repo.Read(id)
	res.WriteHeader(200)
	if _, err := res.Write([]byte(user.ToJson())); err != nil {
		_, _, line, _ := runtime.Caller(0)
		log.Printf("Error en la línea %d: %s\n", line, err.Error())
	}
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
