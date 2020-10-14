package model

import "fmt"

// Customer struct
type Customer struct {
	Name string
	Id   int
}

// Customer class method create - create Customer given name
func (customer *Customer) Create(name string) *Customer {
	fmt.Println("creating customer")
	customer.Name = name
	return customer
}
