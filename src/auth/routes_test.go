package auth

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	tokenmanager "github.com/EzequielBPullolil/auth_service/src/token_manager"
	"github.com/EzequielBPullolil/auth_service/src/types"
	"github.com/stretchr/testify/assert"
)

var server *http.ServeMux
var url = "/auth"

// Simulated repo
type MockedRepo struct {
	types.Repository
}

func (c MockedRepo) Create(t types.User) (types.User, error) {
	t.Id = "test_id_fake"
	t.HashPassword()
	return t, nil
}
func (c MockedRepo) Read(t string) (*types.User, error) {
	return &types.User{
		Id:    "fake_id",
		Name:  "ezequiel",
		Email: "anEmail@gogo.com",
	}, nil
}
func init() {
	server = http.NewServeMux()
	HandleAuthRoutes(server, MockedRepo{})
}
func createBody(name, password, email string) string {
	return fmt.Sprintf(`{"name":"%s", "password":"%s", "email":"%s"}`, name, password, email)
}
func TestAuthSingup(t *testing.T) {
	endpoint := url + "/signup"
	valid_name := "ezequiel"
	valid_email := "test_email@email.com"
	valid_password := "ValidPassword45#"
	t.Run("Should be bad repsonse if", func(t *testing.T) {
		t.Run("Any field is empty", func(t *testing.T) {
			fields := make([]string, 3)
			fields = append(fields, createBody("", valid_password, valid_email))
			fields = append(fields, createBody(valid_name, "", valid_email))
			fields = append(fields, createBody(valid_name, valid_password, ""))
			for i := range fields {
				t.Run("status code 400", func(t *testing.T) {
					rr := httptest.NewRecorder()
					body := bytes.NewReader([]byte(fields[i]))
					req, err := http.NewRequest("POST", endpoint, body)
					assert.NoError(t, err)
					server.ServeHTTP(rr, req)
					assert.Equal(t, http.StatusBadRequest, rr.Code)
				})
			}
		})
		t.Run("Try signup user with invalid fields", func(t *testing.T) {
			var invalid_fields = [][]struct{ Name, Body string }{
				{
					{"it's not long enough", createBody("abcd", valid_password, valid_email)},
					{"it's not empty", createBody("", valid_password, valid_email)},
					{"Not contains numbers", createBody("abcdf#", valid_password, valid_email)},
					{"Not contains symbols", createBody("abcdf5", valid_password, valid_email)},
				},
				{
					{"it's not long enough", createBody(valid_name, "Abcde#2", valid_email)},
					{"it's empty", createBody(valid_name, "", valid_email)},
					{"Not contains numbers", createBody(valid_name, "Abcdf#A", valid_email)},
					{"Not contains symbols", createBody(valid_name, "Abcde22", valid_email)},
				},
			}
			for i := range invalid_fields {
				for _, field := range invalid_fields[i] {
					t.Run(field.Name, func(t *testing.T) {
						rr := httptest.NewRecorder()
						body := bytes.NewReader([]byte(field.Body))
						req, err := http.NewRequest("POST", endpoint, body)
						assert.NoError(t, err)
						server.ServeHTTP(rr, req)
						log.Println(rr.Body.String())
						assert.Equal(t, http.StatusBadRequest, rr.Code)
					})
				}
			}
		})
	})

	t.Run("Should response with status code 201 if all fields are valid", func(t *testing.T) {
		expected_response, _ := json.Marshal(types.ResponseWithData{
			Status: "Succesful user registration",
			Data: types.UserDAO{
				User: types.User{
					Id:       "test_id_fake",
					Name:     valid_name,
					Email:    valid_email,
					Password: types.HashPassword(valid_password),
				},
			},
		})
		rr := httptest.NewRecorder()
		body := bytes.NewReader([]byte(createBody(valid_name, valid_password, valid_email)))
		req, err := http.NewRequest("POST", endpoint, body)
		assert.NoError(t, err)
		server.ServeHTTP(rr, req)

		assert.Equal(t, http.StatusCreated, rr.Code)

		response := strings.TrimSuffix(rr.Body.String(), "\n")
		assert.Equal(t, string(expected_response), response)
	})
}

func TestAuthLogin(t *testing.T) {
	expected_response, _ := json.Marshal(types.ResponseWithData{
		Status: "Successful user login",
		Data: types.TokenData{
			Token: tokenmanager.CreateToken("anEmail@gogo.com"),
			User: types.User{
				Id:    "fake_id",
				Name:  "ezequiel",
				Email: "anEmail@gogo.com",
			},
		},
	})
	endpoint := url + "/login"
	body := bytes.NewReader([]byte(`{
		"email": "anEmail@gogo.com",
		"password": "original_password"
	}`))
	req, err := http.NewRequest("POST", endpoint, body)
	assert.NoError(t, err)

	rr := httptest.NewRecorder()
	server.ServeHTTP(rr, req)
	assert.Equal(t, http.StatusCreated, rr.Code)

	response := strings.TrimSuffix(rr.Body.String(), "\n")
	assert.Equal(t, string(expected_response), response)
}

func TestAuthValidate(t *testing.T) {
	endpoint := url + "/validate"
	cookie := &http.Cookie{
		Name: "auth_token",
	}

	t.Run("Should response with status code 200 if the auth_token is valid", func(t *testing.T) {
		expected_response, _ := json.Marshal(types.ResponseWithData{
			Status: "Valid auth token",
			Data:   struct{}{},
		})
		rr := httptest.NewRecorder()
		req, err := http.NewRequest("POST", endpoint, nil)
		assert.NoError(t, err)

		cookie.Value = tokenmanager.CreateToken("fake_id")
		req.AddCookie(cookie)
		assert.NotEmpty(t, req.Cookies())

		server.ServeHTTP(rr, req)
		assert.Equal(t, http.StatusOK, rr.Code)
		response := strings.TrimSuffix(rr.Body.String(), "\n")
		assert.Equal(t, string(expected_response), response)
	})
	t.Run("Should response with status code 400 if the auth_token is invalid", func(t *testing.T) {
		expected_response, _ := json.Marshal(types.ResponseError{
			Status: "Invalid auth token",
			Error:  "",
		})

		rr := httptest.NewRecorder()
		req, err := http.NewRequest("POST", endpoint, nil)
		assert.NoError(t, err)

		cookie.Value = "a fake token"
		req.AddCookie(cookie)
		assert.NotEmpty(t, req.Cookies())
		server.ServeHTTP(rr, req)

		response := strings.TrimSuffix(rr.Body.String(), "\n")
		assert.Equal(t, http.StatusBadRequest, rr.Code)
		assert.Equal(t, string(expected_response), response)
	})

}
