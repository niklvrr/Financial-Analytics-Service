package app

import (
	"context"
	"github.com/niklvrr/Financial-Analytics-Service/internal/infrastructure/proxy"
	"github.com/niklvrr/Financial-Analytics-Service/internal/infrastructure/repository"
	"github.com/niklvrr/Financial-Analytics-Service/internal/transport/cli/command/bank_account_commands"
	"github.com/niklvrr/Financial-Analytics-Service/internal/transport/cli/command/category_commands"
	"github.com/niklvrr/Financial-Analytics-Service/internal/transport/cli/command/decorator"
	"github.com/niklvrr/Financial-Analytics-Service/internal/transport/cli/command/operation_commands"
	"github.com/niklvrr/Financial-Analytics-Service/internal/usecase/facade"
	"github.com/niklvrr/Financial-Analytics-Service/internal/usecase/service"
	"log/slog"

	"github.com/niklvrr/Financial-Analytics-Service/internal/config"
	"github.com/niklvrr/Financial-Analytics-Service/internal/infrastructure"
	"github.com/niklvrr/Financial-Analytics-Service/internal/transport/cli/menu"
	"github.com/niklvrr/Financial-Analytics-Service/pkg/logger"
)

type Container struct {
	Ctx context.Context
	Cfg *config.Config
	Log *slog.Logger

	DB infrastructure.Database

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
	db, err := infrastructure.NewPostgreSqlDb(ctx, cfg.Database.URL)
	if err != nil {
		lg.Error(err.Error())
		return nil, err
	}

	// инициализация репозиториев
	bankAccountRepo := repository.NewBankAccountRepo(db.Db)
	categoryRepo := repository.NewCategoryRepo(db.Db)
	operationRepo := repository.NewOperationRepo(db.Db)

	// инициализация прокси
	bankAccountProxy := proxy.NewBankAccountProxy(bankAccountRepo, lg)
	categoryProxy := proxy.NewCategoryProxy(categoryRepo, lg)
	operationProxy := proxy.NewOperationProxy(operationRepo, lg)

	// инициализация сервисов
	bankAccountService := service.NewBankAccountService(bankAccountProxy)
	categoryService := service.NewCategoryService(categoryProxy)
	operationService := service.NewOperationService(operationProxy)

	// инициализация фасадов
	bankAccountFacade := facade.NewBankAccountFacade(bankAccountService)
	categoryFacade := facade.NewCategoryFacade(categoryService)
	operationFacade := facade.NewOperationFacade(operationService)

	// инициализация команд банковского счета
	createBankAccountCommand := bank_account_commands.NewCreateBankAccountCommand(bankAccountFacade)
	getBankAccountCommand := bank_account_commands.NewGetBankAccountCommand(bankAccountFacade)
	updateBankAccountCommand := bank_account_commands.NewUpdateBankAccountCommand(bankAccountFacade)
	deleteBankAccountCommand := bank_account_commands.NewDeleteBankAccountCommand(bankAccountFacade)
	getAllBankAccountsCommand := bank_account_commands.NewGetAllBankAccountsCommand(bankAccountFacade)
	bankAccountCommands := []menu.Command{
		decorator.WithLogging(createBankAccountCommand, lg),
		decorator.WithLogging(getBankAccountCommand, lg),
		decorator.WithLogging(updateBankAccountCommand, lg),
		decorator.WithLogging(deleteBankAccountCommand, lg),
		decorator.WithLogging(getAllBankAccountsCommand, lg),
	}

	// инициализация команд категории
	createCategoryCommand := category_commands.NewCreateCategoryCommand(categoryFacade)
	getCategoryCommand := category_commands.NewGetCategoryCommand(categoryFacade)
	updateCategoryCommand := category_commands.NewUpdateCategoryCommand(categoryFacade)
	deleteCategoryCommand := category_commands.NewDeleteCategoryCommand(categoryFacade)
	getAllCategoriesCommand := category_commands.NewGetAllCategoriesCommand(categoryFacade)
	categoryCommands := []menu.Command{
		decorator.WithLogging(createCategoryCommand, lg),
		decorator.WithLogging(getCategoryCommand, lg),
		decorator.WithLogging(updateCategoryCommand, lg),
		decorator.WithLogging(deleteCategoryCommand, lg),
		decorator.WithLogging(getAllCategoriesCommand, lg),
	}

	// инициализация команд операций
	createOperationCommand := operation_commands.NewCreateOperationCommand(operationFacade)
	getOperationCommand := operation_commands.NewGetOperationCommand(operationFacade)
	updateOperationCommand := operation_commands.NewUpdateOperationCommand(operationFacade)
	deleteOperationCommand := operation_commands.NewDeleteOperationCommand(operationFacade)
	getAllOperationsCommand := operation_commands.NewGetAllOperationsCommand(operationFacade)
	operationCommands := []menu.Command{
		decorator.WithLogging(createOperationCommand, lg),
		decorator.WithLogging(getOperationCommand, lg),
		decorator.WithLogging(updateOperationCommand, lg),
		decorator.WithLogging(deleteOperationCommand, lg),
		decorator.WithLogging(getAllOperationsCommand, lg),
	}

	// инициализация меню
	root := menu.NewMenu("=== Главное меню ===")
	root.Build(bankAccountCommands, categoryCommands, operationCommands)

	return &Container{
		Ctx:  ctx,
		Cfg:  cfg,
		Log:  lg,
		Menu: root,
	}, nil
}

func (c *Container) Close() {
	if c.DB != nil {
		c.DB.Close()
	}
}

func (c *Container) Run() {
	c.Menu.Run(c.Ctx)
}
