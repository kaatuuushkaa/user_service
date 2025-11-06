package main

import (
	"fmt"
	"log"
	"user_service/internal/db"
)

func main() {
	database, err := db.InitDB()
	if err != nil {
		log.Fatalf("Could not connect to DB: %v", err)
	}
	fmt.Println(database)
}
