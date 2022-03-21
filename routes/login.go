package routes

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func LoginRoutes(server *gin.Engine, db *gorm.DB) {
	routes := server.Group("/login")
	{
		routes.POST("/", func(ctx *gin.Context) {
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
}
