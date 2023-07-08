package entity

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestTransactionWithAmountGreaterThan100(t *testing.T) {

	tt := Transaction{}
	tt.ID = "1"
	tt.AccountID = "1"
	tt.Amount = 2000

	transaction := tt.NewTransaction()

	err := transaction.IsValid()
	assert.Error(t, err)
	assert.Equal(t, "you dont have limit for this transaction", err.Error())
}

func TestTransactionWithAmountLesserThan1(t *testing.T) {
	tt := Transaction{}
	tt.ID = "1"
	tt.AccountID = "1"
	tt.Amount = 0

	transaction := tt.NewTransaction()

	err := transaction.IsValid()
	assert.Error(t, err)
	assert.Equal(t, "the amount must be greater than 1", err.Error())
}
