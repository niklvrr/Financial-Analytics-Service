package exporter

import (
	"context"
	"fmt"
	"strconv"
	"time"

	"github.com/niklvrr/Financial-Analytics-Service/internal/usecase/service"
)

type DateRangeStrategy struct {
	entityType string
	service    interface{}
}

func NewDateRangeStrategy(entityType string) (*DateRangeStrategy, error) {
	if entityType != "operation" {
		return nil, fmt.Errorf("стратегия date_range применима только для операций")
	}
	return &DateRangeStrategy{
		entityType: entityType,
	}, nil
}

func (s *DateRangeStrategy) SetService(svc interface{}) {
	s.service = svc
}

func (s *DateRangeStrategy) Collect(ctx context.Context, params ExportParams) ([]map[string]string, error) {
	if s.service == nil {
		return nil, fmt.Errorf("сервис не установлен")
	}

	svc, ok := s.service.(*service.OperationService)
	if !ok {
		return nil, fmt.Errorf("неверный тип сервиса для операции")
	}

	var dateFrom, dateTo time.Time
	var err error

	if params.DateFrom != "" {
		dateFrom, err = time.Parse("2006-01-02", params.DateFrom)
		if err != nil {
			return nil, fmt.Errorf("неверный формат даты начала: %w", err)
		}
	} else {
		dateFrom = time.Time{}
	}

	if params.DateTo != "" {
		dateTo, err = time.Parse("2006-01-02", params.DateTo)
		if err != nil {
			return nil, fmt.Errorf("неверный формат даты конца: %w", err)
		}
		dateTo = time.Date(dateTo.Year(), dateTo.Month(), dateTo.Day(), 23, 59, 59, 0, dateTo.Location())
	} else {
		dateTo = time.Now()
	}

	operations, err := svc.GetAllOperations(ctx)
	if err != nil {
		return nil, err
	}

	result := make([]map[string]string, 0)
	for _, op := range operations {
		opDate := op.Date
		afterFrom := dateFrom.IsZero() || opDate.After(dateFrom) || opDate.Equal(dateFrom)
		beforeTo := opDate.Before(dateTo) || opDate.Equal(dateTo)

		if afterFrom && beforeTo {
			result = append(result, map[string]string{
				"id":            strconv.FormatInt(op.Id, 10),
				"kind":          op.Kind,
				"bank_account_id": strconv.FormatInt(op.BankAccountId, 10),
				"amount":        strconv.FormatFloat(op.Amount, 'f', 2, 64),
				"date":          op.Date.Format(time.RFC3339),
				"description":   op.Description,
				"category_id":   strconv.FormatInt(op.CategoryId, 10),
			})
		}
	}

	return result, nil
}

func (s *DateRangeStrategy) GetHeaders() []string {
	return []string{"id", "kind", "bank_account_id", "amount", "date", "description", "category_id"}
}

