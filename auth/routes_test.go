package auth

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

var server *http.ServeMux
var url = "/auth"

func init() {
	server = http.NewServeMux()
	HandleAuthRoutes(server)
}
func TestAuthSingup(t *testing.T) {
	req, err := http.NewRequest("POST", url+"/singup", http.NoBody)
	assert.NoError(t, err)

	rr := httptest.NewRecorder()
	server.ServeHTTP(rr, req)

	assert.Equal(t, http.StatusCreated, rr.Code)
	assert.Contains(t, "status", rr.Body.String())
}
