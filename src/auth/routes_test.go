package auth

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	tokenmanager "github.com/EzequielBPullolil/auth_service/src/token_manager"
	"github.com/EzequielBPullolil/auth_service/src/types"
	"github.com/google/uuid"
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

		t.Run("Any field is invalid", func(t *testing.T) {
			var invalid_cases = []struct{ Title, Error, Body string }{
				{Title: "Invalid name", Error: "Invalid Name", Body: createBody("Abcd#", valid_password, valid_email)},
				{Title: "Invalid password", Error: "Invalid Password", Body: createBody(valid_name, "Abc#3aa", valid_email)},
				{Title: "Invalid email", Error: "Invalid Email", Body: createBody(valid_name, valid_password, "nodomain.com")},
			}

			for _, c := range invalid_cases {
				t.Run(c.Title, func(t *testing.T) {
					expected_response, _ := json.Marshal(types.ResponseError{
						Status: "error signup user",
						Error:  c.Error,
					})
					rr := httptest.NewRecorder()
					body := bytes.NewReader([]byte(c.Body))
					req, err := http.NewRequest("POST", endpoint, body)
					assert.NoError(t, err)
					server.ServeHTTP(rr, req)
					assert.Equal(t, http.StatusBadRequest, rr.Code)

					response := strings.TrimSuffix(rr.Body.String(), "\n")
					assert.Equal(t, string(expected_response), response)
				})
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
	token, _ := tokenmanager.CreateToken(types.User{
		Id:    "fake_id",
		Email: "anEmail@gogo.com",
	})
	expected_response, _ := json.Marshal(types.ResponseWithData{
		Status: "Successful user login",
		Data: types.TokenData{
			Token: token,
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
		token, _ := tokenmanager.CreateToken(types.User{
			Id:    uuid.New().String(),
			Email: "email@test.com",
		})
		expected_response, _ := json.Marshal(types.ResponseWithData{
			Status: "Valid auth token",
			Data:   struct{}{},
		})
		rr := httptest.NewRecorder()
		req, err := http.NewRequest("POST", endpoint, nil)
		assert.NoError(t, err)

		cookie.Value = token
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
