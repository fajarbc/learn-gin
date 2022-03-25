package service

import (
	"os"
	"time"

	"github.com/fajarbc/learn-gin/models"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

// import "github.com/fajarbc/learn-gin/models"

type AuthorService interface {
	Login(db *gorm.DB, username string, password string) (bool, string, models.Author)
	Save(*gin.Context, models.Author) (bool, string)
}

type authorService struct {
	authors []models.Author
}

func (service *authorService) Login(db *gorm.DB, username string, password string) (bool, string, models.Author) {
	var author models.Author
	enryptedPassword, err := Encrypt(password, os.Getenv("encryption_key"))
	if err != nil {
		panic(err.Error())
	}
	result := db.Where(&models.Author{
		Username: username,
		Password: enryptedPassword,
	}).First((&author))
	if result.Error != nil {
		return false, "username & password did not match", models.Author{}
	}
	return true, "Berhasil", author
}

func (service *authorService) Save(ctx *gin.Context, author models.Author) (bool, string) {
	service.authors = append(service.authors, author)
	password, err := Encrypt(author.Password, os.Getenv("encryption_key"))
	if err != nil {
		return false, err.Error()
	}
	save := &models.Author{
		Name:      author.Name,
		Email:     author.Email,
		Status:    author.Status,
		Username:  author.Username,
		Password:  password,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	db := ctx.MustGet("db").(*gorm.DB)
	tx := db.Create(&save)
	if tx.Error != nil {
		return false, tx.Error.Error()
	}
	return true, "success"
}

func NewAuthorService() AuthorService {
	return &authorService{
		authors: []models.Author{},
	}
}
