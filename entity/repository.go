package entity

type TransactionRepository interface {
	Insert(id string, account_id string, amount float64, status string, errorMessage string) error
}
