package auth

import (
	"bytes"
	"fmt"
	"net/http"
	"net/http/httptest"
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
func TestAuthSingup(t *testing.T) {
	valid_name := "ezequiel"
	valid_email := "test_email@email.com"
	valid_password := "ValidPassword45#"
	fields := make([]string, 3)
	fields = append(fields, fmt.Sprintf(`{"name":"%s", "password": "%s", "email": "%s"}`, "", valid_password, valid_email))
	fields = append(fields, fmt.Sprintf(`{"name":"%s", "password": "%s", "email": "%s"}`, valid_name, "", valid_email))
	fields = append(fields, fmt.Sprintf(`{"name":"%s", "password": "%s", "email": "%s"}`, valid_name, valid_password, ""))
	for i := range fields {
		t.Run("Should repsonse status 400 if field is empty", func(t *testing.T) {
			rr := httptest.NewRecorder()
			body := bytes.NewReader([]byte(fields[i]))
			req, err := http.NewRequest("POST", url+"/signup", body)
			assert.NoError(t, err)
			server.ServeHTTP(rr, req)
			assert.Equal(t, http.StatusBadRequest, rr.Code)
		})
	}

	t.Run("Should response with status code 201 if all fields are valid", func(t *testing.T) {
		rr := httptest.NewRecorder()
		body := bytes.NewReader([]byte(fmt.Sprintf(`{"name":"%s", "password": "%s", "email": "%s"}`, valid_name, valid_password, valid_email)))
		req, err := http.NewRequest("POST", url+"/signup", body)
		assert.NoError(t, err)
		server.ServeHTTP(rr, req)
		assert.Equal(t, http.StatusCreated, rr.Code)
	})
}

func TestAuthLogin(t *testing.T) {
	expectedToken := fmt.Sprintf(`"token": "%s"`, tokenmanager.CreateToken("anEmail@gogo.com"))
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

		cookie.Value = tokenmanager.CreateToken("fake_id")
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
