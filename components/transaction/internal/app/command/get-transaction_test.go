package command

import (
	"context"
	"errors"
	"testing"

	tx "github.com/LerianStudio/midaz/components/transaction/internal/domain/transaction"
	mock "github.com/LerianStudio/midaz/components/transaction/internal/gen/mock/transaction"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"go.uber.org/mock/gomock"
)

func TestGetTransactionByID(t *testing.T) {
	ID := uuid.New()
	organizationID := uuid.New()
	ledgerID := uuid.New()

	tran := &tx.Transaction{
		ID:             ID.String(),
		OrganizationID: organizationID.String(),
		LedgerID:       ledgerID.String(),
	}

	uc := UseCase{
		TransactionRepo: mock.NewMockRepository(gomock.NewController(t)),
	}

	uc.TransactionRepo.(*mock.MockRepository).
		EXPECT().
		Find(gomock.Any(), gomock.Any(), gomock.Any(), gomock.Any()).
		Return(tran, nil).
		Times(1)
	res, err := uc.TransactionRepo.Find(context.TODO(), organizationID, ledgerID, ID)

	assert.Equal(t, tran, res)
	assert.Nil(t, err)
}

func TestGetTransactionByIDError(t *testing.T) {
	errMSG := "err to create account on database"
	ID := uuid.New()
	organizationID := uuid.New()
	ledgerID := uuid.New()

	uc := UseCase{
		TransactionRepo: mock.NewMockRepository(gomock.NewController(t)),
	}

	uc.TransactionRepo.(*mock.MockRepository).
		EXPECT().
		Find(gomock.Any(), gomock.Any(), gomock.Any(), gomock.Any()).
		Return(nil, errors.New(errMSG)).
		Times(1)
	res, err := uc.TransactionRepo.Find(context.TODO(), organizationID, ledgerID, ID)

	assert.NotEmpty(t, err)
	assert.Equal(t, err.Error(), errMSG)
	assert.Nil(t, res)
}
