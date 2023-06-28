package entity

import "errors"

type Transaction struct {
	ID           string
	AccountID    string
	Amount       float64
	Status       string
	ErrorMessage string
}

func NewTransaction(transaction Transaction) *Transaction {
	return &Transaction{
		ID:        transaction.ID,
		AccountID: transaction.AccountID,
		Amount:    transaction.Amount,
		Status:    transaction.Status,
	}
}

func (t *Transaction) IsValid() error {
	if t.Amount > 1000 {
		return errors.New("you dont have limit for this transaction")
	}
	return nil
}
