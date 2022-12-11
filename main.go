package main

import (
	database "learn/go-gin/config"
	"learn/go-gin/model"
	route "learn/go-gin/routes"
	"log"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	loadEnv()
	loadDatabase()
	serve()
}

func loadEnv() {
	err := godotenv.Load(".env")

	if err != nil {
		log.Fatal("Can't loading .env file")
	}
}

func loadDatabase() {
	database.Connect()
	database.Database.AutoMigrate(&model.Product{})
}

func serve() {
	gin.SetMode(gin.DebugMode)
	r := route.SetupRouter()

	r.Run()
}
