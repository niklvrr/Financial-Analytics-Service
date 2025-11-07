package proxy

import (
	"context"
	"github.com/niklvrr/Financial-Analytics-Service/internal/domain/model"
	"github.com/niklvrr/Financial-Analytics-Service/internal/infrastructure/repository"
	"log/slog"
)

type BankAccountProxy struct {
	repository *repository.BankAccountRepo
	cache      map[int64]*model.BankAccount
	log        *slog.Logger
}

func NewBankAccountProxy(repository *repository.BankAccountRepo, log *slog.Logger) *BankAccountProxy {
	return &BankAccountProxy{
		repository: repository,
		cache:      make(map[int64]*model.BankAccount),
		log:        log,
	}
}

func (p *BankAccountProxy) CreateBankAccount(ctx context.Context, account *model.BankAccount) error {
	err := p.repository.CreateBankAccount(ctx, account)
	if err != nil {
		return err
	}
	p.cache[account.ID()] = account
	p.log.Debug("Счет добавлен в кэш")
	return nil
}

func (p *BankAccountProxy) GetBankAccount(ctx context.Context, accountId int64) (*model.BankAccount, error) {
	if account, ok := p.cache[accountId]; ok {
		p.log.Debug("Счет взят из кэша")
		return account, nil
	}

	account, err := p.repository.GetBankAccount(ctx, accountId)
	if err != nil {
		return nil, err
	}
	p.cache[accountId] = account
	p.log.Debug("Счет добавлен в кэш")
	return account, nil
}

func (p *BankAccountProxy) UpdateBankAccount(ctx context.Context, account *model.BankAccount) error {
	err := p.repository.UpdateBankAccount(ctx, account)
	if err != nil {
		return err
	}
	p.cache[account.ID()] = account
	p.log.Debug("Счет добавлен в кэш")
	return nil
}

func (p *BankAccountProxy) DeleteBankAccount(ctx context.Context, accountId int64) error {
	err := p.repository.DeleteBankAccount(ctx, accountId)
	if err != nil {
		return err
	}

	delete(p.cache, accountId)
	p.log.Debug("Счет удален из кэша")
	return nil
}

func (p *BankAccountProxy) GetAllBankAccounts(ctx context.Context) ([]*model.BankAccount, error) {
	return p.repository.GetAllBankAccounts(ctx)
}
