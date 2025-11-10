package handlers

import (
	"github.com/gin-gonic/gin"
	appjwt "user_service/internal/jwt"

	//"github.com/golang-jwt/jwt/v5"
	"net/http"
	"time"
	//"user_service/internal/jwt"
	"user_service/internal/userService"
)

type AuthHandler struct {
	service userService.UserService
	jwt     appjwt.IJWT
}

func NewAuthHandler(service userService.UserService, jwt appjwt.IJWT) *AuthHandler {
	return &AuthHandler{service: service, jwt: jwt}
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
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid username or password"})
		return
	}

	accessToken := h.jwt.GenerateJWT(user.ID, true, appjwt.Minute*10)
	refreshToken, expAfter := h.jwt.GenerateRefreshToken(user.ID, true, appjwt.SixMonth)

	c.SetCookie("REFRESH_TOKEN", refreshToken, int(expAfter.Sub(time.Now()).Seconds()), "/", "", false, true)

	c.JSON(http.StatusOK, gin.H{
		"message":      "Login successful",
		"access_token": accessToken,
	})
}
