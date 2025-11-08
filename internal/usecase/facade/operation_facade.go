package facade

import (
	"context"
	"github.com/niklvrr/Financial-Analytics-Service/internal/domain/request"
	"github.com/niklvrr/Financial-Analytics-Service/internal/domain/response"
	"github.com/niklvrr/Financial-Analytics-Service/internal/usecase/importer/operation_importer"
	"github.com/niklvrr/Financial-Analytics-Service/internal/usecase/service"
)

type OperationFacade struct {
	svc *service.OperationService
}

func NewOperationFacade(svc *service.OperationService) *OperationFacade {
	return &OperationFacade{svc: svc}
}

func (f *OperationFacade) CreateOperation(ctx context.Context, req *request.CreateOperationRequest) error {
	return f.svc.CreateOperation(ctx, req)
}

func (f *OperationFacade) GetOperation(ctx context.Context, req *request.GetOperationRequest) (*response.OperationResponse, error) {
	return f.svc.GetOperation(ctx, req)
}

func (f *OperationFacade) UpdateOperation(ctx context.Context, req *request.UpdateOperationRequest) error {
	return f.svc.UpdateOperation(ctx, req)
}

func (f *OperationFacade) DeleteOperation(ctx context.Context, req *request.DeleteOperationRequest) error {
	return f.svc.DeleteOperation(ctx, req)
}

func (f *OperationFacade) GetAllOperations(ctx context.Context) ([]*response.OperationResponse, error) {
	return f.svc.GetAllOperations(ctx)
}

func (f *OperationFacade) ImportOperationFromFile(ctx context.Context, path, format string) error {
	var importer operation_importer.OperationImporter
	switch format {
	case ".csv":
		importer = operation_importer.NewCSVOperationImporter(f.svc)
	case ".json":
		importer = operation_importer.NewJSONOperationImporter(f.svc)
	case ".yaml":
		importer = operation_importer.NewYamlOperationImporter(f.svc)
	default:
		return unsupportedFileFormatError
	}

	t := operation_importer.OperationTemplate{
		Impl: importer,
	}
	return t.Run(ctx, path)
}
