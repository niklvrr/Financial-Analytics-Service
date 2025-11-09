package exporter

import (
	"context"
	"gopkg.in/yaml.v3"
)

type YAMLBuilder struct {
	rows   []map[string]string
	header []string
}

func NewYAMLBuilder() *YAMLBuilder {
	return &YAMLBuilder{
		rows:   make([]map[string]string, 0),
		header: make([]string, 0),
	}
}

func (b *YAMLBuilder) Begin(ctx context.Context, title string) error {
	b.rows = make([]map[string]string, 0)
	b.header = make([]string, 0)
	return nil
}

func (b *YAMLBuilder) AddHeader(ctx context.Context, cols ...string) error {
	b.header = cols
	return nil
}

func (b *YAMLBuilder) AddRow(ctx context.Context, values ...string) error {
	if len(b.header) == 0 {
		return nil
	}

	row := make(map[string]string)
	for i, col := range b.header {
		if i < len(values) {
			row[col] = values[i]
		} else {
			row[col] = ""
		}
	}
	b.rows = append(b.rows, row)
	return nil
}

func (b *YAMLBuilder) End(ctx context.Context) (*Report, error) {
	data, err := yaml.Marshal(b.rows)
	if err != nil {
		return nil, err
	}

	return &Report{
		Content: data,
	}, nil
}

