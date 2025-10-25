package infrastructure

import (
	"context"
	"errors"
	"fmt"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/niklvrr/Financial-Analytics-Service/internal/domain/model"
)

const (
	createBankAccountQuery = `
INSERT INTO bank_accounts (name, balance)
VALUES ($1, $2)
RETURNING id`
	getBankAccountQuery = `
SELECT name, balance
FROM bank_accounts
WHERE id = $1`
	updateBankAccountQuery = `
UPDATE bank_accounts
SET name = $1, balance = $2
WHERE id = $3`
	deleteBankAccountQuery = `
DELETE FROM bank_accounts
WHERE id = $1`
	getAllBankAccountsQuery = `
SELECT id, name, balance
FROM bank_accounts`
)

var (
	createBankAccountError  = errors.New("Ошибка создания банковского счета в базу данных")
	getBankAccountError     = errors.New("Ошибка чтения банковского счета из базы данных")
	updateBankAccountError  = errors.New("Ошибка изменения банковского счета в базу данных")
	deleteBankAccountError  = errors.New("Ошибка удаления банковского счета из базы данных")
	getAllBankAccountsError = errors.New("Ошибка чтения всех банковских счетов из базы данных")
)

type BankAccountRepo struct {
	db *pgxpool.Pool
}

func NewBankAccountRepo(db *pgxpool.Pool) *BankAccountRepo {
	return &BankAccountRepo{db: db}
}

func (r *BankAccountRepo) CreateBankAccount(ctx context.Context, account *model.BankAccount) error {
	err := r.db.QueryRow(
		ctx,
		createBankAccountQuery,
		account.Name,
		account.Balance,
	)

	if err != nil {
		return fmt.Errorf("%w: %w", createBankAccountError, err)
	}

	return nil
}

func (r *BankAccountRepo) GetBankAccount(ctx context.Context, accountId int64) (*model.BankAccount, error) {
	var (
		account = &model.BankAccount{}
		name    string
		balance float64
	)
	err := r.db.QueryRow(ctx, getBankAccountQuery, accountId).Scan(
		&name,
		&balance,
	)
	if err != nil {
		return nil, fmt.Errorf("%w: %w", getBankAccountError, err)
	}

	account.SetID(accountId)
	account.SetName(name)
	account.SetBalance(balance)
	return account, nil
}

func (r *BankAccountRepo) UpdateBankAccount(ctx context.Context, account *model.BankAccount) error {
	var (
		name    string
		balance float64
	)

	cmdTag, err := r.db.Exec(
		ctx, updateBankAccountQuery,
		&name, &balance,
		account.ID,
	)

	if err != nil {
		return fmt.Errorf("%w: %w", updateBankAccountError, err)
	}

	if cmdTag.RowsAffected() == 0 {
		return fmt.Errorf("%w: %w", updateBankAccountError, getBankAccountError)
	}

	account.SetName(name)
	account.SetBalance(balance)
	return nil
}

func (r *BankAccountRepo) DeleteBankAccount(ctx context.Context, accountId int64) error {
	cmdTag, err := r.db.Exec(ctx, deleteBankAccountQuery, accountId)
	if err != nil {
		return fmt.Errorf("%w: %w", deleteBankAccountError, err)
	}

	if cmdTag.RowsAffected() == 0 {
		return fmt.Errorf("%w: %w", deleteBankAccountError, getBankAccountError)
	}
	return nil
}

func (r *BankAccountRepo) GetAllBankAccounts(ctx context.Context) ([]*model.BankAccount, error) {
	rows, err := r.db.Query(ctx, getAllBankAccountsQuery)
	if err != nil {
		return nil, fmt.Errorf("%w: %w", getAllBankAccountsError, err)
	}
	defer rows.Close()

	var accounts []*model.BankAccount
	for rows.Next() {
		var (
			account = &model.BankAccount{}
			id      int64
			name    string
			balance float64
		)
		err := rows.Scan(
			&id,
			&name,
			&balance,
		)

		if err != nil {
			return nil, fmt.Errorf("%w: %w", getAllBankAccountsError, err)
		}
		account.SetID(id)
		account.SetName(name)
		account.SetBalance(balance)

		accounts = append(accounts, account)
	}

	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("%w: %w", getAllBankAccountsError, err)
	}
	return accounts, nil
}
