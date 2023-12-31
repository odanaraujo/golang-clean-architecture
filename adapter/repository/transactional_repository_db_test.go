package repository

import (
	"os"
	"testing"

	_ "github.com/mattn/go-sqlite3"
	"github.com/odanaraujo/golang/clean-architecture/adapter/repository/fixture"
	"github.com/odanaraujo/golang/clean-architecture/entity"
	"github.com/stretchr/testify/assert"
)

func TestTransactionalRepositoryDB_Insert(t *testing.T) {
	migrationDir := os.DirFS("fixture/sql")
	db := fixture.Up(migrationDir)
	defer fixture.Down(db, migrationDir)

	repository := NewTransactionRepositoryDB(db)

	err := repository.Insert("1", "1", 2, entity.APPROVED, "")

	assert.Nil(t, err)
}
