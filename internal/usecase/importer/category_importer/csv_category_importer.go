package category_importer

import (
	"context"
	"github.com/gocarina/gocsv"
	"github.com/niklvrr/Financial-Analytics-Service/internal/domain/request"
	"github.com/niklvrr/Financial-Analytics-Service/internal/usecase/service"
	"github.com/niklvrr/Financial-Analytics-Service/pkg/utils"
	"os"
)

type CSVCategoryImporter struct {
	svc *service.CategoryService
}

func NewCSVCategoryImporter(svc *service.CategoryService) *CSVCategoryImporter {
	return &CSVCategoryImporter{
		svc: svc,
	}
}

func (imp *CSVCategoryImporter) Load(path string) ([]*request.CreateCategoryRequest, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var reqs []*request.CreateCategoryRequest
	if err := gocsv.UnmarshalFile(file, &reqs); err != nil {
		return nil, err
	}
	return reqs, nil
}

func (imp *CSVCategoryImporter) Validate(reqs []*request.CreateCategoryRequest) error {
	for _, req := range reqs {
		err := utils.ValidateKind(req.Kind)
		if err != nil {
			return err
		}

		err = utils.ValidateString(req.Name)
		if err != nil {
			return err
		}
	}
	return nil
}

func (imp *CSVCategoryImporter) Save(ctx context.Context, reqs []*request.CreateCategoryRequest) error {
	for _, req := range reqs {
		err := imp.svc.CreateCategory(ctx, req)
		if err != nil {
			return err
		}
	}
	return nil
}
