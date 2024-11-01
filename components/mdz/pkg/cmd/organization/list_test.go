package organization

import (
	"bytes"
	"fmt"
	"testing"
	"time"

	"github.com/LerianStudio/midaz/components/mdz/internal/domain/repository"
	"github.com/LerianStudio/midaz/components/mdz/internal/model"
	"github.com/LerianStudio/midaz/components/mdz/pkg/factory"
	"github.com/LerianStudio/midaz/components/mdz/pkg/iostreams"
	"github.com/LerianStudio/midaz/components/mdz/pkg/ptr"
	"github.com/stretchr/testify/assert"
	"go.uber.org/mock/gomock"
	"gotest.tools/golden"
)

func Test_newCmdOrganizationList(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRepo := repository.NewMockOrganization(ctrl)

	orgFactory := factoryOrganizationList{
		factory: &factory.Factory{IOStreams: &iostreams.IOStreams{
			Out: &bytes.Buffer{},
			Err: &bytes.Buffer{},
		}},
		repoOrganization: mockRepo,
	}

	cmd := newCmdOrganizationList(&orgFactory)

	list := model.OrganizationList{
		Items: []model.OrganizationItem{
			{
				ID:                   "123",
				ParentOrganizationID: ptr.StringPtr(""),
				LegalName:            "Test Organization",
				DoingBusinessAs:      "The ledger.io",
				LegalDocument:        "48784548000104",
				Address: model.Address{
					Country: "BR",
				},
				Status: model.Status{
					Description: "Test Ledger",
					Code:        ptr.StringPtr("2123"),
				},
				CreatedAt: time.Date(2024, 10, 31, 11, 31, 22, 369928000, time.UTC),
				UpdatedAt: time.Date(2024, 10, 31, 11, 31, 22, 369928000, time.UTC),
			},
		},
		Limit: 10,
		Page:  1,
	}

	mockRepo.EXPECT().Get(gomock.Any(), gomock.Any()).
		Return(&list, nil)

	err := cmd.Execute()
	assert.NoError(t, err)

	output := orgFactory.factory.IOStreams.Out.(*bytes.Buffer).Bytes()
	fmt.Println(string(output))
	golden.AssertBytes(t, output, "output_list.golden")
}