package request

// Модель запроса для создания категории
type CreateCategoryRequest struct {
	Kind string
	Name string
}

// Модель запроса для получения категории
type GetCategoryRequest struct {
	Id int64
}

// Модель запроса для изменения категории
type UpdateCategoryRequest struct {
	Id   int64
	Kind string
	Name string
}

// Модель запроса для удаления категории
type DeleteCategoryRequest struct {
	Id int64
}
