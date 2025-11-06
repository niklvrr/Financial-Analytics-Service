package category_commands

import (
	"bufio"
	"context"
	"fmt"
	"github.com/niklvrr/Financial-Analytics-Service/internal/usecase/facade"
	"os"
)

type GetAllCategoriesCommand struct {
	f *facade.CategoryFacade
	in *bufio.Reader
}

func NewGetAllCategoriesCommand(f *facade.CategoryFacade) *GetAllCategoriesCommand {
	return &GetAllCategoriesCommand{
		f: f,
		in: bufio.NewReader(os.Stdin),
	}
}

func (c *GetAllCategoriesCommand) Execute(ctx context.Context) error {
	categories, err := c.f.GetAllCategories(ctx)
	if err != nil {
		return err
	}

	if len(categories) == 0 {
		fmt.Println("Сохраненных категорий не найдено")
		return nil
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