package usecase

import (
	"context"
	"github.com/niklvrr/Financial-Analytics-Service/internal/domain/model"
	"github.com/niklvrr/Financial-Analytics-Service/internal/domain/request"
	"github.com/niklvrr/Financial-Analytics-Service/internal/domain/response"
)

type BankAccountRepo interface {
	CreateBankAccount(ctx context.Context, account *model.BankAccount) error
	GetBankAccount(ctx context.Context, accountId int64) (*model.BankAccount, error)
	UpdateBankAccount(ctx context.Context, account *model.BankAccount) error
	DeleteBankAccount(ctx context.Context, accountId int64) error
	GetAllBankAccounts(ctx context.Context) ([]*model.BankAccount, error)
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
		Id:      account.ID(),
		Name:    account.Name(),
		Balance: account.Balance(),
	}

	return resp, nil
}

func (s *BankAccountService) UpdateBankAccount(ctx context.Context, req *request.UpdateBankAccountRequest) error {
	account := model.NewBankAccount(req.Id, req.Name, req.Balance)
	if err := s.repo.UpdateBankAccount(ctx, account); err != nil {
		return err
	}

	return nil
}

func (s *BankAccountService) DeleteBankAccount(ctx context.Context, req *request.DeleteBankAccountRequest) error {
	if err := s.repo.DeleteBankAccount(ctx, req.Id); err != nil {
		return err
	}

	return nil
}

func (s *BankAccountService) GetAllBankAccounts(ctx context.Context) ([]*response.BankAccountResponse, error) {
	accounts, err := s.repo.GetAllBankAccounts(ctx)
	if err != nil {
		return nil, err
	}

	resp := make([]*response.BankAccountResponse, len(accounts))
	for i, account := range accounts {
		resp[i] = &response.BankAccountResponse{
			Id:      account.ID(),
			Name:    account.Name(),
			Balance: account.Balance(),
		}
	}

	return resp, nil
}
