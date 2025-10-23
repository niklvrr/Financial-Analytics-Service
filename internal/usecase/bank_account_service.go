package usecase

import (
	"context"
	"github.com/niklvrr/FinancialAnalyticsService/internal/domain/model"
	"github.com/niklvrr/FinancialAnalyticsService/internal/domain/request"
	"github.com/niklvrr/FinancialAnalyticsService/internal/domain/response"
)

type BankAccountRepo interface {
	CreateBankAccount(ctx context.Context, account *model.BankAccount) error
	GetBankAccount(ctx context.Context, accountId int64) (*model.BankAccount, error)
	UpdateBankAccount(ctx context.Context, account *model.BankAccount) error
	DeleteBankAccount(ctx context.Context, accountId int64) error
}

type BankAccountService struct {
	repo BankAccountRepo
}

func NewBankAccountService(repo BankAccountRepo) *BankAccountService {
	return &BankAccountService{repo: repo}
}

func (s *BankAccountService) CreateBankAccount(ctx context.Context, req *request.CreateBankAccountRequest) error {
	account := model.NewBankAccount(0, req.Name, 0)
	if err := s.repo.CreateBankAccount(ctx, account); err != nil {
		return err
	}

	return nil
}

func (s *BankAccountService) GetBankAccount(ctx context.Context, req *request.GetBankAccountsRequest) (*response.BankAccountResponse, error) {
	account, err := s.repo.GetBankAccount(ctx, req.Id)
	if err != nil {
		return nil, err
	}

	resp := &response.BankAccountResponse{
		Id: account.
	}

	return resp, nil
}