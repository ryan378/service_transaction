package usecase

import(
	//"api_gateway/handler"
	//"github.com/gin-gonic/gin"
)

type Login struct {
}

type LoginInterface interface {
	Autentikasi(username, password string) bool
}

func NewLogin() LoginInterface {
	return &Login{}
}

func (a *Login) Autentikasi(username, password string) bool {
	if username == "admin" && password == "admin123" {
		return true
	}
	return false
}
