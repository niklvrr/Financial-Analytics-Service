package operation_commands

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

type ExportOperationCommand struct {
	f     *facade.OperationFacade
	in    *bufio.Reader
	title string
}

func NewExportOperationCommand(f *facade.OperationFacade) *ExportOperationCommand {
	return &ExportOperationCommand{
		f:     f,
		in:    bufio.NewReader(os.Stdin),
		title: "Экспорт операций в файл",
	}
}

func (c *ExportOperationCommand) Execute(ctx context.Context) error {
	fmt.Println("Доступные стратегии:")
	fmt.Println("  - full: экспорт всех операций")
	fmt.Println("  - by_account: экспорт операций по ID счета")
	fmt.Println("  - by_category: экспорт операций по ID категории")
	fmt.Println("  - date_range: экспорт операций за период")
	strategyPrompt := "Введите стратегию экспорта (full/by_account/by_category/date_range): "
	strategy, err := utils.AskString(c.in, strategyPrompt)
	if err != nil {
		return err
	}
	strategy = strings.ToLower(strings.TrimSpace(strategy))

	params := exporter.ExportParams{
		Format:   "",
		Strategy: strategy,
		Path:     "",
	}

	switch strategy {
	case "by_account":
		accountPrompt := "Введите ID счета: "
		accountID, err := utils.AskInt(c.in, accountPrompt)
		if err != nil {
			return err
		}
		params.AccountID = int64(accountID)

	case "by_category":
		categoryPrompt := "Введите ID категории: "
		categoryID, err := utils.AskInt(c.in, categoryPrompt)
		if err != nil {
			return err
		}
		params.CategoryID = int64(categoryID)

	case "date_range":
		dateFromPrompt := "Введите дату начала (YYYY-MM-DD) или оставьте пустым: "
		dateFrom, err := utils.AskString(c.in, dateFromPrompt)
		if err != nil {
			return err
		}
		params.DateFrom = strings.TrimSpace(dateFrom)

		dateToPrompt := "Введите дату конца (YYYY-MM-DD) или оставьте пустым: "
		dateTo, err := utils.AskString(c.in, dateToPrompt)
		if err != nil {
			return err
		}
		params.DateTo = strings.TrimSpace(dateTo)
	}

	pathPrompt := "Введите путь к файлу для сохранения: "
	path, err := utils.AskString(c.in, pathPrompt)
	if err != nil {
		return err
	}
	params.Path = strings.TrimSpace(path)

	format := filepath.Ext(path)
	params.Format = format

	report, err := c.f.Export(ctx, params)
	if err != nil {
		return err
	}

	if err := os.WriteFile(params.Path, report.Content, 0644); err != nil {
		return err
	}

	fmt.Printf("Данные успешно экспортированы в файл: %s\n", params.Path)
	return nil
}

func (c *ExportOperationCommand) Title() string {
	return c.title
}
