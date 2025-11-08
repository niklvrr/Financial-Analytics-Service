package bank_account_commands

import (
	"bufio"
	"context"
	"fmt"
	"github.com/niklvrr/Financial-Analytics-Service/internal/usecase/facade"
	"github.com/niklvrr/Financial-Analytics-Service/pkg/utils"
	"os"
	"path/filepath"
)

type ImportBankAccountCommand struct {
	f     *facade.BankAccountFacade
	in    *bufio.Reader
	title string
}

func NewImportBankAccountCommand(f *facade.BankAccountFacade) *ImportBankAccountCommand {
	return &ImportBankAccountCommand{
		f:     f,
		in:    bufio.NewReader(os.Stdin),
		title: "Импорт банковских счетов из файла",
	}
}

func (c *ImportBankAccountCommand) Execute(ctx context.Context) error {
	prompt := "Введите полный путь до файла: "
	path, err := utils.AskString(c.in, prompt)
	if err != nil {
		return err
	}

	format := filepath.Ext(path)
	err = c.f.ImportBankAccountsFromFile(ctx, path, format)
	if err != nil {
		return err
	}

	fmt.Println("Данные успешно прочитаны из файла!")
	return nil
}

func (c *ImportBankAccountCommand) Title() string {
	return c.title
}
