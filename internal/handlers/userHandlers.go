package handlers

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"user_service/internal/userService"
)

type UserHandler struct {
	service userService.UserService
}

func NewUserHandler(service userService.UserService) *UserHandler {
	return &UserHandler{service}
}

func (h *UserHandler) GetUserById(c *gin.Context) {
	user, err := h.service.GetUserById(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, user)
}

func (h *UserHandler) GetLeaderboard(c *gin.Context) {
	users, err := h.service.GetLeaderboard()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, users)
}
