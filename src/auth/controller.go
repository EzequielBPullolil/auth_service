package auth

import (
	"errors"
	"log"
	"net/http"

	tokenmanager "github.com/EzequielBPullolil/auth_service/src/token_manager"
	"github.com/EzequielBPullolil/auth_service/src/types"
)

type AuthController struct {
	repo types.Repository
	types.Controller
}

func NewAuthController(db_repository types.Repository) AuthController {
	return AuthController{
		repo: db_repository,
	}
}

func (uc AuthController) SignupUser(res http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		return
	}
	u := uc.GetUserData(r)

	if err, ok := u.ValidateFields(); !ok {
		uc.ResponseError("error signup user", err, res)
		return
	}
	entity, err := uc.repo.Create(u)
	if err != nil {
		uc.ResponseError("Error while persisting user", err, res)
		return
	}
	res.WriteHeader(201)
	uc.ResponseWithData("Succesful user registration", types.UserDAO{
		User: entity,
	}, res)
}
func (uc AuthController) LoginUser(res http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		return
	}
	u := uc.GetUserData(r)
	user, err := uc.repo.Read(u.GetEmail())
	if err != nil {
		uc.ResponseError("Error while login user", err, res)
		return
	}
	res.WriteHeader(201)
	uc.ResponseWithData("Successful user login", types.TokenData{
		Token: tokenmanager.CreateToken(user.GetEmail()),
		User:  *user,
	}, res)

}

func (uc AuthController) ValidateUserToken(res http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		return
	}
	c, err := r.Cookie("auth_token")

	if err != nil {
		uc.ResponseError("Missing auth token", err, res)
		return
	}

	log.Println("\n" + c.Value)
	if !tokenmanager.ValidateToken(c.Value) {
		uc.ResponseError("Invalid auth token", errors.New(""), res)
		return
	}

	id := tokenmanager.GetTokenId(c.Value)
	if id == "" {
		uc.ResponseError("Invalid auth token", errors.New(""), res)
		return

	}
	uc.ResponseWithData("Valid auth token", struct{}{}, res)

}
