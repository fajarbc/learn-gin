package routes

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func SetupRoutes(db *gorm.DB) *gin.Engine {
	server := gin.Default()

	server.Use(func(c *gin.Context) {
		c.Set("db", db)
	})

	ApiRoutes(server, db)
	AuthorRoutes(server, db)
	ViewRoutes(server, db)

	return server
}
