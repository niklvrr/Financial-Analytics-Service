package exporter

import (
	"context"
)

type Builder interface {
	Begin(ctx context.Context, title string) error
	AddHeader(ctx context.Context, cols ...string) error
	AddRow(ctx context.Context, values ...string) error
	End(ctx context.Context) (*Report, error)
}

type Report struct {
	Content []byte
}

func NewBuilder(format string) (Builder, error) {
	switch format {
	case ".csv":
		return NewCSVBuilder(), nil
	case ".json":
		return NewJSONBuilder(), nil
	case ".yaml":
		return NewYAMLBuilder(), nil
	default:
		return nil, ErrUnsupportedFormat
	}
}
