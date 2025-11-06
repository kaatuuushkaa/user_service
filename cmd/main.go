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

	repoUser := userService.NewUserRepositiry(database)
	serviceUser := userService.NewUserService(repoUser)
	handlerUser := handlers.NewUserHandler(serviceUser)
	handlerAuth := handlers.NewAuthHandler()

	r := gin.Default()

	r.POST("/login", handlerAuth.Login)

	auth := r.Group("/users")
	auth.Use(middleware.JWTMiddleware())
	{
		auth.GET("/:id/status", handlerUser.GetUserById)
		auth.GET("/leaderboard", handlerUser.GetLeaderboard)
	}

	r.Run(":8080")
}
