package model

import "fmt"

// Transaction struct
type Transaction struct {
	Id            string
	Amount        float32
	SrcAccountId  string
	DestAccountId string
}

func (transaction *Transaction) Create(srcAccountId, destAccoutnId string, amount float32) *Transaction {
	fmt.Println("Creating transaction")
	transaction.SrcAccountId = srcAccountId
	transaction.DestAccountId = destAccoutnId
	transaction.Amount = amount
	return transaction
}
