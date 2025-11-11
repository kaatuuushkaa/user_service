package main

import (
	"github.com/gin-gonic/gin"
	"log"
	"os"
	"user_service/internal/db"
	"user_service/internal/handlers"
	"user_service/internal/http"
	"user_service/internal/jwt"
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

	key := os.Getenv("JWT_KEY")
	j := jwt.New(key)
	handlerUser := handlers.NewUserHandler(serviceUser)
	handlerAuth := handlers.NewAuthHandler(serviceUser, j)

	r := gin.Default()
	http.AuthRoutes(r, *handlerAuth)

	r.Use(middleware.JWTMiddleware(j))
	http.UserRoutes(r, *handlerUser)

	if err := r.Run(":8080"); err != nil {
		log.Fatal(err)
	}
}
