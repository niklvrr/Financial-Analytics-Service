package operation_importer

import (
	"context"
	"github.com/niklvrr/Financial-Analytics-Service/internal/domain/request"
	"github.com/niklvrr/Financial-Analytics-Service/internal/usecase/service"
	"github.com/niklvrr/Financial-Analytics-Service/pkg/utils"
	"gopkg.in/yaml.v3"
	"os"
)

type YamlOperationImporter struct {
	svc *service.OperationService
}

func NewYamlOperationImporter(svc *service.OperationService) *YamlOperationImporter {
	return &YamlOperationImporter{
		svc: svc,
	}
}

func (c *YamlOperationImporter) Load(path string) ([]*request.CreateOperationRequest, error) {
	data, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}

	var reqs []*request.CreateOperationRequest
	err = yaml.Unmarshal(data, &reqs)
	if err != nil {
		return nil, err
	}
	return reqs, nil
}

func (c *YamlOperationImporter) Validate(reqs []*request.CreateOperationRequest) error {
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

func (c *YamlOperationImporter) Save(ctx context.Context, reqs []*request.CreateOperationRequest) error {
	for _, req := range reqs {
		err := c.svc.CreateOperation(ctx, req)
		if err != nil {
			return err
		}
	}
	return nil
}
