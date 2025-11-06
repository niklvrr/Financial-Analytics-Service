package facade

import (
	"context"
	"github.com/niklvrr/Financial-Analytics-Service/internal/domain/request"
	"github.com/niklvrr/Financial-Analytics-Service/internal/domain/response"
	"github.com/niklvrr/Financial-Analytics-Service/internal/usecase/service"
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
