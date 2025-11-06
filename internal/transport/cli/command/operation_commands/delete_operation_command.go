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

type DeleteOperationCommand struct {
	f *facade.OperationFacade
	in *bufio.Reader
}

func NewDeleteOperationCommand(f *facade.OperationFacade) *DeleteOperationCommand {
	return &DeleteOperationCommand{
		f: f,
		in: bufio.NewReader(os.Stdin),
	}
}

func (c *DeleteOperationCommand) Execute(ctx context.Context) error {
	idPrompt := "Введите уникальный номер опреации: "
	id, err := utils.AskInt(c.in, idPrompt)
	if err != nil {
		return err
	}

	req := &request.DeleteOperationRequest{
		Id: int64(id),
	}

	err = c.f.DeleteOperation(ctx, req)
	if err != nil {
		return err
	}

	fmt.Println("Операция успешно удалена!")
	return nil
}
