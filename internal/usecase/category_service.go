package usecase

import (
	"context"
	"github.com/niklvrr/Financial-Analytics-Service/internal/domain/model"
	"github.com/niklvrr/Financial-Analytics-Service/internal/domain/request"
	"github.com/niklvrr/Financial-Analytics-Service/internal/domain/response"
)

type CategoryRepo interface {
	CreateCategory(ctx context.Context, c *model.Category) error
	GetCategory(ctx context.Context, id int64) (*model.Category, error)
	UpdateCategory(ctx context.Context, c *model.Category) error
	DeleteCategory(ctx context.Context, id int64) error
	GetAllCategories(ctx context.Context) ([]*model.Category, error)
}

type CategoryService struct {
	repo CategoryRepo
}

func NewCategoryService(repo CategoryRepo) *CategoryService {
	return &CategoryService{repo: repo}
}

func (s *CategoryService) CreateCategory(ctx context.Context, req *request.CreateCategoryRequest) error {
	category := model.NewCategory(0, req.Kind, req.Name)
	err := s.repo.CreateCategory(ctx, category)
	if err != nil {
		return err
	}

	return nil
}

func (s *CategoryService) GetCategory(ctx context.Context, id int64) (*response.CategoryResponse, error) {
	category, err := s.repo.GetCategory(ctx, id)
	if err != nil {
		return nil, err
	}

	resp := &response.CategoryResponse{
		Id:   id,
		Kind: category.Kind(),
		Name: category.Name(),
	}

	return resp, nil
}

func (s *CategoryService) UpdateCategory(ctx context.Context, req *request.UpdateCategoryRequest) error {
	category := model.NewCategory(req.Id, req.Kind, req.Name)
	err := s.repo.UpdateCategory(ctx, category)
	if err != nil {
		return err
	}

	return nil
}

func (s *CategoryService) DeleteCategory(ctx context.Context, req *request.DeleteCategoryRequest) error {
	err := s.repo.DeleteCategory(ctx, req.Id)
	if err != nil {
		return err
	}

	return nil
}

func (s *CategoryService) GetAllCategories(ctx context.Context) ([]*response.CategoryResponse, error) {
	categories, err := s.repo.GetAllCategories(ctx)
	if err != nil {
		return nil, err
	}

	resp := make([]*response.CategoryResponse, len(categories))
	for i, category := range categories {
		resp[i] = &response.CategoryResponse{
			Id:   category.ID(),
			Kind: category.Kind(),
			Name: category.Name(),
		}
	}

	return resp, nil
}
