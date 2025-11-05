package di

import (
	"context"
	"log/slog"

	"github.com/niklvrr/Financial-Analytics-Service/internal/config"
	"github.com/niklvrr/Financial-Analytics-Service/internal/infrastructure"
	"github.com/niklvrr/Financial-Analytics-Service/internal/transport/cli/handlers"
	"github.com/niklvrr/Financial-Analytics-Service/internal/transport/cli/menu"
	"github.com/niklvrr/Financial-Analytics-Service/internal/usecase"
	"github.com/niklvrr/Financial-Analytics-Service/pkg/logger"
)

type Container struct {
	Ctx context.Context
	Cfg *config.Config
	Log *slog.Logger

	DB *infrastructure.Database

	// репозитории
	BankAccountRepo *infrastructure.BankAccountRepo
	CategoryRepo    *infrastructure.CategoryRepo
	OperationRepo   *infrastructure.OperationRepo

	// сервисы
	BankAccountService *usecase.BankAccountService
	CategoryService    *usecase.CategoryService
	OperationService   *usecase.OperationService

	// хэндлеры
	BankAccountHandler *handlers.BankAccountHandler
	CategoryHandler    *handlers.CategoryHandler
	OperationHandler   *handlers.OperationHandler

	Menu *menu.Menu
}

func New(ctx context.Context) (*Container, error) {
	// инициализация конфига
	cfg, err := config.LoadConfig()
	if err != nil {
		return nil, err
	}

	// инициализация логгера
	lg := logger.NewLog(cfg.App.Env)

	// инициализация базы данных
	db, err := infrastructure.NewDB(ctx, cfg.Database.URL)
	if err != nil {
		lg.Error(err.Error())
		return nil, err
	}

	// инициализация репозиториев
	bankAccountRepo := infrastructure.NewBankAccountRepo(db.Db)
	categoryRepo := infrastructure.NewCategoryRepo(db.Db)
	operationRepo := infrastructure.NewOperationRepo(db.Db)

	// инициализация сервисов
	bankAccountService := usecase.NewBankAccountService(bankAccountRepo)
	categoryService := usecase.NewCategoryService(categoryRepo)
	operationService := usecase.NewOperationService(operationRepo)

	// инициализация хэндлеров
	bankAccountHandler := handlers.NewBankAccountHandler(bankAccountService)
	categoryHandler := handlers.NewCategoryHandler(categoryService)
	operationHandler := handlers.NewOperationHandler(operationService)

	// инициализация меню
	root := menu.NewMenu("=== Главное меню ===")
	root.Build(bankAccountHandler, categoryHandler, operationHandler)

	return &Container{
		Ctx: ctx,
		Cfg: cfg,
		Log: lg,
		DB:  db,

		BankAccountRepo: bankAccountRepo,
		CategoryRepo:    categoryRepo,
		OperationRepo:   operationRepo,

		BankAccountService: bankAccountService,
		CategoryService:    categoryService,
		OperationService:   operationService,

		BankAccountHandler: bankAccountHandler,
		CategoryHandler:    categoryHandler,
		OperationHandler:   operationHandler,

		Menu: root,
	}, nil
}

func (c *Container) Close() {
	if c.DB != nil {
		c.DB.Close()
	}
}

func (c *Container) Run() error {
	return c.Menu.Run(c.Ctx)
}
