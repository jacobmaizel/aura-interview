package main

import (
	"log"
	"os"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func Init_gin() *gin.Engine {
	var ginEngine *gin.Engine

	env := os.Getenv("ENVIRONMENT")
	if env == "" {
		log.Fatal("MISSING ENVIRONMENT in env vars")
	}

	switch env {
	case "DEVELOPMENT":
		ginEngine = gin.Default()
		corsConfig := cors.DefaultConfig()

		corsConfig.AllowAllOrigins = true

		ginEngine.Use(cors.New(corsConfig))
	case "PRODUCTION":
		ginEngine = gin.New()
	default:
		log.Fatal("Invalid ENVIRONMENT env var set, needs DEVELOPMENT or PRODUCTION")
	}

	baseRouter := ginEngine.Group("/api")

	AddV1Routes(baseRouter)

	return ginEngine

}
