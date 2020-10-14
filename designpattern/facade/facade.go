package main

import (
	"facade/model"
	"fmt"
)

// BranchManagerFacade struct
type BranchManagerFacade struct {
	account     *model.Account
	customer    *model.Customer
	transaction *model.Transaction
}

// method NewBranchManagerFacade
func NewBranchManagerFacade() *BranchManagerFacade {
	return &BranchManagerFacade{
		&model.Account{},
		&model.Customer{},
		&model.Transaction{},
	}
}

// BranchManagerFacade class method createCustomerAccount
func (facade *BranchManagerFacade) createCustomerAccount(customerName, accountType string) (*model.Customer, *model.Account) {
	customer := facade.customer.Create(customerName)
	account := facade.account.Create(accountType)
	return customer, account
}

// BranchManagerFacade class method createTransaction
func (facade *BranchManagerFacade) createTransaction(srcAccountId, destAccountId string, amount float32) *model.Transaction {
	transaction := facade.transaction.Create(srcAccountId, destAccountId, amount)
	return transaction
}

func main() {
	facade := NewBranchManagerFacade()
	customer, account := facade.createCustomerAccount("Dylan Nguyen", "Savings")
	fmt.Println(customer.Name)
	fmt.Println(account.AccountType)

	transaction := facade.createTransaction("1123", "4352", 10000)
	fmt.Println(transaction.Amount)
}
