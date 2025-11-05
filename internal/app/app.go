package app

import (
	"context"
	"errors"
	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/niklvrr/Financial-Analytics-Service/internal/config"
	"github.com/niklvrr/Financial-Analytics-Service/internal/infrastructure"
	"github.com/niklvrr/Financial-Analytics-Service/internal/transport/cli/handlers"
	"github.com/niklvrr/Financial-Analytics-Service/internal/transport/cli/menu"
	"github.com/niklvrr/Financial-Analytics-Service/internal/usecase"
	"github.com/niklvrr/Financial-Analytics-Service/pkg/logger"
	"log"
	"log/slog"
)

type App struct {
	ctx context.Context
	db  *infrastructure.Database
	cfg *config.Config
	log *slog.Logger
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

	// инициализация бд
	db, err := infrastructure.NewDB(c, cfg.Database.URL)
	if err != nil {
		log.Error(err.Error())
	}
	log.Debug("База данных инициализирована")

	mustRunMigrations(cfg.Database.URL, log)

	return &App{
		ctx: c,
		db:  db,
		cfg: cfg,
		log: log,
	}
}

func (app *App) Run() error {
	err := setupApp(app.db.Db, app.ctx)
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

func setupApp(db *pgxpool.Pool, ctx context.Context) error {
	// инициализация репозитория
	bankAccountRepo := infrastructure.NewBankAccountRepo(db)
	categoryRepo := infrastructure.NewCategoryRepo(db)
	operationRepo := infrastructure.NewOperationRepo(db)

	// инициализация сервисов
	bankAccountService := usecase.NewBankAccountService(bankAccountRepo)
	categoryService := usecase.NewCategoryService(categoryRepo)
	operationService := usecase.NewOperationService(operationRepo)

	// инициализация хэндлеров
	bankAccountHandler := handlers.NewBankAccountHandler(bankAccountService)
	categoryHandler := handlers.NewCategoryHandler(categoryService)
	operationHandler := handlers.NewOperationHandler(operationService)

	// инициализация меню
	menuTitle := "=== Главное меню ==="
	m := menu.NewMenu(menuTitle)
	m.Build(bankAccountHandler, categoryHandler, operationHandler)
	err := m.Run(ctx)
	if err != nil {
		return err
	}

	return nil
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
