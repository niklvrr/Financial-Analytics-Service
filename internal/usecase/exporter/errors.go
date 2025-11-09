package exporter

import "errors"

var (
	ErrUnsupportedFormat = errors.New("неподдерживаемый формат экспорта")
	ErrInvalidStrategy   = errors.New("неверная стратегия экспорта")
	ErrEmptyData         = errors.New("нет данных для экспорта")
)

