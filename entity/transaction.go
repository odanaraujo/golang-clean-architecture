package entity

import "errors"

const (
	APPROVED = "approved"
	REJECTED = "rejected"
)

type Transaction struct {
	ID           string
	AccountID    string
	Amount       float64
	Status       string
	ErrorMessage string
}

func NewTransaction() *Transaction {
	// return &Transaction{
	// 	ID:        transaction.ID,
	// 	AccountID: transaction.AccountID,
	// 	Amount:    transaction.Amount,
	// 	Status:    transaction.Status,
	// }
	return &Transaction{}
}

func (t *Transaction) IsValid() error {
	if t.Amount > 1000 {
		return errors.New("you dont have limit for this transaction")
	}

	if t.Amount < 1 {
		return errors.New("the amount must be greater than 1")
	}
	return nil
}
