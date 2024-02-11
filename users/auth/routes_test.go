package auth

import (
	"bytes"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/EzequielBPullolil/auth_service/users"
	"github.com/stretchr/testify/assert"
)

var server *http.ServeMux
var url = "/auth"

// Simulated repo
type MockedRepo struct {
	users.Repository
}

func (c MockedRepo) Create(t users.User) (users.User, error) {
	return users.User{
		Id:    "fake_id",
		Email: "anEmail@gogo.com",
		Name:  "ezequiel",
	}, nil
}
func (c MockedRepo) Read(t string) (*users.User, error) {
	return &users.User{
		Id:    "fake_id",
		Name:  "ezequiel",
		Email: "anEmail@gogo.com",
	}, nil
}
func init() {
	server = http.NewServeMux()
	HandleAuthRoutes(server, MockedRepo{})
}
func TestAuthSingup(t *testing.T) {
	body := bytes.NewReader([]byte(`{
		"name": "ezequiel",
		"email": "anEmail@gogo.com",
		"password": "original_password"
	}`))
	req, err := http.NewRequest("POST", url+"/singup", body)
	assert.NoError(t, err)

	rr := httptest.NewRecorder()
	server.ServeHTTP(rr, req)

	assert.Equal(t, http.StatusCreated, rr.Code)
	assert.Contains(t, rr.Body.String(), `"status": "Successful user registration",`)
	assert.Contains(t, rr.Body.String(), `"id": "fake_id"`)
	assert.Contains(t, rr.Body.String(), `"name": "ezequiel"`)
	assert.Contains(t, rr.Body.String(), `"email": "anEmail@gogo.com",`)
}

func TestAuthLogin(t *testing.T) {
	expectedToken := fmt.Sprintf(`"token": "%s"`, users.CreateToken("fake_id"))
	body := bytes.NewReader([]byte(`{
		"email": "anEmail@gogo.com",
		"password": "original_password"
	}`))
	req, err := http.NewRequest("POST", url+"/login", body)
	assert.NoError(t, err)

	rr := httptest.NewRecorder()
	server.ServeHTTP(rr, req)
	assert.Equal(t, http.StatusCreated, rr.Code)

	assert.Contains(t, rr.Body.String(), `"status": "Successful user login",`)
	assert.Contains(t, rr.Body.String(), expectedToken)
	assert.Contains(t, rr.Body.String(), `"email": "anEmail@gogo.com",`)
}

func TestAuthValidate(t *testing.T) {
	endpoint := url + "/validate"
	cookie := &http.Cookie{
		Name: "auth_token",
	}

	t.Run("Should response valid auth_token", func(t *testing.T) {
		rr := httptest.NewRecorder()
		req, err := http.NewRequest("POST", endpoint, nil)
		assert.NoError(t, err)

		cookie.Value = users.CreateToken("fake_id")
		req.AddCookie(cookie)
		assert.NotEmpty(t, req.Cookies())

		server.ServeHTTP(rr, req)
		assert.Equal(t, http.StatusOK, rr.Code)
		assert.Contains(t, rr.Body.String(), `"status": "Valid auth token",`)
	})
	t.Run("Should be invalid response if auth_token is invalid", func(t *testing.T) {
		rr := httptest.NewRecorder()
		req, err := http.NewRequest("POST", endpoint, nil)
		assert.NoError(t, err)

		cookie.Value = "a fake token"
		req.AddCookie(cookie)
		assert.NotEmpty(t, req.Cookies())

		server.ServeHTTP(rr, req)
		assert.Equal(t, http.StatusBadRequest, rr.Code)
		assert.Contains(t, rr.Body.String(), `"status": "Invalid auth token",`)
	})

}
