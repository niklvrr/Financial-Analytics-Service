package proxy

import (
	"context"
	"github.com/niklvrr/Financial-Analytics-Service/internal/domain/model"
	"github.com/niklvrr/Financial-Analytics-Service/internal/infrastructure/repository"
	"log/slog"
)

type CategoryProxy struct {
	repository *repository.CategoryRepo
	cache      map[int64]*model.Category
	log        *slog.Logger
}

func NewCategoryProxy(repository *repository.CategoryRepo, log *slog.Logger) *CategoryProxy {
	return &CategoryProxy{
		repository: repository,
		cache:      make(map[int64]*model.Category),
		log:        log,
	}
}

func (p *CategoryProxy) CreateCategory(ctx context.Context, c *model.Category) error {
	err := p.repository.CreateCategory(ctx, c)
	if err != nil {
		return err
	}
	p.cache[c.ID()] = c
	p.log.Debug("Категория добавлена в кэш")
	return nil
}

func (p *CategoryProxy) GetCategory(ctx context.Context, id int64) (*model.Category, error) {
	if c, ok := p.cache[id]; ok {
		p.log.Debug("Категория взята из кэша")
		return c, nil
	}

	c, err := p.repository.GetCategory(ctx, id)
	if err != nil {
		return nil, err
	}
	p.cache[id] = c
	p.log.Debug("Категория добавлена в кэш")
	return c, nil
}

func (p *CategoryProxy) UpdateCategory(ctx context.Context, c *model.Category) error {
	err := p.repository.UpdateCategory(ctx, c)
	if err != nil {
		return err
	}
	p.cache[c.ID()] = c
	p.log.Debug("Категория добавлена в кэш")
	return nil
}

func (p *CategoryProxy) DeleteCategory(ctx context.Context, id int64) error {
	err := p.repository.DeleteCategory(ctx, id)
	if err != nil {
		return err
	}
	delete(p.cache, id)
	p.log.Debug("Категория удалена из кэша")
	return nil
}
func (p *CategoryProxy) GetAllCategories(ctx context.Context) ([]*model.Category, error) {
	return p.repository.GetAllCategories(ctx)
}
