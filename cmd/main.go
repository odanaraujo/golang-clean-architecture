package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"

	_ "github.com/mattn/go-sqlite3"
	"github.com/odanaraujo/golang/clean-architecture/adapter/repository"
	processtransaction "github.com/odanaraujo/golang/clean-architecture/usecase/process_transaction"
)

func main() {

	db, err := sql.Open("sqlite3", "test.db")

	if err != nil {
		log.Fatal(err)
	}

	repo := repository.NewTransactionRepositoryDB(db)

	usecase := processtransaction.NewProcessTransaction(repo)

	input := processtransaction.TransactionDtoInput{
		ID:        "1",
		AccountID: "1",
		Amount:    0,
	}

	output, err := usecase.Execute(input)

	if err != nil {
		fmt.Println(err.Error())
	}

	body, err := json.Marshal(output)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(string(body))
}
