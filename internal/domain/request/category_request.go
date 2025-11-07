package request

// Модель запроса для создания категории
type CreateCategoryRequest struct {
	Kind string `csv:"kind" json:"kind" yaml:"kind"`
	Name string `csv:"name" json:"name" yaml:"name"`
}

// Модель запроса для получения категории
type GetCategoryRequest struct {
	Id int64 `csv:"id" json:"id" yaml:"id"`
}

// Модель запроса для изменения категории
type UpdateCategoryRequest struct {
	Id   int64  `csv:"id" json:"id" yaml:"id"`
	Kind string `csv:"kind" json:"kind" yaml:"kind"`
	Name string `csv:"name" json:"name" yaml:"name"`
}

// Модель запроса для удаления категории
type DeleteCategoryRequest struct {
	Id int64 `csv:"id" json:"id" yaml:"id"`
}
