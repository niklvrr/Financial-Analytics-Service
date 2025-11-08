package bank_account_importer

import (
	"context"
	"encoding/json"
	"github.com/niklvrr/Financial-Analytics-Service/internal/domain/request"
	"github.com/niklvrr/Financial-Analytics-Service/internal/usecase/service"
	"github.com/niklvrr/Financial-Analytics-Service/pkg/utils"
	"os"
)

type JSONBankAccountImporter struct {
	svc *service.BankAccountService
}

func NewJSONBankAccountImporter(svc *service.BankAccountService) *JSONBankAccountImporter {
	return &JSONBankAccountImporter{
		svc: svc,
	}
}

func (b *JSONBankAccountImporter) Load(path string) ([]*request.CreateBankAccountRequest, error) {
	data, err := os.ReadFile(path)
	if err != nil {
		panic(err)
	}

	var reqs []*request.CreateBankAccountRequest
	err = json.Unmarshal(data, &reqs)
	if err != nil {
		return nil, err
	}
	return reqs, nil
}

func (b *JSONBankAccountImporter) Validate(data []*request.CreateBankAccountRequest) error {
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

func (b *JSONBankAccountImporter) Save(ctx context.Context, data []*request.CreateBankAccountRequest) error {
	for _, req := range data {
		err := b.svc.CreateBankAccount(ctx, req)
		if err != nil {
			return err
		}
	}
	return nil
}
