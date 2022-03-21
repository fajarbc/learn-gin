package routes

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func ViewRoutes(server *gin.Engine, db *gorm.DB) {
	routes := server.Group("view")
	{
		routes.GET("/articles", articleController.ShowAll)
	}
}
