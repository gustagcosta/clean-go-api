package types

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDogValidateSucess(t *testing.T) {
	dog := &Dog{
		Name: "Gusta",
		Age:  10,
	}

	error := dog.Validate()

	assert.Equal(t, nil, error, "Error must be null")
}

func TestDogValidateNameError(t *testing.T) {
	dog := &Dog{
		Name: "Gu",
		Age:  10,
	}

	error := dog.Validate()

	assert.Equal(t, "name must have at least 3 characters", error.Error(), "Failed at name rule")
}

func TestDogValidateAgeError(t *testing.T) {
	dog := &Dog{
		Name: "Gusta",
		Age:  -2,
	}

	error := dog.Validate()

	assert.Equal(t, "invalid age", error.Error(), "Failed at age rule")
}
