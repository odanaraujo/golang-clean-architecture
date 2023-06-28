package entity

import "testing"

func TestTransactionWithAmountGreaterThan100(t *testing.T) {
	transaction := NewTransaction()
	transaction.ID = 1
	transaction.A
}