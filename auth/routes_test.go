package auth

import (
	"bytes"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/EzequielBPullolil/auth_service/common"
	"github.com/stretchr/testify/assert"
)

var server *http.ServeMux
var url = "/auth"

// Simulated repo
type MockedRepo struct {
	common.Repository
}

func (c MockedRepo) Create(t common.Entity) (common.Entity, error) {
	return common.User{
		Id: "fake_id",
	}, nil
}
func (c MockedRepo) Read(t string) (common.Entity, error) {
	return common.User{
		Id: "fake_id",
	}, nil
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
	assert.Contains(t, rr.Body.String(), `"id": "fake_id"`)
	assert.Contains(t, rr.Body.String(), `"name": "ezequiel"`)
	assert.Contains(t, rr.Body.String(), `"email": "anEmail@gogo.com",`)
}

func TestAuthLogin(t *testing.T) {
	expectedToken := fmt.Sprintf(`"token": "%s"`, CreateToken("fake_id"))
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
	assert.Contains(t, rr.Body.String(), expectedToken)
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
