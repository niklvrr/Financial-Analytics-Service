package category_commands

import (
	"bufio"
	"context"
	"github.com/niklvrr/Financial-Analytics-Service/internal/usecase/facade"
	"github.com/niklvrr/Financial-Analytics-Service/pkg/utils"
	"os"
	"path/filepath"
)

type ImportCategoryCommand struct {
	f  *facade.CategoryFacade
	in *bufio.Reader
}

func NewImportCategoryCommand(f *facade.CategoryFacade) *ImportCategoryCommand {
	return &ImportCategoryCommand{
		f:  f,
		in: bufio.NewReader(os.Stdin),
	}
}

func (c *ImportCategoryCommand) Execute(ctx context.Context) error {
	prompt := "Введите полный путь до файла: "
	path, err := utils.AskString(c.in, prompt)
	if err != nil {
		return err
	}

	format := filepath.Ext(path)
	return c.f.ImportCategoryFromFile(ctx, path, format)
}
