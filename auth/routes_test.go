package auth

import (
	"bytes"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/EzequielBPullolil/auth_service/common"
	"github.com/stretchr/testify/assert"
)

var server *http.ServeMux
var url = "/auth"

type MockedRepo struct {
	common.Repository
}

func init() {
	server = http.NewServeMux()
	HandleAuthRoutes(server, MockedRepo{})
}
func TestAuthSingup(t *testing.T) {
	body := bytes.NewReader([]byte(`{
		"name":"ezequiel",
		"email":"anEmail@gogo.com",
		"password": "original_password"
	}`))
	req, err := http.NewRequest("POST", url+"/singup", body)
	assert.NoError(t, err)

	rr := httptest.NewRecorder()
	server.ServeHTTP(rr, req)

	assert.Equal(t, http.StatusCreated, rr.Code)
	assert.Contains(t, rr.Body.String(), `"status": "Successful user registration",`)
	assert.Contains(t, rr.Body.String(), `"name": "ezequiel"`)
	assert.Contains(t, rr.Body.String(), `"email": "anEmail@gogo.com",`)
}

func TestAuthLogin(t *testing.T) {
	body := bytes.NewReader([]byte(`{
		"email":"anEmail@gogo.com",
		"password": "original_password"
	}`))
	req, err := http.NewRequest("POST", url+"/login", body)
	assert.NoError(t, err)

	rr := httptest.NewRecorder()
	server.ServeHTTP(rr, req)
	assert.Equal(t, http.StatusCreated, rr.Code)

	assert.Contains(t, rr.Body.String(), `"status": "Successful user login",`)
	assert.Contains(t, rr.Body.String(), `"token": "fake_token"`)
	assert.Contains(t, rr.Body.String(), `"email": "anEmail@gogo.com",`)
}

func TestAuthValidate(t *testing.T) {
	req, err := http.NewRequest("POST", url+"/validate", nil)
	assert.NoError(t, err)

	rr := httptest.NewRecorder()
	server.ServeHTTP(rr, req)
	assert.Equal(t, http.StatusOK, rr.Code)
	assert.Contains(t, rr.Body.String(), `"status": "Valid auth token",`)
}
