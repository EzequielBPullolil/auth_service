package common

type Entity interface {
	GetId() string
}

type Repository interface {
	Create(Entity) (Entity, error)
	Read(Entity) (Entity, error)
	Delete(Entity) error
}
