package Service

import "strings"

type LoginService interface {
	LoginUser(email string, password string) bool
}

type loginInformation struct {
	email    string
	password string
}

func StaticLoginService() *loginInformation {
	return &loginInformation{
		email:    "shreyasraut123@gmail.com",
		password: "String",
	}
}

func (info *loginInformation) LoginUser(email string, password string) bool {
	info.email = "shreyasraut123@gmail.com"
	info.password = "string"
	return strings.EqualFold(info.email, email) && strings.EqualFold(info.password, password)

}
