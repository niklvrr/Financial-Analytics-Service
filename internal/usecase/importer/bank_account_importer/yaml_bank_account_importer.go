package bank_account_importer

import (
	"context"
	"github.com/niklvrr/Financial-Analytics-Service/internal/domain/request"
	"github.com/niklvrr/Financial-Analytics-Service/internal/usecase/service"
	"github.com/niklvrr/Financial-Analytics-Service/pkg/utils"
	"gopkg.in/yaml.v3"
	"os"
)

type BankAccountYamlImporter struct {
	svc *service.BankAccountService
}

func NewBankAccountYamlImporter(svc *service.BankAccountService) *BankAccountYamlImporter {
	return &BankAccountYamlImporter{
		svc: svc,
	}
}

func (b *BankAccountYamlImporter) Load(path string) ([]*request.CreateBankAccountRequest, error) {
	data, err := os.ReadFile("config.yaml")
	if err != nil {
		return nil, err
	}

	var reqs []*request.CreateBankAccountRequest
	err = yaml.Unmarshal(data, &reqs)
	if err != nil {
		return nil, err
	}
	return reqs, nil
}

func (b *BankAccountYamlImporter) Validate(data []*request.CreateBankAccountRequest) error {
	for _, req := range data {
		err := utils.ValidateString(req.Name)
		if err != nil {
			return err
		}

		err = utils.ValidateFloat(req.Balance)
		if err != nil {
			return err
		}
	}
	return nil
}

func (b *BankAccountYamlImporter) Save(ctx context.Context, data []*request.CreateBankAccountRequest) error {
	for _, req := range data {
		err := b.svc.CreateBankAccount(ctx, req)
		if err != nil {
			return err
		}
	}
	return nil
}
