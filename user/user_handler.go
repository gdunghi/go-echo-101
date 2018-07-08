package user

import (
	"net/http"
	"strconv"

	"github.com/labstack/echo"
)

// Handler ... for test
type Handler struct {
	UserModelInterface
}

//NewHandler ... return handler
func NewHandler(u UserModelInterface) *Handler {
	return &Handler{u}
}

//GetUserByID ... GetUserByID
func (h *Handler) GetUserByID(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	u := h.UserModelInterface.GetUserByID(id)
	return c.JSON(http.StatusOK, u)
}

//GetAll ... GetAll
func (h *Handler) GetAll(c echo.Context) error {
	u, err := h.UserModelInterface.GetAll()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, "")
	}
	return c.JSON(http.StatusOK, u)
}
