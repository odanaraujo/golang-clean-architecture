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

	isValid := transaction.IsValid()
	assert.Error
}