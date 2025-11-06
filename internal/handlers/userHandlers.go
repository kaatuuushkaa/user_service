package handlers

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
	"user_service/internal/userService"
)

type UserHandler struct {
	service userService.UserService
}

func NewUserHandler(service userService.UserService) *UserHandler {
	return &UserHandler{service}
}

func (h *UserHandler) GetUserById(c *gin.Context) {
	userIDFromToken := c.GetInt("user_id")
	requestedID := c.Param("id")

	if strconv.Itoa(userIDFromToken) != requestedID {
		c.JSON(http.StatusForbidden, gin.H{"error": "access denied"})
		return
	}
	user, err := h.service.GetUserById(requestedID)
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
