package routes

import (
	"net/http"

	"github.com/fajarbc/learn-gin/controller"
	"github.com/fajarbc/learn-gin/service"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

var (
	loginService service.LoginService = service.NewLoginService()
	jwtService   service.JWTService   = service.NewJWTService()

	loginController controller.LoginController = controller.NewLoginController(loginService, jwtService)
)

func LoginRoutes(server *gin.Engine, db *gorm.DB) {
	routes := server.Group("/login")
	{
		routes.POST("/", func(ctx *gin.Context) {
			token := loginController.Login_(ctx)
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
