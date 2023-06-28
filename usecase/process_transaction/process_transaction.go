package processtransaction

import (
	"fmt"

	"github.com/odanaraujo/golang/clean-architecture/entity"
)

type ProcessTransaction struct {
	Repository entity.TransactionRepository
}

func NewProcessTransaction(repository entity.TransactionRepository) *ProcessTransaction {
	return &ProcessTransaction{Repository: repository}
}

func (p *ProcessTransaction) Execute(input TransactionDtoInput) (TransactionDtoOutput, error) {

	transaction := entity.NewTransaction()
	transaction.ID = input.ID
	transaction.AccountID = input.AccountID
	transaction.Amount = input.Amount
	transaction.Status = entity.APPROVED

	output := TransactionDtoOutput{
		ID:           transaction.ID,
		Status:       entity.APPROVED,
		ErrorMessage: "",
	}

	if invalidTransaction := transaction.IsValid(); invalidTransaction != nil {
		output.Status = entity.REJECTED
		output.ErrorMessage = invalidTransaction.Error()
	}

	fmt.Println("Exemplo de Dan", output.ErrorMessage)
	fmt.Println("Exemplo de Dan", transaction.Status)

	if err := p.Repository.Insert(transaction.ID, transaction.AccountID, transaction.Amount, output.Status, output.ErrorMessage); err != nil {
		return TransactionDtoOutput{}, err
	}

	return output, nil

}
