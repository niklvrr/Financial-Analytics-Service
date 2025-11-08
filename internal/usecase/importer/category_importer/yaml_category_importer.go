package category_importer

import (
	"context"
	"github.com/niklvrr/Financial-Analytics-Service/internal/domain/request"
	"github.com/niklvrr/Financial-Analytics-Service/internal/usecase/service"
	"github.com/niklvrr/Financial-Analytics-Service/pkg/utils"
	"gopkg.in/yaml.v3"
	"os"
)

type YamlCategoryImporter struct {
	svc *service.CategoryService
}

func NewYamlCategoryImporter(svc *service.CategoryService) *YamlCategoryImporter {
	return &YamlCategoryImporter{
		svc: svc,
	}
}

func (imp *YamlCategoryImporter) Load(path string) ([]*request.CreateCategoryRequest, error) {
	data, err := os.ReadFile("config.yaml")
	if err != nil {
		return nil, err
	}

	var reqs []*request.CreateCategoryRequest
	err = yaml.Unmarshal(data, &reqs)
	if err != nil {
		return nil, err
	}
	return reqs, nil
}

func (imp *YamlCategoryImporter) Validate(reqs []*request.CreateCategoryRequest) error {
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

func (imp *YamlCategoryImporter) Save(ctx context.Context, reqs []*request.CreateCategoryRequest) error {
	for _, req := range reqs {
		err := imp.svc.CreateCategory(ctx, req)
		if err != nil {
			return err
		}
	}
	return nil
}
