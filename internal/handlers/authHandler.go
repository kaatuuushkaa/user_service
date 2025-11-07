package handlers

import (
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"net/http"
	"strconv"
	"time"
	"user_service/internal/middleware"
	"user_service/internal/userService"
)

type AuthHandler struct {
	service userService.UserService
}

func NewAuthHandler(service userService.UserService) *AuthHandler {
	return &AuthHandler{service}
}

func (h *AuthHandler) LoginHandler(c *gin.Context) {
	var body struct {
		UserID int `json:"user_id"`
	}
	if err := c.ShouldBindJSON(&body); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}

	_, err := h.service.GetUserById(strconv.Itoa(body.UserID))
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "user not found"})
		return
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"user_id": body.UserID,
		"exp":     jwt.NewNumericDate(time.Now().Add(time.Hour * 24)),
	})

	tokenString, err := token.SignedString(middleware.JwtSecret)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "could not create token"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"access_token": tokenString,
	})
}
