package category_commands

import (
	"bufio"
	"context"
	"fmt"
	"github.com/niklvrr/Financial-Analytics-Service/internal/usecase/facade"
	"github.com/niklvrr/Financial-Analytics-Service/pkg/utils"
	"os"
	"path/filepath"
)

type ImportCategoryCommand struct {
	f     *facade.CategoryFacade
	in    *bufio.Reader
	title string
}

func NewImportCategoryCommand(f *facade.CategoryFacade) *ImportCategoryCommand {
	return &ImportCategoryCommand{
		f:     f,
		in:    bufio.NewReader(os.Stdin),
		title: "Импорт категорий из файла",
	}
}

func (c *ImportCategoryCommand) Execute(ctx context.Context) error {
	prompt := "Введите полный путь до файла: "
	path, err := utils.AskString(c.in, prompt)
	if err != nil {
		return err
	}

	format := filepath.Ext(path)
	err = c.f.ImportCategoryFromFile(ctx, path, format)
	if err != nil {
		return err
	}
	fmt.Println("Данные успешно прочитаны из файла!")
	return nil
}

func (c *ImportCategoryCommand) Title() string {
	return c.title
}
