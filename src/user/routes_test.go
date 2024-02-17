package user

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	tokenmanager "github.com/EzequielBPullolil/auth_service/src/token_manager"
	"github.com/EzequielBPullolil/auth_service/src/types"
	"github.com/stretchr/testify/assert"
)

var server *http.ServeMux
var url = "/users"

// Simulated repo
type MockedRepo struct {
	types.Repository
}

func (c MockedRepo) Update(t string, e types.User) (*types.User, error) {
	e.Id = "fake_id"
	e.HashPassword()
	return &e, nil

}
func (c MockedRepo) Delete(id string) error {
	return nil
}
func (c MockedRepo) Read(t string) (*types.User, error) {
	return &types.User{
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
	token, _ := tokenmanager.CreateToken(types.User{})
	expected_response, _ := json.Marshal(types.ResponseWithData{
		Status: "Successful user find",
		Data: types.UserDAO{
			User: types.User{
				Id:       "fake_id",
				Name:     "palacios",
				Email:    "palacios@gmail.com",
				Password: "",
			},
		},
	})
	req, err := http.NewRequest("GET", url, nil)
	req.AddCookie(&http.Cookie{
		Name:  "auth_token",
		Value: token,
	})
	assert.NoError(t, err)

	rr := httptest.NewRecorder()
	server.ServeHTTP(rr, req)

	assert.Equal(t, http.StatusOK, rr.Code)
	response := strings.TrimSuffix(rr.Body.String(), "\n")
	assert.Equal(t, string(expected_response), response)
}

func TestUpdateUser(t *testing.T) {
	token, _ := tokenmanager.CreateToken(types.User{})
	expected_response, _ := json.Marshal(types.ResponseWithData{
		Status: "Successful user update",
		Data: types.UserDAO{
			User: types.User{
				Id:       "fake_id",
				Name:     "new_name",
				Email:    "anEmail@gogo.com",
				Password: types.HashPassword("fasdsad"),
			},
		},
	})
	body := bytes.NewReader([]byte(`{
		"name": "new_name",
		"password": "fasdsad",
		"email": "anEmail@gogo.com"
	}`))
	req, err := http.NewRequest("PUT", url, body)
	req.AddCookie(&http.Cookie{
		Name:  "auth_token",
		Value: token,
	})
	assert.NoError(t, err)

	rr := httptest.NewRecorder()
	server.ServeHTTP(rr, req)
	assert.Equal(t, http.StatusOK, rr.Code)

	response := strings.TrimSuffix(rr.Body.String(), "\n")
	assert.Equal(t, string(expected_response), response)
}

func TestDeleteUser(t *testing.T) {
	token, _ := tokenmanager.CreateToken(types.User{})
	expected_response, _ := json.Marshal(types.ResponseWithData{
		Status: "Successful user delete",
		Data:   struct{}{},
	})
	req, err := http.NewRequest("DELETE", url, nil)
	req.AddCookie(&http.Cookie{
		Name:  "auth_token",
		Value: token,
	})
	assert.NoError(t, err)
	rr := httptest.NewRecorder()
	server.ServeHTTP(rr, req)
	assert.Equal(t, http.StatusOK, rr.Code)
	response := strings.TrimSuffix(rr.Body.String(), "\n")
	assert.Equal(t, string(expected_response), response)
}

func TestGetUserById(t *testing.T) {
	req, err := http.NewRequest("GET", url+"/fake_id", nil)
	assert.NoError(t, err)
	rr := httptest.NewRecorder()
	server.ServeHTTP(rr, req)
	assert.Equal(t, http.StatusOK, rr.Code)
}
