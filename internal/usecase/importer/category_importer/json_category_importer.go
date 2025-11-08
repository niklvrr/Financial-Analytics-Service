package category_importer

import (
	"context"
	"encoding/json"
	"github.com/niklvrr/Financial-Analytics-Service/internal/domain/request"
	"github.com/niklvrr/Financial-Analytics-Service/internal/usecase/service"
	"github.com/niklvrr/Financial-Analytics-Service/pkg/utils"
	"os"
)

type JSONCategoryImporter struct {
	svc *service.CategoryService
}

func NewJSONCategoryImporter(svc *service.CategoryService) *JSONCategoryImporter {
	return &JSONCategoryImporter{svc: svc}
}

func (imp *JSONCategoryImporter) Load(path string) ([]*request.CreateCategoryRequest, error) {
	data, err := os.ReadFile(path)
	if err != nil {
		panic(err)
	}

	var reqs []*request.CreateCategoryRequest
	err = json.Unmarshal(data, &reqs)
	if err != nil {
		return nil, err
	}
	return reqs, nil
}

func (imp *JSONCategoryImporter) Validate(reqs []*request.CreateCategoryRequest) error {
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

func (imp *JSONCategoryImporter) Save(ctx context.Context, reqs []*request.CreateCategoryRequest) error {
	for _, req := range reqs {
		err := imp.svc.CreateCategory(ctx, req)
		if err != nil {
			return err
		}
	}
	return nil
}
