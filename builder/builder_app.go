package builder

import "fmt"

func init() {
	// This is Old customer
	customer1 := NewCustomer(&Customer{
		Id:        "0001",
		FirstName: "cust",
		LastName:  "omer",
		Email:     "customer1@g.co",
		Phone:     "0811",
	})
	fmt.Println(customer1)
	
	// This is new Customer with new Property
	customer2 := NewCustomer(&Customer{
		Id:			"0002",
		FirstName:	"cust",
		LastName:	"omer",
		Email:		"customer2@g.co",
		Phone:		"0811",
		Adress: 	"California",
		Age:		20,	
	})
	fmt.Println(customer2)
}