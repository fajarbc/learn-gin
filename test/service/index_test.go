package service_test

import (
	"fmt"
	"os"
	"testing"

	"github.com/fajarbc/learn-gin/models"
	"github.com/fajarbc/learn-gin/service"
	"github.com/joho/godotenv"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"

	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/suite"
)

func TestMain(m *testing.M) {
	// load env
	err := godotenv.Load("../../.env")
	if err != nil {
		panic("Can not load .env")
	}

	fmt.Println("service test started")
	m.Run()
	fmt.Println("service test finished")
}

type SuiteService struct {
	suite.Suite
	DB *gorm.DB

	// service
	authorService service.AuthorService
	jwtService    service.JWTService

	// default data
	secretKey     string
	plainText     string
	base64Text    string
	encryptedText string
	authorValid   models.AuthorLogin
	authorInvalid models.AuthorLogin
}

func (s *SuiteService) SetupTest() {
	var err error

	// set db
	LINK := models.GetDatabaseLink()
	s.DB, err = gorm.Open(mysql.Open(LINK))
	require.NoError(s.T(), err)

	s.authorService = service.NewAuthorService()
	s.jwtService = service.NewJWTService()

	s.secretKey = os.Getenv("encryption_key")
	s.authorValid = models.AuthorLogin{
		Username: "user",
		Password: "user",
	}

	s.authorInvalid = models.AuthorLogin{
		Username: "user",
		Password: "wrong",
	}

	s.plainText = "Test Passed"
	s.base64Text = "VGVzdCBQYXNzZWQ="
	s.encryptedText = "iga/KLS2pGoiQWA="

}

func Test_Service(t *testing.T) {
	suite.Run(t, new(SuiteService))
}
