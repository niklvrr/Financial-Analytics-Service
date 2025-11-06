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

type GetOperationCommand struct {
	f *facade.OperationFacade
	in *bufio.Reader
}

func NewGetOperationCommand(f *facade.OperationFacade) *GetOperationCommand {
	return &GetOperationCommand{
		f: f,
		in: bufio.NewReader(os.Stdin),
	}
}

func (c *GetOperationCommand) Execute(ctx context.Context) error {
	idPrompt := "Введите уникальный номер операции"
	id, err := utils.AskInt(c.in, idPrompt)
	if err != nil {
		return err
	}

	req := &request.GetOperationRequest{
		Id: int64(id),
	}

	operation, err := c.f.GetOperation(ctx, req)
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
