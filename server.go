package main

import (
	"io"
	"os"

	"github.com/fajarbc/learn-gin/models"
	"github.com/fajarbc/learn-gin/routes"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func setupLogOutput() {
	f, _ := os.Create("gin.log")
	gin.DefaultWriter = io.MultiWriter(f, os.Stdout)
}

func main() {
	setupLogOutput()
	db := models.ConnectDB()
	models.AutoMigrateDB(db)

	err := godotenv.Load(".env")
	if err != nil {
		panic("Can not load .env")
	}

	server := gin.New()

	// load static files. example: css, js
	server.Static("/css", "./templates/css")

	// load html
	server.LoadHTMLGlob("templates/*.html")

	server.Use(gin.Recovery(), gin.Logger())

	// routes
	server = routes.SetupRoutes(db)

	// running server, default in port 8080
	server.Run()
}
