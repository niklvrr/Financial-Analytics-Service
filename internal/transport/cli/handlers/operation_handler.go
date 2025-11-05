package handlers

import (
	"bufio"
	"context"
	"fmt"
	"github.com/niklvrr/Financial-Analytics-Service/internal/domain/request"
	"github.com/niklvrr/Financial-Analytics-Service/internal/domain/response"
	"github.com/niklvrr/Financial-Analytics-Service/pkg/utils"
	"os"
)

type OperationService interface {
	CreateOperation(ctx context.Context, req *request.CreateOperationRequest) error
	GetOperation(ctx context.Context, req *request.GetOperationRequest) (*response.OperationResponse, error)
	UpdateOperation(ctx context.Context, req *request.UpdateOperationRequest) error
	DeleteOperation(ctx context.Context, req *request.DeleteOperationRequest) error
	GetAllOperations(ctx context.Context) ([]*response.OperationResponse, error)
}

type OperationHandler struct {
	svc OperationService
	in  *bufio.Reader
}

func NewOperationHandler(svc OperationService) *OperationHandler {
	return &OperationHandler{
		svc: svc,
		in:  bufio.NewReader(os.Stdin),
	}
}

func (h *OperationHandler) CreateOperation(ctx context.Context) error {
	kindPrompt := "Введите тип операции(доход/расход): "
	kind, err := utils.AskString(h.in, kindPrompt)
	if err != nil {
		return err
	}

	bankAccountIdPrompt := "Введите уникальный номер счета: "
	bankAccountId, err := utils.AskInt(h.in, bankAccountIdPrompt)
	if err != nil {
		return err
	}

	amountPrompt := "Введите сумму операции: "
	amount, err := utils.AskFloat(h.in, amountPrompt)
	if err != nil {
		return err
	}

	descPrompt := "Введите описание операции"
	desc, err := utils.AskString(h.in, descPrompt)
	if err != nil {
		return err
	}

	categoryIdPrompt := "Введите уникальный номер категории: "
	categoryId, err := utils.AskInt(h.in, categoryIdPrompt)
	if err != nil {
		return err
	}

	req := &request.CreateOperationRequest{
		Kind:          kind,
		BankAccountId: int64(bankAccountId),
		Amount:        amount,
		Description:   desc,
		CategoryId:    int64(categoryId),
	}

	err = h.svc.CreateOperation(ctx, req)
	if err != nil {
		return err
	}

	fmt.Println("Операция успешно создана!")
	return nil
}

func (h *OperationHandler) GetOperation(ctx context.Context) error {
	idPrompt := "Введите уникальный номер операции"
	id, err := utils.AskInt(h.in, idPrompt)
	if err != nil {
		return err
	}

	req := &request.GetOperationRequest{
		Id: int64(id),
	}

	operation, err := h.svc.GetOperation(ctx, req)
	if err != nil {
		return err
	}

	fmt.Print("=== Данные операции ===\n")
	fmt.Printf("Номер операции: %d\n"+
		"Тип: %s\n"+
		"Номер счета: %d\n"+
		"Сумма: %g\n"+
		"Дата: %s\n"+
		"Описание: %s\n"+
		"Номер: %d\n",
		operation.Id,
		operation.Kind,
		operation.BankAccountId,
		operation.Amount,
		operation.Date,
		operation.Description,
		operation.CategoryId)
	return nil
}

func (h *OperationHandler) UpdateOperation(ctx context.Context) error {
	idPrompt := "Введите уникальный номер опреации: "
	id, err := utils.AskInt(h.in, idPrompt)
	if err != nil {
		return err
	}

	kindPrompt := "Введите тип операции(доход/расход): "
	kind, err := utils.AskString(h.in, kindPrompt)
	if err != nil {
		return err
	}

	bankAccountIdPrompt := "Введите новый уникальный номер счета: "
	bankAccountId, err := utils.AskInt(h.in, bankAccountIdPrompt)
	if err != nil {
		return err
	}

	amountPrompt := "Введите новую сумму операции: "
	amount, err := utils.AskFloat(h.in, amountPrompt)
	if err != nil {
		return err
	}

	descPrompt := "Введите новое описание операции: "
	desc, err := utils.AskString(h.in, descPrompt)
	if err != nil {
		return err
	}

	categoryIdPrompt := "Введите новый уникальный номер категории: "
	categoryId, err := utils.AskInt(h.in, categoryIdPrompt)
	if err != nil {
		return err
	}

	req := &request.UpdateOperationRequest{
		Id:            int64(id),
		Kind:          kind,
		BankAccountId: int64(bankAccountId),
		Amount:        amount,
		Description:   desc,
		CategoryId:    int64(categoryId),
	}

	err = h.svc.UpdateOperation(ctx, req)
	if err != nil {
		return err
	}
	fmt.Println("Операция успешно изменена!")
	return nil
}

func (h *OperationHandler) DeleteOperation(ctx context.Context) error {
	idPrompt := "Введите уникальный номер опреации: "
	id, err := utils.AskInt(h.in, idPrompt)
	if err != nil {
		return err
	}

	req := &request.DeleteOperationRequest{
		Id: int64(id),
	}

	err = h.svc.DeleteOperation(ctx, req)
	if err != nil {
		return err
	}

	fmt.Println("Операция успешно удалена!")
	return nil
}

func (h *OperationHandler) GetAllOperations(ctx context.Context) error {
	ops, err := h.svc.GetAllOperations(ctx)
	if err != nil {
		return err
	}

	if len(ops) == 0 {
		fmt.Println("Сохраненных операций не найдено")
		return nil
	}

	fmt.Print("=== Данные операций ===\n")
	for _, operation := range ops {
		fmt.Printf("Номер операции: %d\n"+
			"Тип: %s\n"+
			"Номер счета: %d\n"+
			"Сумма: %g\n"+
			"Дата: %s\n"+
			"Описание: %s\n"+
			"Номер: %d\n",
			operation.Id,
			operation.Kind,
			operation.BankAccountId,
			operation.Amount,
			operation.Date,
			operation.Description,
			operation.CategoryId)
	}
	return nil
}
