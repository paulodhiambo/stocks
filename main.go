package main

import (
	"github.com/joho/godotenv"
	"log"
	"os"
	"stocks/app/router"
	"stocks/config"
)

func init() {
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Unable to read environment variables %s", err)
	}
	config.InitLog()
}

func main() {
	port := os.Getenv("PORT")

	init := config.Init()
	app := router.Init(init)

	err := app.Run(":" + port)
	if err != nil {
		log.Fatalf("Unable to start application on the provided port %s", err)
	}
}
