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

type DeleteBankAccountCommand struct {
	f  *facade.BankAccountFacade
	in *bufio.Reader
}

func NewDeleteBankAccountCommand(f *facade.BankAccountFacade) *DeleteBankAccountCommand {
	return &DeleteBankAccountCommand{
		f:  f,
		in: bufio.NewReader(os.Stdin),
	}
}

func (c *DeleteBankAccountCommand) Execute(ctx context.Context) error {
	idPrompt := "Введите уникальный номер счета: "
	id, err := utils.AskInt(c.in, idPrompt)
	if err != nil {
		return err
	}

	req := &request.DeleteBankAccountRequest{
		Id: int64(id),
	}

	err = c.f.DeleteBankAccount(ctx, req)
	if err != nil {
		return err
	}
	fmt.Println("Счет успешно удален!")

	return nil
}
