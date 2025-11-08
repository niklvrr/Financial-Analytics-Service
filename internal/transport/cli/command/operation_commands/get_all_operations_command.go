package operation_commands

import (
	"bufio"
	"context"
	"fmt"
	"github.com/niklvrr/Financial-Analytics-Service/internal/usecase/facade"
	"os"
)

type GetAllOperationsCommand struct {
	f     *facade.OperationFacade
	in    *bufio.Reader
	title string
}

func NewGetAllOperationsCommand(f *facade.OperationFacade) *GetAllOperationsCommand {
	return &GetAllOperationsCommand{
		f:     f,
		in:    bufio.NewReader(os.Stdin),
		title: "Найти все опарции",
	}
}

func (c *GetAllOperationsCommand) Execute(ctx context.Context) error {
	ops, err := c.f.GetAllOperations(ctx)
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

func (c *GetAllOperationsCommand) Title() string {
	return c.title
}
