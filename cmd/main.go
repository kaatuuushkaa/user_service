package main

import (
	echo "github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"log"
	"user_service/internal/db"
	"user_service/internal/handlers"
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

	e := echo.New()

	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	e.GET("/users/:id/status", handlerUser.GetUserById)
	e.GET("/users/leaderboard", handlerUser.GetLeaderboard)

	if err := e.Start(":8080"); err != nil {
		log.Fatalf("Failed to start with err: %v", err)
	}
}
