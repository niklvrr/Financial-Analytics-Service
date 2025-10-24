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

func (c *Category) ID() int64 {
	return c.id
}

func (c *Category) SetID(id int64) {
	c.id = id
}

func (c *Category) Kind() string {
	return c.kind
}

func (c *Category) SetKind(kind string) {
	c.kind = kind
}

func (c *Category) Name() string {
	return c.name
}

func (c *Category) SetName(name string) {
	c.name = name
}
