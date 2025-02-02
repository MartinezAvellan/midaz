package query

import (
	"context"
	"errors"
	"go.uber.org/mock/gomock"
	"testing"

	"github.com/LerianStudio/midaz/components/ledger/internal/adapters/postgres/account"
	"github.com/LerianStudio/midaz/pkg"
	"github.com/LerianStudio/midaz/pkg/mmodel"

	"github.com/stretchr/testify/assert"
)

// TestGetAllAccounts is responsible to test GetAllAccounts with success and error
func TestGetAllAccounts(t *testing.T) {
	organizationID := pkg.GenerateUUIDv7()
	ledgerID := pkg.GenerateUUIDv7()
	portfolioID := pkg.GenerateUUIDv7()
	limit := 10
	page := 1

	t.Parallel()
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	mockAccountRepo := account.NewMockRepository(ctrl)

	uc := UseCase{
		AccountRepo: mockAccountRepo,
	}

	t.Run("Success", func(t *testing.T) {
		accounts := []*mmodel.Account{{}}
		mockAccountRepo.
			EXPECT().
			FindAll(gomock.Any(), organizationID, ledgerID, &portfolioID, limit, page).
			Return(accounts, nil).
			Times(1)
		res, err := uc.AccountRepo.FindAll(context.TODO(), organizationID, ledgerID, &portfolioID, limit, page)

		assert.NoError(t, err)
		assert.Len(t, res, 1)
	})

	t.Run("Error", func(t *testing.T) {
		errMsg := "errDatabaseItemNotFound"
		mockAccountRepo.
			EXPECT().
			FindAll(gomock.Any(), organizationID, ledgerID, &portfolioID, limit, page).
			Return(nil, errors.New(errMsg)).
			Times(1)
		res, err := uc.AccountRepo.FindAll(context.TODO(), organizationID, ledgerID, &portfolioID, limit, page)

		assert.EqualError(t, err, errMsg)
		assert.Nil(t, res)
	})
}

// TestGetAllAccountsWithoutPortfolio is responsible to test GetAllAccounts without portfolio with success and error
func TestGetAllAccountsWithoutPortfolio(t *testing.T) {
	organizationID := pkg.GenerateUUIDv7()
	ledgerID := pkg.GenerateUUIDv7()
	limit := 10
	page := 1

	t.Parallel()
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	mockAccountRepo := account.NewMockRepository(ctrl)

	uc := UseCase{
		AccountRepo: mockAccountRepo,
	}

	t.Run("Success", func(t *testing.T) {
		accounts := []*mmodel.Account{{}}
		mockAccountRepo.
			EXPECT().
			FindAll(gomock.Any(), organizationID, ledgerID, nil, limit, page).
			Return(accounts, nil).
			Times(1)
		res, err := uc.AccountRepo.FindAll(context.TODO(), organizationID, ledgerID, nil, limit, page)

		assert.NoError(t, err)
		assert.Len(t, res, 1)
	})

	t.Run("Error", func(t *testing.T) {
		errMsg := "errDatabaseItemNotFound"
		mockAccountRepo.
			EXPECT().
			FindAll(gomock.Any(), organizationID, ledgerID, nil, limit, page).
			Return(nil, errors.New(errMsg)).
			Times(1)
		res, err := uc.AccountRepo.FindAll(context.TODO(), organizationID, ledgerID, nil, limit, page)

		assert.EqualError(t, err, errMsg)
		assert.Nil(t, res)
	})
}
