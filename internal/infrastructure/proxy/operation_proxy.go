package proxy

import (
	"context"
	"github.com/niklvrr/Financial-Analytics-Service/internal/domain/model"
	"github.com/niklvrr/Financial-Analytics-Service/internal/infrastructure/repository"
)

type OperationProxy struct {
	repo  *repository.OperationRepo
	cache map[int64]*model.Operation
}

func NewOperationProxy(repository *repository.OperationRepo) *OperationProxy {
	return &OperationProxy{
		repo:  repository,
		cache: make(map[int64]*model.Operation),
	}
}

func (p *OperationProxy) CreateOperation(ctx context.Context, op *model.Operation) error {
	err := p.repo.CreateOperation(ctx, op)
	if err != nil {
		return err
	}
	p.cache[op.ID()] = op
	return nil
}

func (p *OperationProxy) GetOperation(ctx context.Context, id int64) (*model.Operation, error) {
	if op, ok := p.cache[id]; ok {
		return op, nil
	}

	op, err := p.repo.GetOperation(ctx, id)
	if err != nil {
		return nil, err
	}
	p.cache[id] = op
	return op, nil
}

func (p *OperationProxy) UpdateOperation(ctx context.Context, op *model.Operation) error {
	err := p.repo.UpdateOperation(ctx, op)
	if err != nil {
		return err
	}
	p.cache[op.ID()] = op
	return nil
}

func (p *OperationProxy) DeleteOperation(ctx context.Context, id int64) error {
	err := p.repo.DeleteOperation(ctx, id)
	if err != nil {
		return err
	}
	delete(p.cache, id)
	return nil
}

func (p *OperationProxy) GetAllOperations(ctx context.Context) ([]*model.Operation, error) {
	return p.repo.GetAllOperations(ctx)
}
