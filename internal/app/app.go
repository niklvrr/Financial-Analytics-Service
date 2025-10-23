package app

import (
	"github.com/niklvrr/Financial-Analytics-Service/internal/config"
	"github.com/niklvrr/Financial-Analytics-Service/internal/infrastructure"
	"github.com/niklvrr/Financial-Analytics-Service/pkg/logger"
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
	log := logger.NewLog(cfg.App.Env)
	log.Debug("Логгер инициализирован")

	// инициализация бд
	db, err := infrastructure.NewDB(cfg.Database.URL)
	if err != nil {
		log.Error(err.Error())
	}
	log.Debug("База данных инициализирована")

	return &App{
		db:  db,
		cfg: cfg,
		log: log,
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
