package bank_account_commands

import (
	"bufio"
	"context"
	"fmt"
	"github.com/niklvrr/Financial-Analytics-Service/internal/usecase/exporter"
	"github.com/niklvrr/Financial-Analytics-Service/internal/usecase/facade"
	"github.com/niklvrr/Financial-Analytics-Service/pkg/utils"
	"os"
	"path/filepath"
	"strings"
)

type ExportBankAccountCommand struct {
	f     *facade.BankAccountFacade
	in    *bufio.Reader
	title string
}

func NewExportBankAccountCommand(f *facade.BankAccountFacade) *ExportBankAccountCommand {
	return &ExportBankAccountCommand{
		f:     f,
		in:    bufio.NewReader(os.Stdin),
		title: "Экспорт банковских счетов в файл",
	}
}

func (c *ExportBankAccountCommand) Execute(ctx context.Context) error {
	pathPrompt := "Введите полный путь к файлу для сохранения: "
	path, err := utils.AskString(c.in, pathPrompt)
	if err != nil {
		return err
	}
	path = strings.TrimSpace(path)

	format := filepath.Ext(path)

	params := exporter.ExportParams{
		Format:   format,
		Strategy: "full",
		Path:     path,
	}

	report, err := c.f.Export(ctx, params)
	if err != nil {
		return err
	}

	if err := os.WriteFile(path, report.Content, 0644); err != nil {
		return err
	}

	fmt.Printf("Данные успешно экспортированы в файл: %s\n", path)
	return nil
}

func (c *ExportBankAccountCommand) Title() string {
	return c.title
}
