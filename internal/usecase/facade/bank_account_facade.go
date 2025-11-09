package facade

import (
	"context"
	"errors"
	"fmt"
	"github.com/niklvrr/Financial-Analytics-Service/internal/domain/request"
	"github.com/niklvrr/Financial-Analytics-Service/internal/domain/response"
	"github.com/niklvrr/Financial-Analytics-Service/internal/usecase/exporter"
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

func (f *BankAccountFacade) ImportBankAccountsFromFile(ctx context.Context, path, format string) error {
	var importer bank_account_importer.BankAccountImporter
	switch format {
	case ".csv":
		importer = bank_account_importer.NewCSVBankAccountImporter(f.svc)
	case ".json":
		importer = bank_account_importer.NewJSONBankAccountImporter(f.svc)
	case ".yaml":
		importer = bank_account_importer.NewYamlBankAccountImporter(f.svc)
	default:
		return unsupportedFileFormatError
	}

	t := bank_account_importer.BankAccountTemplate{
		Impl: importer,
	}
	return t.Run(ctx, path)
}

func (f *BankAccountFacade) Export(ctx context.Context, params exporter.ExportParams) (*exporter.Report, error) {
	if params.Strategy != "full" {
		return nil, fmt.Errorf("для банковских счетов поддерживается только стратегия 'full'")
	}

	strategy, err := exporter.NewStrategy(params.Strategy, "bank_account")
	if err != nil {
		return nil, err
	}

	if fullStrategy, ok := strategy.(*exporter.FullExportStrategy); ok {
		fullStrategy.SetService(f.svc)
	}

	data, err := strategy.Collect(ctx, params)
	if err != nil {
		return nil, err
	}

	if len(data) == 0 {
		return nil, exporter.ErrEmptyData
	}

	builder, err := exporter.NewBuilder(params.Format)
	if err != nil {
		return nil, err
	}

	if err := builder.Begin(ctx, "Банковские счета"); err != nil {
		return nil, err
	}

	headers := strategy.GetHeaders()
	if err := builder.AddHeader(ctx, headers...); err != nil {
		return nil, err
	}

	for _, row := range data {
		values := make([]string, 0, len(headers))
		for _, header := range headers {
			values = append(values, row[header])
		}
		if err := builder.AddRow(ctx, values...); err != nil {
			return nil, err
		}
	}

	report, err := builder.End(ctx)
	if err != nil {
		return nil, err
	}

	return report, nil
}
