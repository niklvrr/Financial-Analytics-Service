package infrastructure

import (
	"context"
	"errors"
	"fmt"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/niklvrr/Financial-Analytics-Service/internal/domain/model"
	"time"
)

const (
	createOperationQuery = `
INSERT INTO operations (kind, bank_account_id, amount, date, description, category_id)
VALUES ($1, $2, $3, $4, $5, $6)`
	getOperationQuery = `
SELECT kind, bank_account_id, amount, date, description, category_id
FROM operations
WHERE id = $1`
	updateOperationQuery = `
UPDATE operations
SET amount = $1, description = $2
WHERE id = $3`
	deleteOperationQuery = `
DELETE FROM operations
WHERE id = $1`
	getAllOperationsQuery = `
SELECT kind, bank_account_id, amount, date, description, category_id
FROM operations`
)

var (
	createOperationError  = errors.New("Ошибка создания опреации в базе данных")
	getOperationError     = errors.New("Ошибка чтения операции из базы данных")
	updateOperationError  = errors.New("Ошибка изменения операции в базе данных")
	deleteOperationError  = errors.New("Ошибка удаления операции из базы данных")
	getAllOperationsError = errors.New("Ошибка чтения всех операци из базы данных")
)

type OperationRepo struct {
	db *pgxpool.Pool
}

func NewOperationRepo(db *pgxpool.Pool) *OperationRepo {
	return &OperationRepo{
		db: db,
	}
}

func (r *OperationRepo) CreateOperation(ctx context.Context, op *model.Operation) error {
	err := r.db.QueryRow(
		ctx,
		createCategoryQuery,
		op.Kind(),
		op.BankAccountId(),
		op.Amount(),
		op.Date(),
		op.Description(),
		op.CategoryId(),
	)

	if err != nil {
		return fmt.Errorf("%w: %w", createOperationError, err)
	}
	return nil
}

func (r *OperationRepo) GetOperation(ctx context.Context, id int64) (*model.Operation, error) {
	var (
		op            = &model.Operation{}
		kind          string
		bankAccountId int64
		amount        float64
		date          time.Time
		description   string
		categoryId    int64
	)

	err := r.db.QueryRow(
		ctx,
		getOperationQuery,
		id,
	).Scan(
		&kind,
		&bankAccountId,
		&amount,
		&date,
		&description,
		&categoryId,
	)

	if err != nil {
		return nil, fmt.Errorf("%w: %w", getOperationError, err)
	}

	op.SetID(id)
	op.SetKind(kind)
	op.SetBankAccountId(bankAccountId)
	op.SetAmount(amount)
	op.SetDate(date)
	op.SetDescription(description)
	op.SetCategoryId(categoryId)
	return op, nil
}

func (r *OperationRepo) UpdateOperation(ctx context.Context, op *model.Operation) error {
	cmdTag, err := r.db.Exec(
		ctx,
		updateOperationQuery,
		op.ID(),
		op.Kind(),
		op.BankAccountId(),
		op.Amount(),
		op.Date(),
		op.Description(),
		op.CategoryId(),
	)

	if err != nil {
		return fmt.Errorf("%w: %w", updateOperationError, err)
	}

	if cmdTag.RowsAffected() == 0 {
		return fmt.Errorf("%w: %w", updateOperationError, getOperationError)
	}

	return nil
}

func (r *OperationRepo) DeleteOperation(ctx context.Context, id int64) error {
	cmdTag, err := r.db.Exec(
		ctx,
		deleteOperationQuery,
		id,
	)

	if err != nil {
		return fmt.Errorf("%w: %w", deleteOperationError, err)
	}

	if cmdTag.RowsAffected() == 0 {
		return fmt.Errorf("%w: %w", deleteOperationError, getAllOperationsError)
	}

	return nil
}

func (r *OperationRepo) GetAllOperations(ctx context.Context) ([]*model.Operation, error) {
	rows, err := r.db.Query(ctx, getAllOperationsQuery)
	if err != nil {
		return nil, fmt.Errorf("%w: %w", getAllOperationsError, err)
	}
	defer rows.Close()

	var ops []*model.Operation
	for rows.Next() {
		var (
			op            = &model.Operation{}
			id            int64
			kind          string
			bankAccountId int64
			amount        float64
			date          time.Time
			description   string
			categoryId    int64
		)

		err := rows.Scan(
			&id,
			&kind,
			&bankAccountId,
			&amount,
			&date,
			&description,
			&categoryId,
		)

		if err != nil {
			return nil, fmt.Errorf("%w: %w", getAllOperationsError, err)
		}

		op.SetID(id)
		op.SetKind(kind)
		op.SetBankAccountId(bankAccountId)
		op.SetAmount(amount)
		op.SetDate(date)
		op.SetDescription(description)
		op.SetCategoryId(categoryId)
		ops = append(ops, op)
	}

	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("%w: %w", getAllOperationsError, err)
	}

	return ops, nil
}
