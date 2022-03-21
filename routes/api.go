package routes

import (
	"net/http"

	"github.com/fajarbc/learn-gin/middleware"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func ApiRoutes(server *gin.Engine, db *gorm.DB) {
	// only authorized user could access "/api"
	routes := server.Group("/api", middleware.AuthorizeJWT())
	{
		routes.GET("/articles", func(ctx *gin.Context) {
			ctx.JSON(http.StatusOK, articleController.FindAll(ctx))
		})

		routes.POST("/articles", func(ctx *gin.Context) {
			err := articleController.Save(ctx)
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
}
