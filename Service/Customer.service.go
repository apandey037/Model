package Service

import "example.com/Tranction/Model"

type CustomerService interface {
	CreateCustomer(*Model.Customer) error
	GetCustomer(int) (Model.Customer, error)
	GetAll() ([]Model.Customer, error)
	UpdateCustomer(AccoutNo int) error
	DeleteCustomer(AccoutNo int) error
}
