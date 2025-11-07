package main

import (
	"github.com/gin-gonic/gin"
	"log"
	"user_service/internal/db"
	"user_service/internal/handlers"
	"user_service/internal/middleware"
	"user_service/internal/userService"
)

func main() {
	database, err := db.InitDB()
	if err != nil {
		log.Fatalf("Could not connect to DB: %v", err)
	}

	repoUser := userService.NewUserRepository(database)
	serviceUser := userService.NewUserService(repoUser)
	handlerUser := handlers.NewUserHandler(serviceUser)
	handlerAuth := handlers.NewAuthHandler(serviceUser)

	r := gin.Default()

	r.POST("/login", handlerAuth.LoginHandler)

	users := r.Group("/users")
	users.Use(middleware.JWTMiddleware())
	{
		users.GET("/:id/status", handlerUser.GetUserByIdHandler)
		users.GET("/leaderboard", handlerUser.GetLeaderboardHandler)
		users.POST("/:id/task/complete", handlerUser.PostTaskCompleteHandler)
		users.POST("/:id/referrer", handlerUser.PostReferrerHandler)
	}

	r.Run(":8080")
}
