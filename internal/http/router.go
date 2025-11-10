package http

import (
	"github.com/gin-gonic/gin"
	"user_service/internal/handlers"
)

func AuthRoutes(r *gin.Engine, handler handlers.AuthHandler) {
	auth := r.Group("/auth")

	auth.POST("/login", handler.LoginHandler)
	auth.POST("/register", handler.RegisterHandler)
}

func UserRoutes(r *gin.Engine, handler handlers.UserHandler) {
	user := r.Group("/users")

	user.GET("/:id/status", handler.GetUserByIdHandler)
	user.GET("/leaderboard", handler.GetLeaderboardHandler)
	user.POST("/:id/task/complete", handler.PostTaskCompleteHandler)
	user.POST("/:id/referrer", handler.PostReferrerHandler)
}
