package service

type LoginService interface {
	Login(username string, password string) bool
}

type loginService struct {
	username string
	password string
}

func (service *loginService) Login(username string, password string) bool {
	return service.username == username &&
		service.password == password
}

func NewLoginService() LoginService {
	return &loginService{
		username: "fajar",
		password: "bece",
	}
}
