package types

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestValidateName(t *testing.T) {
	var user_suject = User{}
	var invalid_cases = []struct{ Name, Value string }{
		{Name: "A name with symbols should not be valid", Value: "Abcdf#"},
		{Name: "A name with numbers should not be valid", Value: "Abcdf3"},
		{Name: "A name with less than 6 characters should not be valid.", Value: "Abcdf"},
	}

	for _, field := range invalid_cases {

		t.Run(field.Name, func(t *testing.T) {
			user_suject.Name = field.Value
			assert.False(t, user_suject.ValidateName())
		})
	}

	t.Run("Should be valid name", func(t *testing.T) {
		user_suject.Name = "Ezequiel"
		assert.True(t, user_suject.ValidateName())
	})
}

func TestValidateEmail(t *testing.T) {
	var user_suject = User{}
	var invalid_cases = []struct{ Name, Value string }{
		{Name: "An email without a name should be invalid", Value: "@domain.com"},
		{Name: "An email without a @ should be invalid", Value: "namedomain.com"},
		{Name: "An email without a domain should be invalid", Value: "name@.com"},
		{Name: "An email without a dot should be invalid", Value: "ezequiel@email com"},
	}

	for _, field := range invalid_cases {

		t.Run(field.Name, func(t *testing.T) {
			user_suject.Email = field.Value
			assert.False(t, user_suject.ValidateEmaiL())
		})
	}

	t.Run("Should be valid Email", func(t *testing.T) {
		user_suject.Email = "name@domain.org"
		assert.True(t, user_suject.ValidateEmaiL())
	})
}
