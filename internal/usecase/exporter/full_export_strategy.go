package exporter

import (
	"context"
	"fmt"
	"strconv"
	"time"

	"github.com/niklvrr/Financial-Analytics-Service/internal/usecase/service"
)

type FullExportStrategy struct {
	entityType string
	service    interface{}
}

func NewFullExportStrategy(entityType string) (*FullExportStrategy, error) {
	return &FullExportStrategy{
		entityType: entityType,
	}, nil
}

func (s *FullExportStrategy) SetService(service interface{}) {
	s.service = service
}

func (s *FullExportStrategy) Collect(ctx context.Context, params ExportParams) ([]map[string]string, error) {
	switch s.entityType {
	case "bank_account":
		return s.collectBankAccounts(ctx)
	case "category":
		return s.collectCategories(ctx)
	case "operation":
		return s.collectOperations(ctx)
	default:
		return nil, fmt.Errorf("неподдерживаемый тип сущности: %s", s.entityType)
	}
}

func (s *FullExportStrategy) GetHeaders() []string {
	switch s.entityType {
	case "bank_account":
		return []string{"id", "name", "balance"}
	case "category":
		return []string{"id", "kind", "name"}
	case "operation":
		return []string{"id", "kind", "bank_account_id", "amount", "date", "description", "category_id"}
	default:
		return []string{}
	}
}

func (s *FullExportStrategy) collectBankAccounts(ctx context.Context) ([]map[string]string, error) {
	svc, ok := s.service.(*service.BankAccountService)
	if !ok {
		return nil, fmt.Errorf("неверный тип сервиса для bank_account")
	}

	accounts, err := svc.GetAllBankAccounts(ctx)
	if err != nil {
		return nil, err
	}

	result := make([]map[string]string, 0, len(accounts))
	for _, acc := range accounts {
		result = append(result, map[string]string{
			"id":      strconv.FormatInt(acc.Id, 10),
			"name":    acc.Name,
			"balance": strconv.FormatFloat(acc.Balance, 'f', 2, 64),
		})
	}

	return result, nil
}

func (s *FullExportStrategy) collectCategories(ctx context.Context) ([]map[string]string, error) {
	svc, ok := s.service.(*service.CategoryService)
	if !ok {
		return nil, fmt.Errorf("неверный тип сервиса для category")
	}

	categories, err := svc.GetAllCategories(ctx)
	if err != nil {
		return nil, err
	}

	result := make([]map[string]string, 0, len(categories))
	for _, cat := range categories {
		result = append(result, map[string]string{
			"id":   strconv.FormatInt(cat.Id, 10),
			"kind": cat.Kind,
			"name": cat.Name,
		})
	}

	return result, nil
}

func (s *FullExportStrategy) collectOperations(ctx context.Context) ([]map[string]string, error) {
	svc, ok := s.service.(*service.OperationService)
	if !ok {
		return nil, fmt.Errorf("неверный тип сервиса для operation")
	}

	operations, err := svc.GetAllOperations(ctx)
	if err != nil {
		return nil, err
	}

	result := make([]map[string]string, 0, len(operations))
	for _, op := range operations {
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

	return result, nil
}

