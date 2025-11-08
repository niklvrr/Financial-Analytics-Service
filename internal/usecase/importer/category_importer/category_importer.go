package category_importer

import (
	"context"
	"github.com/niklvrr/Financial-Analytics-Service/internal/domain/request"
)

type CategoryImporter interface {
	Load(path string) ([]*request.CreateCategoryRequest, error)
	Validate(req []*request.CreateCategoryRequest) error
	Save(ctx context.Context, req []*request.CreateCategoryRequest) error
}

type CategoryTemplate struct {
	Impl CategoryImporter
}

func (c *CategoryTemplate) Run(ctx context.Context, path string) error {
	data, err := c.Impl.Load(path)
	if err != nil {
		return err
	}

	err = c.Impl.Validate(data)
	if err != nil {
		return err
	}

	err = c.Impl.Save(data)
	if err != nil {
		return err
	}
	return nil
}
