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

type GetBankAccountsCommand struct {
	f     *facade.BankAccountFacade
	in    *bufio.Reader
	title string
}

func NewGetBankAccountCommand(facade *facade.BankAccountFacade) *GetBankAccountsCommand {
	return &GetBankAccountsCommand{
		f:     facade,
		in:    bufio.NewReader(os.Stdin),
		title: "Найти банковкий счет",
	}
}

func (c *GetBankAccountsCommand) Execute(ctx context.Context) error {
	prompt := "Введите уникальный номер счета: "
	id, err := utils.AskInt(c.in, prompt)
	if err != nil {
		return err
	}

	req := &request.GetBankAccountsRequest{
		Id: int64(id),
	}
	res, err := c.f.GetBankAccount(ctx, req)
	if err != nil {
		return err
	}

	fmt.Printf("=== Данные счета ===\n"+"Номер: %d\n"+"Название: %s\n"+"Баланс: %g\n",
		res.Id, res.Name, res.Balance)

	return nil
}

func (c *GetBankAccountsCommand) Title() string {
	return c.title
}
