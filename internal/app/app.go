package app

import (
	"context"
	"errors"
	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	"github.com/niklvrr/Financial-Analytics-Service/internal/config"
	"github.com/niklvrr/Financial-Analytics-Service/pkg/di"
	"github.com/niklvrr/Financial-Analytics-Service/pkg/logger"
	"log"
	"log/slog"
)

type App struct {
	ctx       context.Context
	cfg       *config.Config
	log       *slog.Logger
	container *di.Container
}

func NewApp(c context.Context) *App {
	// инициализация конфига
	cfg, err := config.LoadConfig()
	if err != nil {
		log.Fatal(err)
	}

	// инициализация логгера
	log := logger.NewLog(cfg.App.Env)
	log.Debug("Логгер инициализирован")

	// DI контейнер (инициализирует БД, репозитории, сервисы, хэндлеры и меню)
	container, err := di.New(c)
	if err != nil {
		log.Error(err.Error())
	}
	log.Debug("Слои приложения инициализированы")

	mustRunMigrations(cfg.Database.URL, log)

	return &App{
		ctx:       c,
		cfg:       cfg,
		log:       log,
		container: container,
	}
}

func (app *App) Run() error {
	if err := app.container.Run(); err != nil {
		app.log.Error(err.Error())
		return err
	}
	return nil
}

func (app *App) Stop() {
	if app.container != nil {
		app.container.Close()
	}

	app.log.Debug("Приложение завершено корректно")
	return
}

func mustRunMigrations(dbUrl string, logger *slog.Logger) {
	if dbUrl == "" {
		logger.Error("dbUrl is empty")
		return
	}

	mg, err := migrate.New(
		"file://migrations",
		dbUrl,
	)
	if err != nil {
		logger.Error("migration init err", err)
		return
	}

	if err := mg.Up(); err != nil && !errors.Is(err, migrate.ErrNoChange) {
		logger.Error("migration run err", err)
		return
	}

	logger.Debug("migration run ok")
}
