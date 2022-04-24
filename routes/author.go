package routes

import (
	"net/http"

	"github.com/fajarbc/learn-gin/controller"
	"github.com/fajarbc/learn-gin/service"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

var (
	jwtService    service.JWTService    = service.NewJWTService()
	authorService service.AuthorService = service.NewAuthorService()

	authorContoller controller.AuthorController = controller.NewAuthorController(authorService, jwtService)
)

func AuthorRoutes(server *gin.Engine, db *gorm.DB) {
	routes := server.Group("/author")
	{
		routes.POST("/login", func(ctx *gin.Context) {
			success, message, author := authorContoller.Login(ctx)
			ctx.JSON(http.StatusOK, gin.H{
				"success": success,
				"message": message,
				"author":  author,
			})
		})

		routes.POST("/register", func(ctx *gin.Context) {
			success, message := authorContoller.Register(ctx)
			ctx.JSON(http.StatusOK, gin.H{
				"success": success,
				"message": message,
			})
		})
	}
}
