package category_commands

import (
	"bufio"
	"context"
	"fmt"
	"github.com/niklvrr/Financial-Analytics-Service/internal/domain/request"
	"github.com/niklvrr/Financial-Analytics-Service/internal/usecase/facade"
	"github.com/niklvrr/Financial-Analytics-Service/pkg/utils"
	"os"
)

type GetCategoryCommand struct {
	f     *facade.CategoryFacade
	in    *bufio.Reader
	title string
}

func NewGetCategoryCommand(f *facade.CategoryFacade) *GetCategoryCommand {
	return &GetCategoryCommand{
		f:     f,
		in:    bufio.NewReader(os.Stdin),
		title: "Найти категорию",
	}
}

func (c *GetCategoryCommand) Execute(ctx context.Context) error {
	idPrompt := "Введите уникальный номер категории: "
	id, err := utils.AskInt(c.in, idPrompt)
	if err != nil {
		return err
	}

	req := &request.GetCategoryRequest{
		Id: int64(id),
	}

	category, err := c.f.GetCategory(ctx, req)
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

func (c *GetCategoryCommand) Title() string {
	return c.title
}
