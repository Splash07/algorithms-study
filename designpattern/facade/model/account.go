package model

import "fmt"

// Account struct
type Account struct {
	Id          string
	AccountType string
}

// Account class method create - create account given Account Type
func (account *Account) Create(accountType string) *Account {
	account.AccountType = accountType
	return account
}

// Account class method getById given id string
func (account *Account) GetById(id string) *Account {
	fmt.Println("getting account by Id")
	return account
}

// Account class method deleteById given id string
func (account *Account) DeleteById(id string) {
	fmt.Println("delete account by id")
}
