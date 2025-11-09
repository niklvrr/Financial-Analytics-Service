package exporter

import (
	"context"
)

type ExportParams struct {
	Format     string
	Strategy   string
	Path       string
	AccountID  int64
	CategoryID int64
	DateFrom   string
	DateTo     string
}

type ExportStrategy interface {
	Collect(ctx context.Context, params ExportParams) ([]map[string]string, error)
	GetHeaders() []string
}

func NewStrategy(name string, entityType string) (ExportStrategy, error) {
	switch name {
	case "full":
		return NewFullExportStrategy(entityType)
	case "by_account":
		return NewByAccountStrategy(entityType)
	case "by_category":
		return NewByCategoryStrategy(entityType)
	case "date_range":
		return NewDateRangeStrategy(entityType)
	default:
		return nil, ErrInvalidStrategy
	}
}

