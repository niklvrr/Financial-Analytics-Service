package facade

import (
	"context"
	"github.com/niklvrr/Financial-Analytics-Service/internal/domain/request"
	"github.com/niklvrr/Financial-Analytics-Service/internal/domain/response"
	"github.com/niklvrr/Financial-Analytics-Service/internal/usecase/exporter"
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

func (f *OperationFacade) Export(ctx context.Context, params exporter.ExportParams) (*exporter.Report, error) {
	strategy, err := exporter.NewStrategy(params.Strategy, "operation")
	if err != nil {
		return nil, err
	}

	switch s := strategy.(type) {
	case *exporter.FullExportStrategy:
		s.SetService(f.svc)
	case *exporter.ByAccountStrategy:
		s.SetService(f.svc)
	case *exporter.ByCategoryStrategy:
		s.SetService(f.svc)
	case *exporter.DateRangeStrategy:
		s.SetService(f.svc)
	}

	data, err := strategy.Collect(ctx, params)
	if err != nil {
		return nil, err
	}

	if len(data) == 0 {
		return nil, exporter.ErrEmptyData
	}

	builder, err := exporter.NewBuilder(params.Format)
	if err != nil {
		return nil, err
	}

	if err := builder.Begin(ctx, "Операции"); err != nil {
		return nil, err
	}

	headers := strategy.GetHeaders()
	if err := builder.AddHeader(ctx, headers...); err != nil {
		return nil, err
	}

	for _, row := range data {
		values := make([]string, 0, len(headers))
		for _, header := range headers {
			values = append(values, row[header])
		}
		if err := builder.AddRow(ctx, values...); err != nil {
			return nil, err
		}
	}

	report, err := builder.End(ctx)
	if err != nil {
		return nil, err
	}

	return report, nil
}
