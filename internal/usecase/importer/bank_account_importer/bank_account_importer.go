package bank_account_importer

import (
	"context"
	"github.com/niklvrr/Financial-Analytics-Service/internal/domain/request"
)

type BankAccountImporter interface {
	Load(path string) ([]*request.CreateBankAccountRequest, error)
	Validate(data []*request.CreateBankAccountRequest) error
	Save(ctx context.Context, data []*request.CreateBankAccountRequest) error
}

type BankAccountTemplate struct {
	Impl BankAccountImporter
}

func (t *BankAccountTemplate) Run(ctx context.Context, path string) error {
	data, err := t.Impl.Load(path)
	if err != nil {
		return err
	}

	err = t.Impl.Validate(data)
	if err != nil {
		return err
	}

	err = t.Impl.Save(ctx, data)
	if err != nil {
		return err
	}
	return nil
}
