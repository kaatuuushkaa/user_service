package handlers

import (
	"github.com/labstack/echo/v4"
	"net/http"
	"user_service/internal/userService"
)

type UserHandler struct {
	service userService.UserService
}

func NewUserHandler(service userService.UserService) *UserHandler {
	return &UserHandler{service}
}

func (h *UserHandler) GetUserById(c echo.Context) error {
	//idStr := c.Param("id")
	//id, err := strconv.Atoi(idStr)
	//if err != nil {
	//	return echo.NewHTTPError(http.StatusBadRequest, "Invalid user id")
	//}

	user, err := h.service.GetUserById(c.Param("id"))
	if err != nil {
		return echo.NewHTTPError(http.StatusNotFound, err.Error())
	}

	return c.JSON(http.StatusOK, user)
}

func (h *UserHandler) GetLeaderboard(c echo.Context) error {
	users, err := h.service.GetLeaderboard()
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, users)
}
