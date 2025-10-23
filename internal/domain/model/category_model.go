package model

type Category struct {
	id   int64
	kind string
	name string
}

func NewCategory(id int64, kind string, name string) *Category {
	return &Category{
		id:   id,
		kind: kind,
		name: name,
	}
}
