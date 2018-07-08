package user

import (
	"net/http"
	"strconv"

	"github.com/labstack/echo"
)

type handler struct {
	UserModelInterface
}

func NewHandler(u UserModelInterface) *handler {
	return &handler{u}
}

func (h *handler) GetUserByID(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	u := h.UserModelInterface.GetUserByID(id)
	return c.JSON(http.StatusOK, u)
}

func (h *handler) GetAllUsers(c echo.Context) error {
	u, err := h.UserModelInterface.GetAll()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, "")
	}
	return c.JSON(http.StatusOK, u)
}
