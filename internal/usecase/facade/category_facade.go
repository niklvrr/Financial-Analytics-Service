package facade

import (
	"context"
	"github.com/niklvrr/Financial-Analytics-Service/internal/domain/request"
	"github.com/niklvrr/Financial-Analytics-Service/internal/domain/response"
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

func (f *CategoryFacade) DeleteCategory(ctx context.Context, req *request.GetCategoryRequest) error {
	return f.DeleteCategory(ctx, req)
}

func (f *CategoryFacade) GetAllCategories(ctx context.Context) ([]*response.CategoryResponse, error) {
	return f.svc.GetAllCategories(ctx)
}
