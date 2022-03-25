package controller

import (
	"github.com/fajarbc/learn-gin/dto"
	"github.com/fajarbc/learn-gin/models"
	"github.com/fajarbc/learn-gin/service"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type AuthorController interface {
	Login(ctx *gin.Context) (bool, string, models.Author)
	Register(ctx *gin.Context) (bool, string)
}

type authorController struct {
	authorService service.AuthorService
	jwtService    service.JWTService
}

func (controller *authorController) Login(ctx *gin.Context) (bool, string, models.Author) {
	var credentials dto.Credentials
	status := false
	err := ctx.ShouldBind(&credentials)
	if err != nil {
		return status, err.Error(), models.Author{}
	}
	db := ctx.MustGet("db").(*gorm.DB)
	isAuthenticated, message, author := controller.authorService.Login(db, credentials.Username, credentials.Password)
	if isAuthenticated {
		status = true
		message = controller.jwtService.GenerateToken(credentials.Username, true)
	}
	return status, message, author
}

func (controller *authorController) Register(ctx *gin.Context) (bool, string) {
	var author models.Author
	err := ctx.ShouldBindJSON(&author)
	if err != nil {
		return false, "error 1" //err.Error()
	}
	err = validate.Struct(author)
	if err != nil {
		return false, "error 2" //err.Error()
	}
	success, message := controller.authorService.Save(ctx, author)
	return success, message
}

func NewAuthorController(authorService service.AuthorService, jwtService service.JWTService) AuthorController {
	return &authorController{
		authorService: authorService,
		jwtService:    jwtService,
	}
}