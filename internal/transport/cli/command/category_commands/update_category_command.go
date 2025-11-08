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

type UpdateCategoryCommand struct {
	f     *facade.CategoryFacade
	in    *bufio.Reader
	title string
}

func NewUpdateCategoryCommand(f *facade.CategoryFacade) *UpdateCategoryCommand {
	return &UpdateCategoryCommand{
		f:     f,
		in:    bufio.NewReader(os.Stdin),
		title: "Изменить категорию",
	}
}

func (c *UpdateCategoryCommand) Execute(ctx context.Context) error {
	idPrompt := "Введите уникальный номер категории: "
	id, err := utils.AskInt(c.in, idPrompt)
	if err != nil {
		return err
	}

	kindPrompt := "Введите новый тип категории(доход/расход): "
	kind, err := utils.AskString(c.in, kindPrompt)
	if err != nil {
		return err
	}

	namePrompt := "Введите новое название категории: "
	name, err := utils.AskString(c.in, namePrompt)
	if err != nil {
		return err
	}
	req := &request.UpdateCategoryRequest{
		Id:   int64(id),
		Kind: kind,
		Name: name,
	}

	err = c.f.UpdateCategory(ctx, req)
	if err != nil {
		return err
	}

	fmt.Println("Категория успещно изменена!")
	return nil
}

func (c *UpdateCategoryCommand) Title() string {
	return c.title
}
