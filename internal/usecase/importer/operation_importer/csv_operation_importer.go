package operation_importer

import (
	"context"
	"github.com/gocarina/gocsv"
	"github.com/niklvrr/Financial-Analytics-Service/internal/domain/request"
	"github.com/niklvrr/Financial-Analytics-Service/internal/usecase/service"
	"github.com/niklvrr/Financial-Analytics-Service/pkg/utils"
	"os"
)

type CSVOperationImporter struct {
	svc *service.OperationService
}

func NewCSVOperationImporter(svc *service.OperationService) *CSVOperationImporter {
	return &CSVOperationImporter{svc: svc}
}

func (c *CSVOperationImporter) Load(path string) ([]*request.CreateOperationRequest, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var reqs []*request.CreateOperationRequest
	if err := gocsv.UnmarshalFile(file, &reqs); err != nil {
		return nil, err
	}
	return reqs, nil
}

func (c *CSVOperationImporter) Validate(reqs []*request.CreateOperationRequest) error {
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

func (c *CSVOperationImporter) Save(ctx context.Context, reqs []*request.CreateOperationRequest) error {
	for _, req := range reqs {
		err := c.svc.CreateOperation(ctx, req)
		if err != nil {
			return err
		}
	}
	return nil
}
