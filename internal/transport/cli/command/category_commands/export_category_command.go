package category_commands

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

type ExportCategoryCommand struct {
	f     *facade.CategoryFacade
	in    *bufio.Reader
	title string
}

func NewExportCategoryCommand(f *facade.CategoryFacade) *ExportCategoryCommand {
	return &ExportCategoryCommand{
		f:     f,
		in:    bufio.NewReader(os.Stdin),
		title: "Экспорт категорий в файл",
	}
}

func (c *ExportCategoryCommand) Execute(ctx context.Context) error {
	pathPrompt := "Введите путь к файлу для сохранения: "
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
		return fmt.Errorf("ошибка экспорта: %w", err)
	}

	if err := os.WriteFile(path, report.Content, 0644); err != nil {
		return fmt.Errorf("ошибка записи в файл: %w", err)
	}

	fmt.Printf("Данные успешно экспортированы в файл: %s\n", path)
	return nil
}

func (c *ExportCategoryCommand) Title() string {
	return c.title
}
