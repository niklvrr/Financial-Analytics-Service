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

type UpdateOperationCommand struct {
	f *facade.OperationFacade
	in *bufio.Reader
}

func NewUpdateOperationCommand(f *facade.OperationFacade) *UpdateOperationCommand {
	return &UpdateOperationCommand{
		f: f,
		in: bufio.NewReader(os.Stdin),
	}
}

func (c *UpdateOperationCommand) Execute(ctx context.Context) error {
	idPrompt := "Введите уникальный номер опреации: "
	id, err := utils.AskInt(c.in, idPrompt)
	if err != nil {
		return err
	}

	kindPrompt := "Введите тип операции(доход/расход): "
	kind, err := utils.AskString(c.in, kindPrompt)
	if err != nil {
		return err
	}

	bankAccountIdPrompt := "Введите новый уникальный номер счета: "
	bankAccountId, err := utils.AskInt(c.in, bankAccountIdPrompt)
	if err != nil {
		return err
	}

	amountPrompt := "Введите новую сумму операции: "
	amount, err := utils.AskFloat(c.in, amountPrompt)
	if err != nil {
		return err
	}

	descPrompt := "Введите новое описание операции: "
	desc, err := utils.AskString(c.in, descPrompt)
	if err != nil {
		return err
	}

	categoryIdPrompt := "Введите новый уникальный номер категории: "
	categoryId, err := utils.AskInt(c.in, categoryIdPrompt)
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

	err = c.f.UpdateOperation(ctx, req)
	if err != nil {
		return err
	}
	fmt.Println("Операция успешно изменена!")
	return nil
}