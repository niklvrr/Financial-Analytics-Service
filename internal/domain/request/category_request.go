package request

// Модель запроса для создания категории
type CreateCategoryRequest struct {
	Kind string
	Name string
}

func NewCreateCategoryRequest(kind, name string) *CreateCategoryRequest {
	return &CreateCategoryRequest{
		Kind: kind,
		Name: name,
	}
}

// Модель запроса для получения категории
type GetCategoryRequest struct {
	Id int64
}

func NewGetCategoryRequest(id int64) *GetCategoryRequest {
	return &GetCategoryRequest{
		Id: id,
	}
}

// Модель запроса для изменения категории
type UpdateCategoryRequest struct {
	Id   int64
	Kind string
	Name string
}

func NewUpdateCategoryRequest(id int64, kind string, name string) *UpdateCategoryRequest {
	return &UpdateCategoryRequest{
		Id:   id,
		Kind: kind,
		Name: name,
	}
}

// Модель запроса для удаления категории
type DeleteCategoryRequest struct {
	Id int64
}

func NewDeleteCategoryRequest(id int64) *DeleteCategoryRequest {
	return &DeleteCategoryRequest{
		Id: id,
	}
}
