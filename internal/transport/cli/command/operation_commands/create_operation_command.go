package operation_commands

import (
	"bufio"
	"context"
	"fmt"
	"github.com/niklvrr/Financial-Analytics-Service/internal/domain/request"
	"github.com/niklvrr/Financial-Analytics-Service/internal/usecase/facade"
	"github.com/niklvrr/Financial-Analytics-Service/pkg/utils"
	"os"
)

type CreateOperationCommand struct {
	f     *facade.OperationFacade
	in    *bufio.Reader
	title string
}

func NewCreateOperationCommand(f *facade.OperationFacade) *CreateOperationCommand {
	return &CreateOperationCommand{
		f:     f,
		in:    bufio.NewReader(os.Stdin),
		title: "Создать операцию",
	}
}

func (c *CreateOperationCommand) Execute(ctx context.Context) error {
	kindPrompt := "Введите тип операции(доход/расход): "
	kind, err := utils.AskString(c.in, kindPrompt)
	if err != nil {
		return err
	}

	bankAccountIdPrompt := "Введите уникальный номер счета: "
	bankAccountId, err := utils.AskInt(c.in, bankAccountIdPrompt)
	if err != nil {
		return err
	}

	amountPrompt := "Введите сумму операции: "
	amount, err := utils.AskFloat(c.in, amountPrompt)
	if err != nil {
		return err
	}

	descPrompt := "Введите описание операции: "
	desc, err := utils.AskString(c.in, descPrompt)
	if err != nil {
		return err
	}

	categoryIdPrompt := "Введите уникальный номер категории: "
	categoryId, err := utils.AskInt(c.in, categoryIdPrompt)
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

	err = c.f.CreateOperation(ctx, req)
	if err != nil {
		return err
	}

	fmt.Println("Операция успешно создана!")
	return nil
}

func (c *CreateOperationCommand) Title() string {
	return c.title
}
