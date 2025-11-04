package handlers

import (
	"bufio"
	"context"
	"fmt"
	"github.com/niklvrr/Financial-Analytics-Service/internal/domain/request"
	"github.com/niklvrr/Financial-Analytics-Service/internal/domain/response"
	"github.com/niklvrr/Financial-Analytics-Service/pkg/utils"
	"os"
)

type CategoryService interface {
	CreateCategory(ctx context.Context, req *request.CreateCategoryRequest) error
	GetCategory(ctx context.Context, req *request.GetCategoryRequest) (*response.CategoryResponse, error)
	UpdateCategory(ctx context.Context, req *request.UpdateCategoryRequest) error
	DeleteCategory(ctx context.Context, req *request.DeleteCategoryRequest) error
	GetAllCategories(ctx context.Context) ([]*response.CategoryResponse, error)
}

type CategoryHandler struct {
	svc CategoryService
	in  *bufio.Reader
}

func NewCategoryHandler(svc CategoryService) *CategoryHandler {
	return &CategoryHandler{
		svc: svc,
		in:  bufio.NewReader(os.Stdin),
	}
}

func (h *CategoryHandler) CreateCategory(ctx context.Context) error {
	kindPrompt := "Введите тип категори(доход/расход): "
	kind, err := utils.AskString(h.in, kindPrompt)
	if err != nil {
		return err
	}

	namePrompt := "Введите название категории: "
	name, err := utils.AskString(h.in, namePrompt)
	if err != nil {
		return err
	}

	req := &request.CreateCategoryRequest{
		Kind: kind,
		Name: name,
	}

	err = h.svc.CreateCategory(ctx, req)
	if err != nil {
		return err
	}

	fmt.Println("Категория успешно создана!")
	return nil
}

func (h *CategoryHandler) GetCategory(ctx context.Context) error {
	idPrompt := "Введите уникальный номер категории: "
	id, err := utils.AskInt(h.in, idPrompt)
	if err != nil {
		return err
	}

	req := &request.GetCategoryRequest{
		Id: int64(id),
	}

	category, err := h.svc.GetCategory(ctx, req)
	if err != nil {
		return err
	}

	fmt.Printf("=== Данные категории ===\n"+
		"Номер: %d\n"+
		"Тип: %s\n"+
		"Название: %s\n",
		category.Id,
		category.Kind,
		category.Name)

	return nil
}

func (h *CategoryHandler) UpdateCategory(ctx context.Context) error {
	idPrompt := "Введите уникальный номер категории: "
	id, err := utils.AskInt(h.in, idPrompt)
	if err != nil {
		return err
	}

	kindPrompt := "Введите новый тип категории(доход/расход): "
	kind, err := utils.AskString(h.in, kindPrompt)
	if err != nil {
		return err
	}

	name, err := utils.AskString(h.in, kindPrompt)
	if err != nil {
		return err
	}
	req := &request.UpdateCategoryRequest{
		Id:   int64(id),
		Kind: kind,
		Name: name,
	}

	err = h.svc.UpdateCategory(ctx, req)
	if err != nil {
		return err
	}

	fmt.Println("Категория успещно изменена!")
	return nil
}

func (h *CategoryHandler) DeleteCategory(ctx context.Context) error {
	idPrompt := "Введите уникальный номер категории: "
	id, err := utils.AskInt(h.in, idPrompt)
	if err != nil {
		return err
	}
	req := &request.DeleteCategoryRequest{
		Id: int64(id),
	}

	err = h.svc.DeleteCategory(ctx, req)
	if err != nil {
		return err
	}

	fmt.Println("Категория успешно удалена!")
	return nil
}

func (h *CategoryHandler) GetAllCategories(ctx context.Context) error {
	categories, err := h.svc.GetAllCategories(ctx)
	if err != nil {
		return err
	}

	fmt.Print("=== Данные категорий ===\n")
	for _, category := range categories {
		fmt.Printf("Номер: %d\n"+
			"Тип: %s\n"+
			"Название: %s\n",
			category.Id,
			category.Kind,
			category.Name)
	}
	return nil
}
