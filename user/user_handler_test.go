package user

import (
	"errors"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/labstack/echo"
	"github.com/stretchr/testify/assert"
)

type (
	UserRepositorySuccess struct{}
	UserRepositoryFail    struct{}
)

func (u *UserRepositorySuccess) GetUserByID(id int) (User, error) {
	return User{
		ID:       1,
		Username: "foo",
		Password: "pw",
	}, nil
}

func (u *UserRepositorySuccess) GetAll() ([]User, error) {
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

func (u *UserRepositorySuccess) Create(user User) (int64, error) {
	return 1, nil
}

func (u *UserRepositoryFail) GetUserByID(id int) (User, error) { return User{}, errors.New("") }

func (u *UserRepositoryFail) GetAll() ([]User, error) { return nil, errors.New("") }

func (u *UserRepositoryFail) Create(user User) (int64, error) { return 0, errors.New("") }

func TestGetUserByID(t *testing.T) {
	e := echo.New()
	req := httptest.NewRequest(echo.GET, "/", nil)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	c.SetPath("/users/:id")
	c.SetParamNames("id")
	c.SetParamValues("1")

	u := &UserRepositorySuccess{}
	h := NewHandler(u)

	var userJSON = `{"id":1,"username":"foo","password":"pw"}`

	if assert.NoError(t, h.GetUserByID(c)) {
		assert.Equal(t, http.StatusOK, rec.Code)
		assert.Equal(t, userJSON, rec.Body.String())
	}
}

func TestGetUserByIDErrorShouldReturnHttp500(t *testing.T) {
	e := echo.New()
	req := httptest.NewRequest(echo.GET, "/", nil)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	c.SetPath("/users/:id")
	c.SetParamNames("id")
	c.SetParamValues("1")

	u := &UserRepositoryFail{}
	h := NewHandler(u)

	h.GetUserByID(c)

	assert.Equal(t, http.StatusInternalServerError, rec.Code)
}

func TestGetUserByIDShouldReturnHttp400WhenIdIsEMPTY(t *testing.T) {
	e := echo.New()
	req := httptest.NewRequest(echo.GET, "/", nil)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	c.SetPath("/users/:id")

	u := &UserRepositorySuccess{}
	h := NewHandler(u)

	h.GetUserByID(c)

	assert.Equal(t, http.StatusBadRequest, rec.Code)
}

func TestGetAllUser(t *testing.T) {
	e := echo.New()
	req := httptest.NewRequest(echo.GET, "/", nil)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	c.SetPath("/users")

	u := &UserRepositorySuccess{}
	h := NewHandler(u)

	var userJSON = `[{"id":1,"username":"foo","password":"pw"},{"id":2,"username":"bar","password":"pw"}]`

	if assert.NoError(t, h.GetAll(c)) {
		assert.Equal(t, http.StatusOK, rec.Code)
		assert.Equal(t, userJSON, rec.Body.String())
	}
}

func TestGetAllUserShouldReturnHttp500WhenError(t *testing.T) {
	e := echo.New()
	req := httptest.NewRequest(echo.GET, "/", nil)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	c.SetPath("/users")

	u := &UserRepositoryFail{}
	h := NewHandler(u)
	h.GetAll(c)
	assert.Equal(t, http.StatusInternalServerError, rec.Code)
}

func TestCreateUserHandler(t *testing.T) {
	var userJSON = `{"username":"foo","password":"pw"}`
	c, rec := newTestContextWithPost("/users", userJSON)
	u := &UserRepositorySuccess{}
	h := NewHandler(u)

	if assert.NoError(t, h.Create(c)) {
		assert.Equal(t, http.StatusOK, rec.Code)
		assert.Equal(t, "1", rec.Body.String())
	}
}

func TestCreateUserHandler_expect_internalServerError(t *testing.T) {
	var userJSON = `{"username":"foo","password":"pw"}`
	c, rec := newTestContextWithPost("/users", userJSON)
	u := &UserRepositoryFail{}
	h := NewHandler(u)
	h.Create(c)
	assert.Equal(t, http.StatusInternalServerError, rec.Code)
}

func TestCreateUserHandler_expect_badRequest_when_bind_error(t *testing.T) {
	var userJSON = `{"username":1}`
	c, rec := newTestContextWithPost("/users", userJSON)
	u := &UserRepositorySuccess{}
	h := NewHandler(u)
	h.Create(c)
	assert.Equal(t, http.StatusBadRequest, rec.Code)
}

func newTestContextWithPost(path, jsonBody string) (echo.Context, *httptest.ResponseRecorder) {
	e := echo.New()
	req := httptest.NewRequest(echo.POST, "/", strings.NewReader(jsonBody))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	c.SetPath(path)
	return c, rec
}
