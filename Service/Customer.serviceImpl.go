package Service

import (
	"example.com/Tranction/Model"
	"fmt"
	as "github.com/aerospike/aerospike-client-go"
	"github.com/mitchellh/mapstructure"
	"log"
)

type CustomerServiceImpl struct {
	client *as.Client
}

func NewCustomerServiceImpl(client *as.Client) CustomerService {
	return &CustomerServiceImpl{
		client: client,
	}

}

func (u *CustomerServiceImpl) CreateCustomer(customer *Model.Customer) error {
	//TODO implement me
	key, err := as.NewKey("test", "Customer", customer.AccountNo)

	err = u.client.PutObject(nil, key, customer)

	if err != nil {

		return err

	}
	return nil
}

func (u *CustomerServiceImpl) GetCustomer(AccountNo int) (Model.Customer, error) {
	//TODO implement me
	var customer Model.Customer
	key, err := as.NewKey("test", "Customer", AccountNo)

	record, err := u.client.Get(nil, key)
	//fmt.Println("Printing name values", record.Bins)
	mapstructure.Decode(record.Bins, &customer)

	//fmt.Println("printing user obj values ", user)
	return customer, err
}

func (u *CustomerServiceImpl) GetAll() ([]Model.Customer, error) {
	//TODO implement me
	var customer []Model.Customer
	AllRecord, err := u.client.ScanAll(nil, "test", "Customer")
	if err != nil {
		return nil, err
	}
	fmt.Println(AllRecord)
	//All.len
	for records := range AllRecord.Records {
		fmt.Println("Enter into records")
		if records != nil {

			var singleCustomer Model.Customer
			mapstructure.Decode(records.Bins, &singleCustomer)
			customer = append(customer, singleCustomer)
		}
	}
	return customer, err
}

func (u *CustomerServiceImpl) UpdateCustomer(AccountNo int) error {
	//TODO implement me
	//var customer Model.Customer
	//key, err := as.NewKey("test", "Customer", AccountNo)
	//record, err := u.client.Get(nil, key)

	return nil
}

func (u *CustomerServiceImpl) DeleteCustomer(AccountNo int) error {
	//var AccountNo Model.Customer
	key, err := as.NewKey("test", "Customer", AccountNo)
	existed, err := u.client.Delete(nil, key)
	if err != nil {
		log.Fatal(err)
	}
	if existed {
		fmt.Println("Record deleted")
	}
	return err
}
