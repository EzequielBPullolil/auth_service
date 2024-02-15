package types

type ResponseError struct {
	Status string `json:"status"`
	Error  string `json:"error"`
}

type ResponseWithData struct {
	Status string `json:"status"`
	Data   any    `json:"data"`
}

type TokenData struct {
	User  User   `json:"user"`
	Token string `json:"token"`
}

type UserDAO struct {
	User User `json:"user"`
}
