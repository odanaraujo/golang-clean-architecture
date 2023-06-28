package entity

import (
	"strconv"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestTransactionWithAmountGreaterThan100(t *testing.T) {
	transaction := NewTransaction()
	transaction.ID = "1"
	transaction.AccountID = strconv.Itoa(1)
	transaction.Amount = 2000

	err := transaction.IsValid()
	assert.Error(t, err)
	assert.Equal(t, "you dont have limit for this transaction", err.Error())
}
