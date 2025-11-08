package operation_importer

import (
	"context"
	"encoding/json"
	"github.com/niklvrr/Financial-Analytics-Service/internal/domain/request"
	"github.com/niklvrr/Financial-Analytics-Service/internal/usecase/service"
	"github.com/niklvrr/Financial-Analytics-Service/pkg/utils"
	"os"
)

type JSONOperationImporter struct {
	svc *service.OperationService
}

func NewJSONOperationImporter(svc *service.OperationService) *JSONOperationImporter {
	return &JSONOperationImporter{
		svc: svc,
	}
}

func (j *JSONOperationImporter) Load(path string) ([]*request.CreateOperationRequest, error) {
	data, err := os.ReadFile(path)
	if err != nil {
		panic(err)
	}

	var reqs []*request.CreateOperationRequest
	err = json.Unmarshal(data, &reqs)
	if err != nil {
		return nil, err
	}
	return reqs, nil
}

func (c *JSONOperationImporter) Validate(reqs []*request.CreateOperationRequest) error {
	for _, req := range reqs {
		err := utils.ValidateKind(req.Kind)
		if err != nil {
			return err
		}

		err = utils.ValidateInt64(req.BankAccountId)
		if err != nil {
			return err
		}

		err = utils.ValidateFloat(req.Amount)
		if err != nil {
			return err
		}

		err = utils.ValidateInt64(req.CategoryId)
		if err != nil {
			return err
		}
	}
	return nil
}

func (c *JSONOperationImporter) Save(ctx context.Context, reqs []*request.CreateOperationRequest) error {
	for _, req := range reqs {
		err := c.svc.CreateOperation(ctx, req)
		if err != nil {
			return err
		}
	}
	return nil
}
