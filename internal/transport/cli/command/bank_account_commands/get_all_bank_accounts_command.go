package bank_account_commands

import (
	"bufio"
	"context"
	"fmt"
	"github.com/niklvrr/Financial-Analytics-Service/internal/usecase/facade"
	"os"
)

type GetAllBankAccountsCommand struct {
	f  *facade.BankAccountFacade
	in *bufio.Reader
}

func NewGetAllBankAccountsCommand(f *facade.BankAccountFacade) *GetAllBankAccountsCommand {
	return &GetAllBankAccountsCommand{
		f:  f,
		in: bufio.NewReader(os.Stdin),
	}
}

func (c *GetAllBankAccountsCommand) Execute(ctx context.Context) error {
	accounts, err := c.f.GetAllBankAccounts(ctx)
	if err != nil {
		return err
	}

	if len(accounts) == 0 {
		fmt.Println("Сохраненных операций не найдено")
		return nil
	}

	fmt.Print("=== Данные счетов ===\n")
	for _, account := range accounts {
		fmt.Printf("Номер: %d\n"+"Название: %s\n"+"Баланс: %g\n",
			account.Id, account.Name, account.Balance)
	}
	return nil
}
