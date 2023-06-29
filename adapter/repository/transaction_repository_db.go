package repository

import (
	"database/sql"
)

type TransactionalRepositoryDb struct {
	db *sql.DB
}

func NewTransactionRepositoryDB(db *sql.DB) *TransactionalRepositoryDb {
	return &TransactionalRepositoryDb{db: db}
}

func (repo *TransactionalRepositoryDb) Insert(id string,
	accountId string,
	amount float64,
	status string,
	errorMessage string) error {

	stmt, err := repo.db.
}
