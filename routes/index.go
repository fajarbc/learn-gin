package routes

import (
	"github.com/fajarbc/learn-gin/controller"
	"github.com/fajarbc/learn-gin/service"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

var (
	loginService   service.LoginService   = service.NewLoginService()
	jwtService     service.JWTService     = service.NewJWTService()
	articleService service.ArticleService = service.New()

	loginController   controller.LoginController   = controller.NewLoginController(loginService, jwtService)
	articleController controller.ArticleController = controller.New(articleService)
)

func SetupRoutes(db *gorm.DB) *gin.Engine {
	server := gin.Default()

	server.Use(func(c *gin.Context) {
		c.Set("db", db)
	})

	ApiRoutes(server, db)
	LoginRoutes(server, db)
	ViewRoutes(server, db)

	return server
}
