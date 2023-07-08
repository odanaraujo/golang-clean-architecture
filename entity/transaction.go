package entity

import (
	"errors"
)

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

func (t *Transaction) NewTransaction() *Transaction {
	return &Transaction{
		ID:        t.ID,
		AccountID: t.AccountID,
		Amount:    t.Amount,
		Status:    t.Status,
	}
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
