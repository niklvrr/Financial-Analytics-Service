package infrastructure

import (
	"context"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/niklvrr/Financial-Analytics-Service/internal/domain/model"
)

const (
	createBankAccountQuery = ``
	getBankAccountQuery    = ``
	updateBankAccountQuery = ``
	deleteBankAccountQuery = ``
)

type BankAccountRepo struct {
	db *pgxpool.Pool
}

func NewBankAccountRepo(db *pgxpool.Pool) *BankAccountRepo {
	return &BankAccountRepo{db: db}
}

func (r *BankAccountRepo) CreateBankAccount(ctx context.Context, account *model.BankAccount) error {
	return nil
}

func (r *BankAccountRepo) GetBankAccount(ctx context.Context, accountId int64) (*model.BankAccount, error) {
	return nil, nil
}

func (r *BankAccountRepo) UpdateBankAccount(ctx context.Context, account *model.BankAccount) error {
	return nil
}

func (r *BankAccountRepo) DeleteBankAccount(ctx context.Context, accountId int64) error {
	return nil
}
