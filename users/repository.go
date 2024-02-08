package users

type Repository interface {
	Create(User) (User, error)
	Read(string) (User, error)
	Delete(string) error
	Update(string, User) (User, error)
	CreateTables() error
}

type UserRepository struct {
}

func NewUserRepository() UserRepository {
	return UserRepository{}
}

func (r UserRepository) CreateTables() {

}

func (r UserRepository) Create(userFields User) (User, error) {
	userFields.Id = "an id"
	return userFields, nil
}
