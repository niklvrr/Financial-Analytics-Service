package infrastructure

import (
	"context"
	"errors"
	"fmt"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/niklvrr/Financial-Analytics-Service/internal/domain/model"
)

const (
	createCategoryQuery = `
INSERT INTO categories (kind, name)
VALUES ($1, $2)`
	getCategoryQuery = `
SELECT kind, name
FROM categories
WHERE id = $1`
	updateCategoryQuery = `
UPDATE categories
SET kind = $1, name = $2
WHERE id = $3`
	deleteCategoryQuery = `
DELETE FROM categories
WHERE id = $1`
	getAllCategoriesQuery = `
SELECT id, kind, name
FROM categories`
)

var (
	createCategoryError   = errors.New("Ошибка создания категории в базу данных")
	getCategoryError      = errors.New("Ошибка чтения категории из базы данных")
	updateCategoryError   = errors.New("Ошибка изменения категории в базу данных")
	deleteCategoryError   = errors.New("Ошибка удаления категории из базы данных")
	getAllCategoriesError = errors.New("Ошибка чтения всех категорий из базы данных")
)

type CategoryRepo struct {
	db *pgxpool.Pool
}

func NewCategoryRepo(db *pgxpool.Pool) *CategoryRepo {
	return &CategoryRepo{
		db: db,
	}
}

func (r *CategoryRepo) CreateCategory(ctx context.Context, c *model.Category) error {
	err := r.db.QueryRow(
		ctx,
		createCategoryQuery,
		c.Kind,
		c.Name,
	)

	if err != nil {
		return fmt.Errorf("%w: %w", createCategoryError, err)
	}

	return nil
}

func (r *CategoryRepo) GetCategory(ctx context.Context, id int64) (*model.Category, error) {
	var (
		c    = &model.Category{}
		kind string
		name string
	)
	err := r.db.QueryRow(
		ctx,
		getCategoryQuery,
		id,
	).Scan(&kind, &name)

	if err != nil {
		return nil, fmt.Errorf("%w: %w", getCategoryError, err)
	}
	c.SetID(id)
	c.SetKind(kind)
	c.SetName(name)
	return c, nil
}

func (r *CategoryRepo) UpdateCategory(ctx context.Context, c *model.Category) error {
	cmdTag, err := r.db.Exec(
		ctx,
		updateCategoryQuery,
		c.ID(),
		c.Kind(),
		c.Name(),
	)

	if err != nil {
		return fmt.Errorf("%w: %w", updateCategoryError, err)
	}

	if cmdTag.RowsAffected() == 0 {
		return fmt.Errorf("%w: %w", updateCategoryError, getCategoryError)
	}

	return nil
}

func (r *CategoryRepo) DeleteCategory(ctx context.Context, id int64) error {
	cmdTag, err := r.db.Exec(
		ctx,
		deleteCategoryQuery,
		id,
	)

	if err != nil {
		return fmt.Errorf("%w: %w", deleteCategoryError, err)
	}

	if cmdTag.RowsAffected() == 0 {
		return fmt.Errorf("%w: %w", deleteCategoryError, getCategoryError)
	}

	return nil
}

func (r *CategoryRepo) GetAllCategories(ctx context.Context) ([]*model.Category, error) {
	rows, err := r.db.Query(ctx, getAllCategoriesQuery)
	if err != nil {
		return nil, fmt.Errorf("%w: %w", getAllCategoriesError, err)
	}
	defer rows.Close()

	var categories []*model.Category
	for rows.Next() {
		var (
			c    = &model.Category{}
			id   int64
			kind string
			name string
		)
		if err := rows.Scan(&id, &kind, &name); err != nil {
			return nil, fmt.Errorf("%w: %w", getAllCategoriesError, err)
		}
		c.SetID(id)
		c.SetKind(kind)
		c.SetName(name)

		categories = append(categories, c)
	}

	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("%w: %w", getAllCategoriesError, err)
	}
	return categories, nil
}
