package operation_importer

import (
	"context"
	"github.com/niklvrr/Financial-Analytics-Service/internal/domain/request"
)

type OperationImporter interface {
	Load(path string) ([]*request.CreateOperationRequest, error)
	Validate(reqs []*request.CreateOperationRequest) error
	Save(ctx context.Context, reqs []*request.CreateOperationRequest) error
}

type OperationTemplate struct {
	Impl OperationImporter
}

func (t *OperationTemplate) Run(ctx context.Context, path string) error {
	data, err := t.Impl.Load(path)
	if err != nil {
		return err
	}

	err = t.Impl.Validate(data)
	if err != nil {
		return err
	}

	err = t.Impl.Save(ctx, data)
	if err != nil {
		return err
	}

	return nil
}
