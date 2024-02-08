package common

type Entity interface {
	GetId() string
	ToJson() string
}

type Repository interface {
	Create(Entity) (Entity, error)
	Read(string) (Entity, error)
	Delete(Entity) error
	Update(string, Entity) (Entity, error)
}
