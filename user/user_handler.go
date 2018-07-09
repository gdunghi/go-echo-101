package user

import (
	"net/http"
	"strconv"

	"github.com/labstack/echo"
)

// Handler ... for test
type Handler struct {
	UserRepositoryInterface
}

//NewHandler ... return handler
func NewHandler(u UserRepositoryInterface) *Handler {
	return &Handler{u}
}

//GetUserByID ... GetUserByID
func (h *Handler) GetUserByID(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, "")
	}
	u, err := h.UserRepositoryInterface.GetUserByID(id)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, "")
	}
	return c.JSON(http.StatusOK, u)
}

//GetAll ... GetAll
func (h *Handler) GetAll(c echo.Context) error {
	u, err := h.UserRepositoryInterface.GetAll()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, "")
	}
	return c.JSON(http.StatusOK, u)
}

//Create ... Create Handler
func (h *Handler) Create(c echo.Context) error {

	u := new(User)
	if err := c.Bind(u); err != nil {
		return c.JSON(http.StatusBadRequest, "")
	}
	id, err := h.UserRepositoryInterface.Create(*u)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, "")
	}
	return c.JSON(http.StatusOK, id)

}
