package bank_account_commands

import (
	"bufio"
	"context"
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
	return c.f.ImportBankAccountsFromFile(ctx, path, format)
}

func (c *ImportBankAccountCommand) Title() string {
	return c.title
}
