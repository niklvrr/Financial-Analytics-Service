package facade

import (
	"context"
	"fmt"
	"github.com/niklvrr/Financial-Analytics-Service/internal/domain/request"
	"github.com/niklvrr/Financial-Analytics-Service/internal/domain/response"
	"github.com/niklvrr/Financial-Analytics-Service/internal/usecase/exporter"
	"github.com/niklvrr/Financial-Analytics-Service/internal/usecase/importer/category_importer"
	"github.com/niklvrr/Financial-Analytics-Service/internal/usecase/service"
)

type CategoryFacade struct {
	svc *service.CategoryService
}

func NewCategoryFacade(svc *service.CategoryService) *CategoryFacade {
	return &CategoryFacade{
		svc: svc,
	}
}

func (f *CategoryFacade) CreateCategory(ctx context.Context, req *request.CreateCategoryRequest) error {
	return f.svc.CreateCategory(ctx, req)
}

func (f *CategoryFacade) GetCategory(ctx context.Context, req *request.GetCategoryRequest) (*response.CategoryResponse, error) {
	return f.svc.GetCategory(ctx, req)
}

func (f *CategoryFacade) UpdateCategory(ctx context.Context, req *request.UpdateCategoryRequest) error {
	return f.svc.UpdateCategory(ctx, req)
}

func (f *CategoryFacade) DeleteCategory(ctx context.Context, req *request.DeleteCategoryRequest) error {
	return f.svc.DeleteCategory(ctx, req)
}

func (f *CategoryFacade) GetAllCategories(ctx context.Context) ([]*response.CategoryResponse, error) {
	return f.svc.GetAllCategories(ctx)
}

func (f *CategoryFacade) ImportCategoryFromFile(ctx context.Context, path, format string) error {
	var importer category_importer.CategoryImporter
	switch format {
	case ".csv":
		importer = category_importer.NewCSVCategoryImporter(f.svc)
	case ".json":
		importer = category_importer.NewJSONCategoryImporter(f.svc)
	case ".yaml":
		importer = category_importer.NewYamlCategoryImporter(f.svc)
	default:
		return unsupportedFileFormatError
	}

	t := category_importer.CategoryTemplate{
		Impl: importer,
	}
	err := t.Run(ctx, path)
	if err != nil {
		return err
	}
	return nil
}

func (f *CategoryFacade) Export(ctx context.Context, params exporter.ExportParams) (*exporter.Report, error) {
	if params.Strategy != "full" {
		return nil, fmt.Errorf("для категорий поддерживается только стратегия 'full'")
	}

	strategy, err := exporter.NewStrategy(params.Strategy, "category")
	if err != nil {
		return nil, err
	}

	if fullStrategy, ok := strategy.(*exporter.FullExportStrategy); ok {
		fullStrategy.SetService(f.svc)
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

	if err := builder.Begin(ctx, "Категории"); err != nil {
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
