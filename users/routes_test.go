package auth

import (
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

func (c MockedRepo) Create(t common.Entity) (common.Entity, error) {
	return common.User{
		Id: "fake_id",
	}, nil
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
