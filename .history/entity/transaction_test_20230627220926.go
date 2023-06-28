package entity

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestTransactionWithAmountGreaterThan100(t *testing.T) {
	transaction := NewTransaction()
	transaction.ID = 1
	transaction.AccountID = 1
	transaction.Amount = 2000

	err := transaction.IsValid()
	assert.Error(t, err)
	assert.Equal(t, "tou dont have limit for this transaction")
}