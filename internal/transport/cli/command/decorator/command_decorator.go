package decorator

import (
	"context"
	"github.com/niklvrr/Financial-Analytics-Service/internal/transport/cli/menu"
	"log/slog"
	"time"
)

type LoggingDecorator struct {
	c   menu.Command
	log *slog.Logger
}

func (d *LoggingDecorator) Execute(ctx context.Context) error {
	start := time.Now()
	err := d.c.Execute(ctx)
	if err != nil {
		d.log.Debug("Команда не выполнена. Ошибка: " + err.Error())
	} else {
		d.log.Debug("Команда выполнена. Время: " + time.Since(start).String())
	}
	return err
}

func WithLogging(command menu.Command, log *slog.Logger) menu.Command {
	return &LoggingDecorator{
		c:   command,
		log: log,
	}
}
