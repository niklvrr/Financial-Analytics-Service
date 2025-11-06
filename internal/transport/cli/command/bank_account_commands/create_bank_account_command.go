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

type CreateBankAccountCommand struct {
	f  *facade.BankAccountFacade
	in *bufio.Reader
}

func NewCreateBankAccountCommand(f *facade.BankAccountFacade) *CreateBankAccountCommand {
	return &CreateBankAccountCommand{
		f:  f,
		in: bufio.NewReader(os.Stdin),
	}
}

func (c *CreateBankAccountCommand) Execute(ctx context.Context) error {
	prompt := "Введите имя счета: "
	name, err := utils.AskString(c.in, prompt)
	if err != nil {
		return err
	}

	req := &request.CreateBankAccountRequest{
		Name: name,
	}

	err = c.f.CreateBankAccount(ctx, req)
	if err != nil {
		return err
	}
	fmt.Println("Счет успешно создан!")
	return nil
}
