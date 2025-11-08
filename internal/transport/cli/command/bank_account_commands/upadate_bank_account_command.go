package bank_account_commands

import (
	"bufio"
	"context"
	"fmt"
	"github.com/niklvrr/Financial-Analytics-Service/internal/domain/request"
	"github.com/niklvrr/Financial-Analytics-Service/internal/usecase/facade"
	"github.com/niklvrr/Financial-Analytics-Service/pkg/utils"
	"os"
)

type UpdateBankAccountCommand struct {
	f     *facade.BankAccountFacade
	in    *bufio.Reader
	title string
}

func NewUpdateBankAccountCommand(f *facade.BankAccountFacade) *UpdateBankAccountCommand {
	return &UpdateBankAccountCommand{
		f:     f,
		in:    bufio.NewReader(os.Stdin),
		title: "Изменить банковский счет",
	}
}

func (c *UpdateBankAccountCommand) Execute(ctx context.Context) error {
	idPrompt := "Введите уникальный номер счета: "
	id, err := utils.AskInt(c.in, idPrompt)
	if err != nil {
		return err
	}

	namePrompt := "Введите новое имя счета: "
	name, err := utils.AskString(c.in, namePrompt)
	if err != nil {
		return err
	}

	balancePrompt := "Введите новое значение баланса счета: "
	balance, err := utils.AskFloat(c.in, balancePrompt)
	if err != nil {
		return err
	}

	req := &request.UpdateBankAccountRequest{
		Id:      int64(id),
		Name:    name,
		Balance: balance,
	}

	err = c.f.UpdateBankAccount(ctx, req)
	if err != nil {
		return err
	}
	fmt.Println("Данные счета успешно изменены!")
	return nil
}

func (c *UpdateBankAccountCommand) Title() string {
	return c.title
}
