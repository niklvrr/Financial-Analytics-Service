package exporter

import (
	"context"
	"encoding/csv"
	"strings"
)

type CSVBuilder struct {
	rows   [][]string
	header []string
}

func NewCSVBuilder() *CSVBuilder {
	return &CSVBuilder{
		rows:   make([][]string, 0),
		header: make([]string, 0),
	}
}

func (b *CSVBuilder) Begin(ctx context.Context, title string) error {
	b.rows = make([][]string, 0)
	b.header = make([]string, 0)
	return nil
}

func (b *CSVBuilder) AddHeader(ctx context.Context, cols ...string) error {
	b.header = cols
	return nil
}

func (b *CSVBuilder) AddRow(ctx context.Context, values ...string) error {
	b.rows = append(b.rows, values)
	return nil
}

func (b *CSVBuilder) End(ctx context.Context) (*Report, error) {
	var buf strings.Builder
	writer := csv.NewWriter(&buf)

	if len(b.header) > 0 {
		if err := writer.Write(b.header); err != nil {
			return nil, err
		}
	}

	for _, row := range b.rows {
		if err := writer.Write(row); err != nil {
			return nil, err
		}
	}

	writer.Flush()
	if err := writer.Error(); err != nil {
		return nil, err
	}

	return &Report{
		Content: []byte(buf.String()),
	}, nil
}

