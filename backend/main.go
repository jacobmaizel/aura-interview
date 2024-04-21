package main

import (
	"log"
	"os"

	"github.com/jacobmaizel/aura/backend/server"
)

func main() {

	port := os.Getenv("PORT")

	if port == "" {
		log.Fatal("PORT Env var is missing")
	}

	server.Pool = Init_db()

	server.Http = Init_gin()

	err := LoadData()

	if err != nil {
		log.Println("Error loading data:", err)
	}

	server.Http.Run(":" + port)

}
