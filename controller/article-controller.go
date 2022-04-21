package controller

import (
	"errors"
	"log"
	"net/http"

	"github.com/fajarbc/learn-gin/middleware"
	"github.com/fajarbc/learn-gin/models"
	"github.com/fajarbc/learn-gin/service"

	"github.com/fajarbc/learn-gin/validators"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"gorm.io/gorm"
)

type ArticleController interface {
	FindAll(ctx *gin.Context) []models.Article
	Save(ctx *gin.Context) error
	ShowAll(ctx *gin.Context)
}

type controller struct {
	articleService service.ArticleService
	authorService  service.AuthorService
}

var validate *validator.Validate

// untuk articleController
func New(articleService service.ArticleService, authorService service.AuthorService) ArticleController {
	validate = validator.New()
	validate.RegisterValidation("has-space", validators.ValidateHasSpace)
	return &controller{
		articleService: articleService,
		authorService:  authorService,
	}
}

func (c *controller) FindAll(ctx *gin.Context) []models.Article {
	return c.articleService.FindAll(ctx)
}

func (c *controller) Save(ctx *gin.Context) error {
	var article models.Article
	var articleInsert models.ArticleInsert
	// make sure json payload is exist
	if err := ctx.ShouldBindJSON(&articleInsert); err != nil {
		return err
	}

	// validate each json key
	if err := validate.Struct(articleInsert); err != nil {
		return err
	}

	// check author and add it to insert data
	claim, err := middleware.GetTokenClaim(ctx)
	if err != nil {
		return err
	}
	// log.Println("GetTokenClaim() result")
	// log.Println(claim)
	authorID := int(claim["id"].(float64))
	db := ctx.MustGet("db").(*gorm.DB)
	isExisted, author := c.authorService.FindById(db, authorID)
	if !isExisted {
		return errors.New("Author not found (" + claim["id"].(string) + ")")
	}
	log.Println("author")
	log.Println(author)
	article.Author = author

	c.articleService.Save(ctx, article)
	return nil
}

func (c *controller) ShowAll(ctx *gin.Context) {
	articles := c.articleService.FindAll(ctx)
	log.Println(articles)
	data := gin.H{
		"title":    "Article Page",
		"articles": articles,
	}

	ctx.HTML(http.StatusOK, "index.html", data)
}
