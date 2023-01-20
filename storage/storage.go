package storage

import "github.com/gustagcosta/go-api/types"

type Storage interface {
	Connect(string) error
	GetDogs() (*[]types.Dog, error)
	GetDog(id int) (*types.Dog, error)
	StoreDog(name string, age int) error
	UpdateDog(*types.Dog) error
	DeleteDog(id int) error
}
