package handlers

import (
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"net/http"
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

func (h *AuthHandler) RegisterHandler(c *gin.Context) {
	var body struct {
		Username string `json:"Username"`
		Password string `json:"Password"`
	}
	if err := c.ShouldBindJSON(&body); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}

	user, err := h.service.Register(body.Username, body.Password)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"message":  "User registered successfully",
		"username": user.Username,
		"user_id":  user.ID,
	})
}

func (h *AuthHandler) LoginHandler(c *gin.Context) {
	var body struct {
		Username string `json:"Username"`
		Password string `json:"Password"`
	}
	if err := c.ShouldBindJSON(&body); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}

	user, err := h.service.Login(body.Username, body.Password)
	if err != nil {
		switch err.Error() {
		case "user not found":
			c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
		case "invalid password":
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid password"})
		default:
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal server error"})
		}
		return
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"user_id": user.ID,
		"exp":     jwt.NewNumericDate(time.Now().Add(time.Hour * 24)),
	})

	tokenString, err := token.SignedString(middleware.JwtSecret)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Could not create token"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message":      "Login successful",
		"access_token": tokenString,
	})
}
