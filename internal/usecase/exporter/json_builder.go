package exporter

import (
	"context"
	"encoding/json"
)

type JSONBuilder struct {
	rows   []map[string]string
	header []string
}

func NewJSONBuilder() *JSONBuilder {
	return &JSONBuilder{
		rows:   make([]map[string]string, 0),
		header: make([]string, 0),
	}
}

func (b *JSONBuilder) Begin(ctx context.Context, title string) error {
	b.rows = make([]map[string]string, 0)
	b.header = make([]string, 0)
	return nil
}

func (b *JSONBuilder) AddHeader(ctx context.Context, cols ...string) error {
	b.header = cols
	return nil
}

func (b *JSONBuilder) AddRow(ctx context.Context, values ...string) error {
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

func (b *JSONBuilder) End(ctx context.Context) (*Report, error) {
	data, err := json.MarshalIndent(b.rows, "", "  ")
	if err != nil {
		return nil, err
	}

	return &Report{
		Content: data,
	}, nil
}

