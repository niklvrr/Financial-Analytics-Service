package bank_account_importer

import (
	"context"
	"github.com/gocarina/gocsv"
	"github.com/niklvrr/Financial-Analytics-Service/internal/domain/request"
	"github.com/niklvrr/Financial-Analytics-Service/internal/usecase/service"
	"github.com/niklvrr/Financial-Analytics-Service/pkg/utils"
	"os"
)

type CSVBankAccountImporter struct {
	svc *service.BankAccountService
}

func NewCSVBankAccountImporter(svc *service.BankAccountService) *CSVBankAccountImporter {
	return &CSVBankAccountImporter{
		svc: svc,
	}
}

func (b *CSVBankAccountImporter) Load(path string) ([]*request.CreateBankAccountRequest, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var reqs []*request.CreateBankAccountRequest
	if err := gocsv.UnmarshalFile(file, &reqs); err != nil {
		return nil, err
	}
	return reqs, nil
}

func (b *CSVBankAccountImporter) Validate(data []*request.CreateBankAccountRequest) error {
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

func (b *CSVBankAccountImporter) Save(ctx context.Context, data []*request.CreateBankAccountRequest) error {
	for _, req := range data {
		err := b.svc.CreateBankAccount(ctx, req)
		if err != nil {
			return err
		}
	}
	return nil
}
