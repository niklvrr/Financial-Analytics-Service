package proxy

import (
	"context"
	"github.com/niklvrr/Financial-Analytics-Service/internal/domain/model"
	"github.com/niklvrr/Financial-Analytics-Service/internal/infrastructure/repository"
	"log/slog"
)

type OperationProxy struct {
	repo  *repository.OperationRepo
	cache map[int64]*model.Operation
	log   *slog.Logger
}

func NewOperationProxy(repository *repository.OperationRepo, log *slog.Logger) *OperationProxy {
	return &OperationProxy{
		repo:  repository,
		cache: make(map[int64]*model.Operation),
		log:   log,
	}
}

func (p *OperationProxy) CreateOperation(ctx context.Context, op *model.Operation) error {
	err := p.repo.CreateOperation(ctx, op)
	if err != nil {
		return err
	}
	p.cache[op.ID()] = op
	p.log.Debug("Операция добавлена в кэш")
	return nil
}

func (p *OperationProxy) GetOperation(ctx context.Context, id int64) (*model.Operation, error) {
	if op, ok := p.cache[id]; ok {
		p.log.Debug("Операция взята из кэша")
		return op, nil
	}

	op, err := p.repo.GetOperation(ctx, id)
	if err != nil {
		return nil, err
	}
	p.cache[id] = op
	p.log.Debug("Операция добавлена в кэш")
	return op, nil
}

func (p *OperationProxy) UpdateOperation(ctx context.Context, op *model.Operation) error {
	err := p.repo.UpdateOperation(ctx, op)
	if err != nil {
		return err
	}
	p.cache[op.ID()] = op
	p.log.Debug("Операция добавлена в кэш")
	return nil
}

func (p *OperationProxy) DeleteOperation(ctx context.Context, id int64) error {
	err := p.repo.DeleteOperation(ctx, id)
	if err != nil {
		return err
	}
	delete(p.cache, id)
	p.log.Debug("Операция удалена из кэша")
	return nil
}

func (p *OperationProxy) GetAllOperations(ctx context.Context) ([]*model.Operation, error) {
	return p.repo.GetAllOperations(ctx)
}
