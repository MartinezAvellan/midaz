package account

import (
	"bytes"
	"testing"

	"github.com/LerianStudio/midaz/common/mmodel"
	"github.com/LerianStudio/midaz/components/mdz/internal/domain/repository"
	"github.com/LerianStudio/midaz/components/mdz/pkg/factory"
	"github.com/LerianStudio/midaz/components/mdz/pkg/iostreams"
	"github.com/LerianStudio/midaz/components/mdz/pkg/ptr"
	"github.com/stretchr/testify/assert"
	"go.uber.org/mock/gomock"
)

func Test_newCmdLedgerUpdate(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRepo := repository.NewMockAccount(ctrl)

	orgFactory := factoryAccountUpdate{
		factory: &factory.Factory{IOStreams: &iostreams.IOStreams{
			Out: &bytes.Buffer{},
			Err: &bytes.Buffer{},
		}},
		repoAccount: mockRepo,
		tuiInput: func(message string) (string, error) {
			return "name", nil
		},
		flagsUpdate: flagsUpdate{
			Name:              "Test Organization",
			StatusCode:        "BLOCKED",
			StatusDescription: "Teste BLOCKED Ledger",
			Metadata:          "{\"chave1\": \"valor1\", \"chave2\": 2, \"chave3\": true}",
		},
	}

	cmd := newCmdAccountUpdate(&orgFactory)
	cmd.SetArgs([]string{
		"--organization-id", "123",
		"--ledger-id", "321",
		"--portfolio-id", "412",
		"--account-id", "777",
		"--name", "Test Organization",
		"--status-code", "BLOCKED",
		"--status-description", "Teste BLOCKED Ledger",
		"--metadata", "{\"chave1\": \"valor1\", \"chave2\": 2, \"chave3\": true}",
	})

	metadata := map[string]any{
		"chave1": "valor1",
		"chave2": 2,
		"chave3": true,
	}

	got := &mmodel.Account{
		ID:   "777",
		Name: "Test Organization",
		Status: mmodel.Status{
			Code:        "BLOCKED",
			Description: ptr.StringPtr("Teste BLOCKED Ledger"),
		},
		Metadata: metadata,
	}

	mockRepo.EXPECT().Update(gomock.Any(), gomock.Any(), gomock.Any(), gomock.Any(), gomock.Any()).Return(got, nil)

	err := cmd.Execute()
	assert.NoError(t, err)

	output := orgFactory.factory.IOStreams.Out.(*bytes.Buffer).String()
	assert.Contains(t, output, "The Account ID 777 has been successfully updated.")
}
