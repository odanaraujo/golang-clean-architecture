package processtransaction

import (
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/odanaraujo/golang/clean-architecture/entity"
	mock_entity "github.com/odanaraujo/golang/clean-architecture/infrastructure/database/sql/mock"
	"github.com/stretchr/testify/assert"
)

func TestProcessTransactionWhenItsValid(t *testing.T) {
	input := TransactionDtoInput{
		ID:        "1",
		AccountID: "1",
		Amount:    200,
	}

	expectedOutput := TransactionDtoOutput{
		ID:           "1",
		Status:       "approved",
		ErrorMessage: "",
	}

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	repositoryMock := mock_entity.NewMockTransactionRepository(ctrl)
	repositoryMock.EXPECT().Insert(input.ID, input.AccountID, input.Amount, entity.APPROVED, "").Return(nil)

	usecase := NewProcessTransaction(repositoryMock)
	output, err := usecase.Execute(input)

	assert.Nil(t, err)
	assert.Equal(t, expectedOutput, output)

}

func TestProcessTransactionWhenDontsValid(t *testing.T) {
	input := TransactionDtoInput{
		ID:        "1",
		AccountID: "1",
		Amount:    0,
	}

	expectedOutput := TransactionDtoOutput{
		ID:           "1",
		Status:       "rejected",
		ErrorMessage: "the amount must be greater than 1",
	}

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	repositoryMock := mock_entity.NewMockTransactionRepository(ctrl)
	repositoryMock.EXPECT().Insert(input.ID, input.AccountID, input.Amount, entity.REJECTED, "the amount must be greater than 1").Return(nil)

	usecase := NewProcessTransaction(repositoryMock)
	output, err := usecase.Execute(input)

	assert.Nil(t, err)
	assert.Equal(t, expectedOutput, output)

}
