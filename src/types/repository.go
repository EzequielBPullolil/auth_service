package types

type Repository interface {
	Create(User) (User, error)
	Read(string) (*User, error)
	FindById(string) (*User, error)
	Delete(string) error
	Update(string, User) (*User, error)
	CreateTables() error
}
