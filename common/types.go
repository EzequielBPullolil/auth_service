package common

type Entity interface {
	GetId() string
}

type Repository interface {
	Create(any)
	Read(any) Entity
	Delete(any)
}
