package app

import (
	"github.com/niklvrr/FinancialAnalyticsService/internal/config"
	"github.com/niklvrr/FinancialAnalyticsService/internal/infrastructure"
	"github.com/niklvrr/FinancialAnalyticsService/pkg/logger"
	"log"
	"log/slog"
)

type App struct {
	db  *infrastructure.Database
	cfg *config.Config
	log *slog.Logger
}

func NewApp() *App {
	// инициализация конфига
	cfg, err := config.LoadConfig()
	if err != nil {
		log.Fatal(err)
	}

	// инициализация логгера
	logger := logger.NewLog(cfg.App.Env)
	logger.Debug("Логгер инициализирован")

	// инициализация бд
	db, err := infrastructure.NewDB(cfg.Database.URL)
	if err != nil {
		logger.Error(err.Error())
	}
	logger.Debug("База данных инициализирована")

	return &App{
		db:  db,
		cfg: cfg,
		log: logger,
	}
}

func (app *App) Run() error {
	err := setupApp()
	if err != nil {
		app.log.Error(err.Error())
	}
	app.log.Debug("Слои приложения инициализированы")

	return nil
}

func (app *App) Stop() {
	if app.db != nil {
		app.db.Close()
	}

	app.log.Debug("Приложение завершено корректно")
	return
}

func setupApp() error {
	// инициализация репозитория

	// инициализация сервисов

	// инициализация хэндлеров

	// инициализация меню
	return nil
}
