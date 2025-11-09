package exporter

import (
	"context"
	"fmt"
	"strconv"
	"time"

	"github.com/niklvrr/Financial-Analytics-Service/internal/usecase/service"
)

type ByCategoryStrategy struct {
	entityType string
	service    interface{}
}

func NewByCategoryStrategy(entityType string) (*ByCategoryStrategy, error) {
	if entityType != "operation" {
		return nil, fmt.Errorf("стратегия by_category применима только для операций")
	}
	return &ByCategoryStrategy{
		entityType: entityType,
	}, nil
}

func (s *ByCategoryStrategy) SetService(svc interface{}) {
	s.service = svc
}

func (s *ByCategoryStrategy) Collect(ctx context.Context, params ExportParams) ([]map[string]string, error) {
	if s.service == nil {
		return nil, fmt.Errorf("сервис не установлен")
	}

	svc, ok := s.service.(*service.OperationService)
	if !ok {
		return nil, fmt.Errorf("неверный тип сервиса для операции")
	}

	operations, err := svc.GetAllOperations(ctx)
	if err != nil {
		return nil, err
	}

	result := make([]map[string]string, 0)
	for _, op := range operations {
		if op.CategoryId == params.CategoryID {
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

func (s *ByCategoryStrategy) GetHeaders() []string {
	return []string{"id", "kind", "bank_account_id", "amount", "date", "description", "category_id"}
}

