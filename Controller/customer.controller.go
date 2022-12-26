package controller

import (
	"net/http"
	"strconv"

	"example.com/Tranction/Model"
	"example.com/Tranction/Service"
	"github.com/gin-gonic/gin"
)

type CustomerController struct {
	CustomerService Service.CustomerService
}

func New(CustomerService Service.CustomerService) CustomerController {
	return CustomerController{
		CustomerService: CustomerService,
	}
}
func (uc *CustomerController) CreateCustomer(ctx *gin.Context) {
	var customer Model.Customer
	if err := ctx.ShouldBindJSON(&customer); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}
	err := uc.CustomerService.CreateCustomer(&customer)
	if err != nil {
		ctx.JSON(http.StatusBadGateway, gin.H{"message": err.Error()})
	}
	ctx.JSON(http.StatusOK, customer)

}
func (uc *CustomerController) GetCustomer(ctx *gin.Context) {

	Accountno, _ := strconv.ParseInt(ctx.Query("account_no"), 10, 64)
	var customer Model.Customer
	//Accountno := customer.AccountNo
	customer, err := uc.CustomerService.GetCustomer(int(Accountno))
	if err != nil {
		ctx.JSON(http.StatusBadGateway, gin.H{"message": err.Error()})
	}
	ctx.JSON(http.StatusOK, customer)
}
func (uc *CustomerController) GetAll(ctx *gin.Context) {
	var customer []Model.Customer
	customer, err := uc.CustomerService.GetAll()
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
	}
	ctx.JSON(http.StatusOK, customer)
}
func (uc *CustomerController) UpdateCustomer(ctx *gin.Context) {

	//Accountno, _ := strconv.ParseInt(ctx.Query("account_no"), 10, 64)
	//var customer Model.Customer
	//if err := ctx.ShouldBindJSON(&customer); err != nil {
	//	ctx.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
	//	return
	//}
	//err := uc.CustomerService.UpdateCustomer(&customer)
	//if err != nil {
	//	ctx.JSON(http.StatusBadGateway, gin.H{"message": err.Error()})
	//}
	ctx.JSON(http.StatusOK, gin.H{"message": "success"})
}
func (uc *CustomerController) DeleteCustomer(ctx *gin.Context) {
	Accountno, _ := strconv.ParseInt(ctx.Query("account_no"), 10, 64)
	var customer Model.Customer
	err := uc.CustomerService.DeleteCustomer(int(Accountno))
	if err != nil {
		ctx.JSON(http.StatusBadGateway, gin.H{"message": err.Error()})
	}
	ctx.JSON(http.StatusOK, customer)

}
func (uc *CustomerController) RegisterUserRoutes(rg *gin.RouterGroup) {
	customerroute := rg.Group("/customer")
	customerroute.POST("/create", uc.CreateCustomer)
	customerroute.GET("/get/AccountNo", uc.GetCustomer)
	customerroute.GET("/getAll", uc.GetAll)
	customerroute.PATCH("/update", uc.UpdateCustomer)
	customerroute.DELETE("/delete/AccountNo", uc.DeleteCustomer)

}
