package common

type User struct {
	Id, Name, Email, Password string
}

func (u User) GetId() string {
	return u.Id
}
