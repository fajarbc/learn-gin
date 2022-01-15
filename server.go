package main

import (
	"io"
	"net/http"
	"os"

	"github.com/fajarbc/learn-gin/controller"
	"github.com/fajarbc/learn-gin/middleware"
	"github.com/fajarbc/learn-gin/service"
	"github.com/gin-gonic/gin"
)

var (
	videoService service.VideoService = service.New()
	loginService service.LoginService = service.NewLoginService()
	jwtService   service.JWTService   = service.NewJWTService()

	videoController controller.VideoController = controller.New(videoService)
	loginController controller.LoginController = controller.NewLoginController(loginService, jwtService)
)

func setupLogOutput() {
	f, _ := os.Create("gin.log")
	gin.DefaultWriter = io.MultiWriter(f, os.Stdout)
}

func main() {
	setupLogOutput()

	server := gin.New()

	// load static files. example: css, js
	server.Static("/css", "./templates/css")

	// load html
	server.LoadHTMLGlob("templates/*.html")

	// server.Use(gin.Recovery(), gin.Logger(), middleware.BasicAuth())
	server.Use(gin.Recovery(), gin.Logger())

	loginRoutes := server.Group("/login")
	{
		loginRoutes.POST("/", func(ctx *gin.Context) {
			token := loginController.Login(ctx)
			message := "OK"
			error := false
			if token == "" {
				message = "Login Failed"
				error = true
			}
			ctx.JSON(http.StatusOK, gin.H{
				"error":   error,
				"message": message,
				"token":   token,
			})
		})
	}

	// only authorized user could access "/api"
	apiRoutes := server.Group("/api", middleware.AuthorizeJWT())
	{
		apiRoutes.GET("/videos", func(ctx *gin.Context) {
			ctx.JSON(http.StatusOK, videoController.FindAll())
		})

		apiRoutes.POST("/videos", func(ctx *gin.Context) {
			err := videoController.Save(ctx)
			status := http.StatusOK
			message := "OK"
			error := false
			if err != nil {
				status = http.StatusBadRequest
				error = true
				message = err.Error()
			}
			ctx.JSON(status, gin.H{
				"error":   error,
				"message": message,
			})
		})
	}

	viewRoutes := server.Group("view")
	{
		viewRoutes.GET("/videos", videoController.ShowAll)
	}
	// running server, default in port 8080
	server.Run()
}
