package usecase

import (
	"context"
	"github.com/niklvrr/Financial-Analytics-Service/internal/domain/model"
	"github.com/niklvrr/Financial-Analytics-Service/internal/domain/request"
	"github.com/niklvrr/Financial-Analytics-Service/internal/domain/response"
	"time"
)

type OperationRepo interface {
	CreateOperation(ctx context.Context, op *model.Operation) error
	GetOperation(ctx context.Context, id int64) (*model.Operation, error)
	UpdateOperation(ctx context.Context, op *model.Operation) error
	DeleteOperation(ctx context.Context, id int64) error
	GetAllOperations(ctx context.Context) ([]*model.Operation, error)
}

type OperationService struct {
	operationRepo OperationRepo
}

func NewOperationService(operationRepo OperationRepo) *OperationService {
	return &OperationService{
		operationRepo: operationRepo,
	}
}

func (s *OperationService) CreateOperation(ctx context.Context, req *request.CreateOperationRequest) error {
	op := model.NewOperation(
		0,
		req.Kind,
		req.BankAccountId,
		req.Amount,
		time.Now(),
		req.Description,
		req.CategoryId,
	)

	err := s.operationRepo.CreateOperation(ctx, op)
	if err != nil {
		return err
	}

	return nil
}

func (s *OperationService) GetOperation(ctx context.Context, req *request.GetOperationRequest) (*response.OperationResponse, error) {
	op, err := s.operationRepo.GetOperation(ctx, req.Id)
	if err != nil {
		return nil, err
	}

	resp := &response.OperationResponse{
		Id:            op.ID(),
		Kind:          op.Kind(),
		BankAccountId: op.BankAccountId(),
		Amount:        op.Amount(),
		Description:   op.Description(),
		CategoryId:    op.CategoryId(),
	}

	return resp, nil
}

func (s *OperationService) UpdateOperation(ctx context.Context, req *request.UpdateOperationRequest) error {
	op := model.NewOperation(
		req.Id,
		req.Kind,
		req.BankAccountId,
		req.Amount,
		req.Date,
		req.Description,
		req.CategoryId,
	)

	err := s.operationRepo.UpdateOperation(ctx, op)
	if err != nil {
		return err
	}

	return nil
}

func (s *OperationService) DeleteOperation(ctx context.Context, req *request.DeleteOperationRequest) error {
	err := s.operationRepo.DeleteOperation(ctx, req.Id)
	if err != nil {
		return err
	}

	return nil
}

func (s *OperationService) GetAllOperations(ctx context.Context) ([]*response.OperationResponse, error) {
	ops, err := s.operationRepo.GetAllOperations(ctx)
	if err != nil {
		return nil, err
	}

	resp := make([]*response.OperationResponse, len(ops))
	for i, op := range ops {
		resp[i] = &response.OperationResponse{
			Id:            op.ID(),
			Kind:          op.Kind(),
			BankAccountId: op.BankAccountId(),
			Amount:        op.Amount(),
			Description:   op.Description(),
			CategoryId:    op.CategoryId(),
		}
	}

	return resp, nil
}
