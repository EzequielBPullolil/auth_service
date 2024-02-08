package users

import (
	"bytes"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/EzequielBPullolil/auth_service/common"
	"github.com/stretchr/testify/assert"
)

var server *http.ServeMux
var url = "/users"

// Simulated repo
type MockedRepo struct {
	common.Repository
}

func (c MockedRepo) Update(t string, e common.Entity) (common.Entity, error) {
	return e, nil

}
func (c MockedRepo) Delete(id string) error {
	return nil
}
func (c MockedRepo) Read(t string) (common.Entity, error) {
	return common.User{
		Id:    "fake_id",
		Name:  "palacios",
		Email: "palacios@gmail.com",
	}, nil
}
func init() {
	server = http.NewServeMux()
	HandleUserRoute(server, MockedRepo{})
}
func TestGetAuthenticatedUser(t *testing.T) {
	req, err := http.NewRequest("GET", url, nil)
	assert.NoError(t, err)

	rr := httptest.NewRecorder()
	server.ServeHTTP(rr, req)

	assert.Equal(t, http.StatusOK, rr.Code)
	assert.Contains(t, rr.Body.String(), `"id": "fake_id"`)
	assert.Contains(t, rr.Body.String(), `"name": "palacios"`)
	assert.Contains(t, rr.Body.String(), `"email": "palacios@gmail.com",`)
}

func TestUpdateUser(t *testing.T) {
	body := bytes.NewReader([]byte(`{
		"name": "new_name",
		"password": "fasdsad",
		"email": "anEmail@gogo.com"
	}`))
	req, err := http.NewRequest("PUT", url, body)
	assert.NoError(t, err)

	rr := httptest.NewRecorder()
	server.ServeHTTP(rr, req)
	assert.Equal(t, http.StatusOK, rr.Code)

	response := rr.Body.String()
	assert.Contains(t, response, `"status": "Successful user update",`)
	assert.Contains(t, response, `"name": "new_name",`)
	assert.Contains(t, response, `"email": "anEmail@gogo.com"`)
}

func TestDeleteUser(t *testing.T) {
	req, err := http.NewRequest("DELETE", url, nil)
	assert.NoError(t, err)
	rr := httptest.NewRecorder()
	server.ServeHTTP(rr, req)
	assert.Equal(t, http.StatusOK, rr.Code)
	response := rr.Body.String()
	assert.Contains(t, response, `"status": "Successful user delete",`)
}

func TestGetUserById(t *testing.T) {
	req, err := http.NewRequest("GET", url+"/fake_id", nil)
	assert.NoError(t, err)
	rr := httptest.NewRecorder()
	server.ServeHTTP(rr, req)
	assert.Equal(t, http.StatusOK, rr.Code)
}
