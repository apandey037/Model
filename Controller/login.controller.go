package controller

import (
	"example.com/Tranction/Model"
	"example.com/Tranction/Security"
	"example.com/Tranction/Service"
	"github.com/gin-gonic/gin"
	"net/http"
)

type LoginControllers interface {
	Login(ctx *gin.Context) string
}

type LoginController struct {
	loginService Service.LoginService
	jWtService   Security.JWTService
}

func LoginHandler(loginService Service.LoginService,
	jWtService Security.JWTService) LoginControllers {
	return &LoginController{
		loginService: loginService,
		jWtService:   jWtService,
	}
}

func (controllers *LoginController) Login(ctx *gin.Context) string {
	var credential Model.Login
	err := ctx.ShouldBind(&credential)
	if err != nil {
		return "no data found"
	}
	isUserAuthenticated := controllers.loginService.LoginUser(credential.Email, credential.Password)
	if isUserAuthenticated {
		return controllers.jWtService.GenerateToken(credential.Email, true)

	}
	return ""
}

func (controllers *LoginController) LoginCustomer(rg *gin.RouterGroup) {
	customerroute := rg.Group("/login")
	var loginService Service.LoginService = Service.StaticLoginService()
	var jwtService = Security.JWTAuthService()
	var loginControllers = LoginHandler(loginService, jwtService)
	customerroute.POST("/signin", func(ctx *gin.Context) {
		//fmt.Println("text")
		token := loginControllers.Login(ctx)

		if token != "" {
			ctx.JSON(http.StatusOK, gin.H{
				"token": token,
			})
		} else {
			ctx.JSON(http.StatusUnauthorized, nil)
		}
	})
}
