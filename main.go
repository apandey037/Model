package main

import (
	"context"
	controller "example.com/Tranction/Controller"
	"fmt"
	"log"

	"example.com/Tranction/Service"
	"github.com/aerospike/aerospike-client-go"
	"github.com/gin-gonic/gin"
)

var (
	server             *gin.Engine
	customerservice    Service.CustomerService
	customercontroller controller.CustomerController
	logincontroller    controller.LoginController
	ctx                context.Context
	client             *aerospike.Client
	key                *aerospike.Key
	err                error
)

func init() {
	const Namespace = "test"
	const Set = "Customer"
	client, err := aerospike.NewClient("localhost", 3000)

	fmt.Println("Aerospike connection established", client.IsConnected())
	if err != nil {
		log.Fatal(err)
	}
	//client.Close()
	customerservice = Service.NewCustomerServiceImpl(client)
	customercontroller = controller.New(customerservice)
	server = gin.Default()
}
func main() {

	defer client.Close()

	basepath := server.Group("/v1")
	customercontroller.RegisterUserRoutes(basepath)
	logincontroller.LoginCustomer(basepath)
	log.Fatal(server.Run(":9090"))
}
