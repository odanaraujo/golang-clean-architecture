package repository

import (
	"database/sql"
	"time"
)

type TransactionalRepositoryDb struct {
	db *sql.DB
}

func NewTransactionRepositoryDB(db *sql.DB) *TransactionalRepositoryDb {
	return &TransactionalRepositoryDb{db: db}
}

func (repo *TransactionalRepositoryDb) Insert(id string, account_id string, amount float64, status string, errorMessage string) error {

	stmt, err := repo.db.Prepare(`
		Insert into transaction(id, account_id, amount, status, error_message, created_at, updated_at)
		values($1, $2, $3, $4, $5, $6, $7)
		`)

	if err != nil {
		return err
	}

	_, err = stmt.Exec(id, account_id, amount, status, errorMessage, time.Now(), time.Now())

	if err != nil {
		return err
	}

	return nil
}
