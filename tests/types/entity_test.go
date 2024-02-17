package types

import (
	"testing"

	"github.com/EzequielBPullolil/auth_service/src/types"
	"github.com/stretchr/testify/assert"
)

func TestValidateName(t *testing.T) {
	var user_suject = types.User{}
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
	var user_suject = types.User{}
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
func TestValidatePassword(t *testing.T) {
	var user_suject = types.User{}
	var invalid_cases = []struct{ Name, Value string }{
		{Name: "A password with less than 8 characters should be invalid", Value: "Aa#2bca"},
		{Name: "A password without a number should be invalid", Value: "Aaaaaaa#"},
		{Name: "A password without a symbol should be invalid", Value: "Aaaaaaa3"},
		{Name: "A password without a capital letter should be invalid", Value: "a#4aaaaa"},
		{Name: "A password without a lowercase should be invalid", Value: "A#4AAAAA"},
	}

	for _, field := range invalid_cases {

		t.Run(field.Name, func(t *testing.T) {
			user_suject.Password = field.Value
			assert.False(t, user_suject.ValidatePassword())
		})
	}

	t.Run("Should be valid Password", func(t *testing.T) {
		user_suject.Password = "Mip@ssw0rd"
		assert.True(t, user_suject.ValidatePassword())
	})
}
