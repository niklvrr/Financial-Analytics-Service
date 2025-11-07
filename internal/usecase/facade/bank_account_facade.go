package facade

import (
	"context"
	"errors"
	"github.com/niklvrr/Financial-Analytics-Service/internal/domain/request"
	"github.com/niklvrr/Financial-Analytics-Service/internal/domain/response"
	"github.com/niklvrr/Financial-Analytics-Service/internal/usecase/importer/bank_account_importer"
	"github.com/niklvrr/Financial-Analytics-Service/internal/usecase/service"
)

var (
	unsupportedFileFormatError = errors.New("Ошибка неподдерживаемый формат файла")
)

type BankAccountFacade struct {
	svc *service.BankAccountService
}

func NewBankAccountFacade(svc *service.BankAccountService) *BankAccountFacade {
	return &BankAccountFacade{
		svc: svc,
	}
}

func (f *BankAccountFacade) CreateBankAccount(ctx context.Context, req *request.CreateBankAccountRequest) error {
	return f.svc.CreateBankAccount(ctx, req)
}

func (f *BankAccountFacade) GetBankAccount(ctx context.Context, req *request.GetBankAccountsRequest) (*response.BankAccountResponse, error) {
	return f.svc.GetBankAccount(ctx, req)
}

func (f *BankAccountFacade) UpdateBankAccount(ctx context.Context, req *request.UpdateBankAccountRequest) error {
	return f.svc.UpdateBankAccount(ctx, req)
}

func (f *BankAccountFacade) DeleteBankAccount(ctx context.Context, req *request.DeleteBankAccountRequest) error {
	return f.svc.DeleteBankAccount(ctx, req)
}

func (f *BankAccountFacade) GetAllBankAccounts(ctx context.Context) ([]*response.BankAccountResponse, error) {
	return f.svc.GetAllBankAccounts(ctx)
}

func (f *BankAccountFacade) ImportDataFromFile(ctx context.Context, path, format string) error {
	var importer bank_account_importer.BankAccountImporter
	switch format {
	case "csv":
		importer = bank_account_importer.NewBankAccountCSVImporter(f.svc)
	case "json":
		importer = bank_account_importer.NewBankAccountJSONImporter(f.svc)
	case "yaml":
		importer = bank_account_importer.NewBankAccountYamlImporter(f.svc)
	default:
		return unsupportedFileFormatError
	}

	t := bank_account_importer.Template{
		Impl: importer,
	}
	return t.Run(ctx, path)
}
