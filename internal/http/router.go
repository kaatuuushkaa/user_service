package http

import (
	"github.com/gin-gonic/gin"
	"user_service/internal/handlers"
)

func AuthRoutes(r *gin.Engine, handler handlers.AuthHandler) {
	auth := r.Group("/auth")

	auth.POST("/signin", handler.SignInHandler)
	auth.POST("/signup", handler.SignUpHandler)
}

func UserRoutes(r *gin.Engine, handler handlers.UserHandler) {
	user := r.Group("/users")

	user.GET("/:id/status", handler.GetUserByIdHandler)
	user.GET("/leaderboard", handler.GetLeaderboardHandler)
	user.POST("/:id/task/complete", handler.PostTaskCompleteHandler)
	user.POST("/:id/referrer", handler.PostReferrerHandler)
}
