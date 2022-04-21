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
	FindByUsername(db *gorm.DB, username string) (bool, models.Author)
	FindById(db *gorm.DB, id int) (bool, models.Author)
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

func (service *authorService) FindByUsername(db *gorm.DB, username string) (bool, models.Author) {
	var author models.Author
	result := db.Where(&models.Author{
		Username: username,
	}).First((&author))
	if result.RowsAffected > 0 {
		return true, author
	}
	return false, models.Author{}
}

func (service *authorService) FindById(db *gorm.DB, id int) (bool, models.Author) {
	var author models.Author
	result := db.Where(&models.Author{
		ID: id,
	}).First((&author))
	if result.Error == nil {
		if result.RowsAffected > 0 {
			return true, author
		}
	}
	return false, models.Author{}
}

func NewAuthorService() AuthorService {
	return &authorService{
		authors: []models.Author{},
	}
}
