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

type CreateCategoryCommand struct {
	f     *facade.CategoryFacade
	in    *bufio.Reader
	title string
}

func NewCreateCategoryCommand(f *facade.CategoryFacade) *CreateCategoryCommand {
	return &CreateCategoryCommand{
		f:     f,
		in:    bufio.NewReader(os.Stdin),
		title: "Создать категорию",
	}
}

func (c *CreateCategoryCommand) Execute(ctx context.Context) error {
	kindPrompt := "Введите тип категори(доход/расход): "
	kind, err := utils.AskString(c.in, kindPrompt)
	if err != nil {
		return err
	}

	namePrompt := "Введите название категории: "
	name, err := utils.AskString(c.in, namePrompt)
	if err != nil {
		return err
	}

	req := &request.CreateCategoryRequest{
		Kind: kind,
		Name: name,
	}

	err = c.f.CreateCategory(ctx, req)
	if err != nil {
		return err
	}

	fmt.Println("Категория успешно создана!")
	return nil
}

func (c *CreateCategoryCommand) Title() string {
	return c.title
}
