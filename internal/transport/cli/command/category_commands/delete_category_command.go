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

type DeleteCategoryCommand struct {
	f     *facade.CategoryFacade
	in    *bufio.Reader
	title string
}

func NewDeleteCategoryCommand(f *facade.CategoryFacade) *DeleteCategoryCommand {
	return &DeleteCategoryCommand{
		f:     f,
		in:    bufio.NewReader(os.Stdin),
		title: "Удалить категорию",
	}
}

func (c *DeleteCategoryCommand) Execute(ctx context.Context) error {
	idPrompt := "Введите уникальный номер категории: "
	id, err := utils.AskInt(c.in, idPrompt)
	if err != nil {
		return err
	}
	req := &request.DeleteCategoryRequest{
		Id: int64(id),
	}

	err = c.f.DeleteCategory(ctx, req)
	if err != nil {
		return err
	}

	fmt.Println("Категория успешно удалена!")
	return nil
}

func (c *DeleteCategoryCommand) Title() string {
	return c.title
}
