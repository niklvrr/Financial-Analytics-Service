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

type BankAccountService interface {
	CreateBankAccount(ctx context.Context, req *request.CreateBankAccountRequest) error
	GetBankAccount(ctx context.Context, req *request.GetBankAccountsRequest) (*response.BankAccountResponse, error)
	UpdateBankAccount(ctx context.Context, req *request.UpdateBankAccountRequest) error
	DeleteBankAccount(ctx context.Context, req *request.DeleteBankAccountRequest) error
	GetAllBankAccounts(ctx context.Context) ([]*response.BankAccountResponse, error)
}

type BankAccountHandler struct {
	svc BankAccountService
	in  *bufio.Reader
}

func NewBankAccountHandler(svc BankAccountService) *BankAccountHandler {
	return &BankAccountHandler{
		svc: svc,
		in:  bufio.NewReader(os.Stdin),
	}
}

func (h *BankAccountHandler) CreateBankAccount(ctx context.Context) error {
	prompt := "Введите имя счета: "
	name, err := utils.AskString(h.in, prompt)
	if err != nil {
		return err
	}

	req := &request.CreateBankAccountRequest{
		Name: name,
	}

	err = h.svc.CreateBankAccount(ctx, req)
	if err != nil {
		return err
	}
	fmt.Println("Счет успешно создан!")
	return nil
}

func (h *BankAccountHandler) GetBankAccount(ctx context.Context) error {
	prompt := "Введите уникальный номер счета: "
	id, err := utils.AskInt(h.in, prompt)
	if err != nil {
		return err
	}

	req := &request.GetBankAccountsRequest{
		Id: int64(id),
	}
	res, err := h.svc.GetBankAccount(ctx, req)
	if err != nil {
		return err
	}

	fmt.Printf("=== Данные счета ===\n"+"Номер: %d\n"+"Название: %s\n"+"Баланс: %d\n",
		res.Id, res.Name, res.Balance)

	return nil
}

func (h *BankAccountHandler) UpdateBankAccount(ctx context.Context) error {
	idPrompt := "Введите уникальный номер счета: "
	id, err := utils.AskInt(h.in, idPrompt)
	if err != nil {
		return err
	}

	namePrompt := "Введите новое имя счета: "
	name, err := utils.AskString(h.in, namePrompt)
	if err != nil {
		return err
	}

	balancePrompt := "Введите новое значение баланса счета: "
	balance, err := utils.AskFloat(h.in, balancePrompt)
	if err != nil {
		return err
	}

	req := &request.UpdateBankAccountRequest{
		Id:      int64(id),
		Name:    name,
		Balance: balance,
	}

	err = h.svc.UpdateBankAccount(ctx, req)
	if err != nil {
		return err
	}
	fmt.Println("Данные счета успешно изменены!")
	return nil
}

func (h *BankAccountHandler) DeleteBankAccount(ctx context.Context) error {
	idPrompt := "Введите уникальный номер счета: "
	id, err := utils.AskInt(h.in, idPrompt)
	if err != nil {
		return err
	}

	req := &request.DeleteBankAccountRequest{
		Id: int64(id),
	}

	err = h.svc.DeleteBankAccount(ctx, req)
	if err != nil {
		return err
	}
	fmt.Println("Счет успешно удален!")

	return nil
}

func (h *BankAccountHandler) GetAllBankAccounts(ctx context.Context) error {
	accounts, err := h.svc.GetAllBankAccounts(ctx)
	if err != nil {
		return err
	}

	fmt.Print("=== Данные счетов ===\n")
	for _, account := range accounts {
		fmt.Printf("Номер: %d\n"+"Название: %s\n"+"Баланс: %d\n",
			account.Id, account.Name, account.Balance)
	}
	return nil
}
