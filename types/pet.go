package types

import "errors"

type Dog struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
	Age  int    `json:"age"`
}

type DogStoreRequest struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
}

func (d *Dog) Validate() error {
	if d.Age <= 0 {
		return errors.New("invalid age")
	}

	if len(d.Name) < 3 {
		return errors.New("name must have at least 3 characters")
	}

	return nil
}
