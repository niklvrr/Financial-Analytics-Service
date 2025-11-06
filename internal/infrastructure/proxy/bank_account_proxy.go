package proxy

import (
	"context"
	"github.com/niklvrr/Financial-Analytics-Service/internal/domain/model"
	"github.com/niklvrr/Financial-Analytics-Service/internal/infrastructure/repository"
)

type BankAccountProxy struct {
	repository *repository.BankAccountRepo
	cache      map[int64]*model.BankAccount
}

func NewBankAccountProxy(repository *repository.BankAccountRepo) *BankAccountProxy {
	return &BankAccountProxy{
		repository: repository,
		cache:      make(map[int64]*model.BankAccount),
	}
}

func (p *BankAccountProxy) CreateBankAccount(ctx context.Context, account *model.BankAccount) error {
	err := p.repository.CreateBankAccount(ctx, account)
	if err != nil {
		return err
	}
	p.cache[account.ID()] = account
	return nil
}

func (p *BankAccountProxy) GetBankAccount(ctx context.Context, accountId int64) (*model.BankAccount, error) {
	if account, ok := p.cache[accountId]; ok {
		return account, nil
	}

	account, err := p.repository.GetBankAccount(ctx, accountId)
	if err != nil {
		return nil, err
	}
	p.cache[accountId] = account
	return account, nil
}

func (p *BankAccountProxy) UpdateBankAccount(ctx context.Context, account *model.BankAccount) error {
	err := p.repository.UpdateBankAccount(ctx, account)
	if err != nil {
		return err
	}
	p.cache[account.ID()] = account
	return nil
}

func (p *BankAccountProxy) DeleteBankAccount(ctx context.Context, accountId int64) error {
	err := p.repository.DeleteBankAccount(ctx, accountId)
	if err != nil {
		return err
	}

	delete(p.cache, accountId)
	return nil
}

func (p *BankAccountProxy) GetAllBankAccounts(ctx context.Context) ([]*model.BankAccount, error) {
	return p.repository.GetAllBankAccounts(ctx)
}
