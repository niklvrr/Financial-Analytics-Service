package di

import (
	"context"
	"github.com/niklvrr/Financial-Analytics-Service/internal/transport/cli/command/bank_account_commands"
	"github.com/niklvrr/Financial-Analytics-Service/internal/transport/cli/command/category_commands"
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
	bankAccountRepo := infrastructure.NewBankAccountRepo(db.Db)
	categoryRepo := infrastructure.NewCategoryRepo(db.Db)
	operationRepo := infrastructure.NewOperationRepo(db.Db)

	// инициализация сервисов
	bankAccountService := service.NewBankAccountService(bankAccountRepo)
	categoryService := service.NewCategoryService(categoryRepo)
	operationService := service.NewOperationService(operationRepo)

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
		createBankAccountCommand,
		getBankAccountCommand,
		updateBankAccountCommand,
		deleteBankAccountCommand,
		getAllBankAccountsCommand,
	}

	// инициализация команд категории
	createCategoryCommand := category_commands.NewCreateCategoryCommand(categoryFacade)
	getCategoryCommand := category_commands.NewGetCategoryCommand(categoryFacade)
	updateCategoryCommand := category_commands.NewUpdateCategoryCommand(categoryFacade)
	deleteCategoryCommand := category_commands.NewDeleteCategoryCommand(categoryFacade)
	getAllCategoriesCommand := category_commands.NewGetAllCategoriesCommand(categoryFacade)
	categoryCommands := []menu.Command{
		createCategoryCommand,
		getCategoryCommand,
		updateCategoryCommand,
		deleteCategoryCommand,
		getAllCategoriesCommand,
	}

	// инициализация команд операций
	createOperationCommand := operation_commands.NewCreateOperationCommand(operationFacade)
	getOperationCommand := operation_commands.NewGetOperationCommand(operationFacade)
	updateOperationCommand := operation_commands.NewUpdateOperationCommand(operationFacade)
	deleteOperationCommand := operation_commands.NewDeleteOperationCommand(operationFacade)
	getAllOperationsCommand := operation_commands.NewGetAllOperationsCommand(operationFacade)
	operationCommands := []menu.Command{
		createOperationCommand,
		getOperationCommand,
		updateOperationCommand,
		deleteOperationCommand,
		getAllOperationsCommand,
	}

	//// инициализация хэндлеров
	//bankAccountHandler := handlers.NewBankAccountHandler(bankAccountService)
	//categoryHandler := handlers.NewCategoryHandler(categoryService)
	//operationHandler := handlers.NewOperationHandler(operationService)

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

func (c *Container) Run() error {
	return c.Menu.Run(c.Ctx)
}
