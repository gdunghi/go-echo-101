package user

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/labstack/echo"
	"github.com/stretchr/testify/assert"
)

type (
	UsersModelStub struct{}
)

func (u *UsersModelStub) GetUserByID(id int) User {
	return User{
		ID:       1,
		Username: "foo",
		Password: "pw",
	}
}

func (u *UsersModelStub) GetAll() ([]User, error) {
	return []User{{
		ID:       1,
		Username: "foo",
		Password: "pw",
	}, {
		ID:       2,
		Username: "bar",
		Password: "pw",
	}}, nil
}

func TestGetUserByID(t *testing.T) {
	e := echo.New()
	req := httptest.NewRequest(echo.GET, "/", nil)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	c.SetPath("/users/:id")
	c.SetParamNames("id")
	c.SetParamValues("1")

	u := &UsersModelStub{}
	h := NewHandler(u)

	var userJSON = `{"id":1,"username":"foo","password":"pw"}`

	if assert.NoError(t, h.GetUserByID(c)) {
		assert.Equal(t, http.StatusOK, rec.Code)
		assert.Equal(t, userJSON, rec.Body.String())
	}
}

func TestGetAllUser(t *testing.T) {
	e := echo.New()
	req := httptest.NewRequest(echo.GET, "/", nil)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	c.SetPath("/users")

	u := &UsersModelStub{}
	h := NewHandler(u)

	var userJSON = `[{"id":1,"username":"foo","password":"pw"},{"id":2,"username":"bar","password":"pw"}]`

	if assert.NoError(t, h.GetAll(c)) {
		assert.Equal(t, http.StatusOK, rec.Code)
		assert.Equal(t, userJSON, rec.Body.String())
	}
}
