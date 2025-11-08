package operation_commands

import (
	"bufio"
	"context"
	"fmt"
	"github.com/niklvrr/Financial-Analytics-Service/internal/usecase/facade"
	"github.com/niklvrr/Financial-Analytics-Service/pkg/utils"
	"os"
	"path/filepath"
)

type ImportOperationCommand struct {
	f     *facade.OperationFacade
	in    *bufio.Reader
	title string
}

func NewImportOperationCommand(f *facade.OperationFacade) *ImportOperationCommand {
	return &ImportOperationCommand{
		f:     f,
		in:    bufio.NewReader(os.Stdin),
		title: "Импорт операций из файла",
	}
}

func (c *ImportOperationCommand) Execute(ctx context.Context) error {
	prompt := "Введите полный путь до файла: "
	path, err := utils.AskString(c.in, prompt)
	if err != nil {
		return err
	}

	format := filepath.Ext(path)
	err = c.f.ImportOperationFromFile(ctx, path, format)
	if err != nil {
		return err
	}
	fmt.Println("Данные успешно прочитаны из файла!")
	return nil
}

func (c *ImportOperationCommand) Title() string {
	return c.title
}
