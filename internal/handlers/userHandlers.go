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

func (h *UserHandler) GetUserByIdHandler(c *gin.Context) {
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

func (h *UserHandler) GetLeaderboardHandler(c *gin.Context) {
	users, err := h.service.GetLeaderboard()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, users)
}

func (h *UserHandler) PostTaskCompleteHandler(c *gin.Context) {
	user, err := h.service.PostTaskComplete(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, user)
}

func (h *UserHandler) PostReferrerHandler(c *gin.Context) {
	var body struct {
		ReferrerID int `json:"referrer_id"`
	}
	if err := c.ShouldBindJSON(&body); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}
	users, err := h.service.PostReferrerHandler(c.Param("id"), strconv.Itoa(body.ReferrerID))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, users)
}
