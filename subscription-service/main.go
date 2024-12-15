package main

import (
	"log"
	"net/http"
	"subscription-service/config"
	"subscription-service/database"
	"subscription-service/routes"
)

func main() {

	//configuration
	appConfig := config.LoadConfig()

	// Initialize database
	database.InitDB(appConfig.DBUrl)

	// Initialize Redis
	database.InitRedis(appConfig.RedisUrl)

	//routes
	r := routes.SetupRouter(appConfig.JwtSecret)

	log.Printf("Starting Go server on port %s", appConfig.Port)
	http.ListenAndServe(":"+appConfig.Port, r)

}
