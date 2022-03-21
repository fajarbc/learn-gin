package controller

import (
	"log"
	"net/http"

	"github.com/fajarbc/learn-gin/models"
	"github.com/fajarbc/learn-gin/service"
	"github.com/fajarbc/learn-gin/validators"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

type ArticleController interface {
	FindAll(ctx *gin.Context) []models.Article
	Save(ctx *gin.Context) error
	ShowAll(ctx *gin.Context)
}

type controller struct {
	service service.ArticleService
}

var validate *validator.Validate

// belum tau untuk apa
func New(service service.ArticleService) ArticleController {
	validate = validator.New()
	validate.RegisterValidation("has-space", validators.ValidateHasSpace)
	return &controller{
		service: service,
	}
}

func (c *controller) FindAll(ctx *gin.Context) []models.Article {
	return c.service.FindAll(ctx)
}

func (c *controller) Save(ctx *gin.Context) error {
	var article models.Article
	err := ctx.ShouldBindJSON(&article)
	if err != nil {
		return err
	}
	err = validate.Struct(article)
	if err != nil {
		return err
	}
	c.service.Save(ctx, article)
	return nil
}

func (c *controller) ShowAll(ctx *gin.Context) {
	articles := c.service.FindAll(ctx)
	log.Println(articles)
	data := gin.H{
		"title":    "Article Page",
		"articles": articles,
	}

	ctx.HTML(http.StatusOK, "index.html", data)
}
