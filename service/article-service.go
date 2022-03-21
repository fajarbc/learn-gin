package service

import (
	"time"

	m "github.com/fajarbc/learn-gin/models"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type ArticleService interface {
	Save(*gin.Context, m.Article) m.Article
	FindAll(*gin.Context) []m.Article
}

type articleService struct {
	articles []m.Article
}

// belum tau ini buat apa
func New() ArticleService {
	return &articleService{
		articles: []m.Article{},
	}
}

// insert data baru
func (service *articleService) Save(ctx *gin.Context, article m.Article) m.Article {
	service.articles = append(service.articles, article)
	save := m.Article{
		Title:   article.Title,
		Content: article.Content,
		Slug:    article.Slug,
		Status:  article.Status,
		Author: m.Author{
			Name:   article.Author.Name,
			Email:  article.Author.Email,
			Status: article.Author.Status,
		},
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	db := ctx.MustGet("db").(*gorm.DB)
	db.Create(&save)
	return article
}

func (service *articleService) FindAll(ctx *gin.Context) []m.Article {
	// Get all records
	var articles []m.Article
	db := ctx.MustGet("db").(*gorm.DB)
	// result := db.Preload("Author").Find(&articles)
	result := db.Joins("Author", db.Where(&m.Author{Status: 1})).Find(&articles)
	if result.Error != nil {
		return []m.Article{}
	}
	// db.Model(&m.Article{}).Limit(10).Find(&m.Author{})
	return articles
}
