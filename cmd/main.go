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

	r := gin.Default()

	users := r.Group("/users")
	users.Use(middleware.JWTMiddleware())
	{
		users.GET("/:id/status", handlerUser.GetUserById)
		users.GET("/leaderboard", handlerUser.GetLeaderboard)
	}

	r.Run(":8080")

	//e := echo.New()
	//
	//e.Use(middleware.Logger())
	//e.Use(middleware.Recover())
	//
	//usersGroup := e.Group("/users")
	//usersGroup.Use(middleware.JWTWithConfig(middleware.JWTConfig{
	//	SigningKey: middleware.JwtSecret,
	//}))
	//
	//usersGroup.GET("/:id/status", handlerUser.GetUserById)
	//usersGroup.GET("/leaderboard", handlerUser.GetLeaderboard)
	//
	//if err := e.Start(":8080"); err != nil {
	//	log.Fatalf("Failed to start with err: %v", err)
	//}
}
