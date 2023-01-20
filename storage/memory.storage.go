package storage

import (
	"errors"
	"math/rand"

	"github.com/gustagcosta/go-api/types"
	"golang.org/x/exp/slices"
)

type MemoryStorage struct {
	dogs []types.Dog
}

func (s *MemoryStorage) Connect(connectionString string) error {
	dogs := []types.Dog{
		{
			ID:   1,
			Name: "Totó",
			Age:  12,
		},
		{
			ID:   2,
			Name: "Pretinho",
			Age:  8,
		},
		{
			ID:   3,
			Name: "Luca",
			Age:  7,
		},
	}

	s.dogs = dogs

	return nil
}

func (s *MemoryStorage) GetDogs() (*[]types.Dog, error) {
	return &s.dogs, nil
}

func (s *MemoryStorage) StoreDog(name string, age int) error {
	dog := &types.Dog{
		ID:   rand.Intn(100),
		Name: name,
		Age:  age,
	}

	dogs := append(s.dogs, *dog)

	s.dogs = dogs

	return nil
}

func (s *MemoryStorage) GetDog(id int) (*types.Dog, error) {
	idx := slices.IndexFunc(s.dogs, func(d types.Dog) bool {
		return d.ID == id
	})

	if idx <= 0 {
		return nil, errors.New("dog not found")
	}

	dog := s.dogs[idx]

	return &dog, nil
}

func (s *MemoryStorage) UpdateDog(newDog *types.Dog) error {
	dog, err := s.GetDog(newDog.ID)
	if err != nil {
		return err
	}

	dog.Name = newDog.Name
	dog.Age = newDog.Age

	err = s.DeleteDog(dog.ID)
	if err != nil {
		return err
	}

	dogs := append(s.dogs, *dog)

	s.dogs = dogs

	return nil
}

func (s *MemoryStorage) DeleteDog(id int) error {
	idx := slices.IndexFunc(s.dogs, func(d types.Dog) bool {
		return d.ID == id
	})

	if idx <= 0 {
		return errors.New("dog not found")
	}

	dogs := append(s.dogs[:idx], s.dogs[idx+1:]...)

	s.dogs = dogs

	return nil
}
