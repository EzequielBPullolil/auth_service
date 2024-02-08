package users

import "github.com/EzequielBPullolil/auth_service/common"

type UserController struct {
	repo common.Repository
}

func NewUserController(db_repository common.Repository) UserController {
	return UserController{
		repo: db_repository,
	}
}
