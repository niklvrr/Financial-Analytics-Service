package menu

import "github.com/niklvrr/Financial-Analytics-Service/internal/transport/cli/handlers"

func registerOperation(menu *Menu, operationHandler *handlers.OperationHandler) {
	subMenuTitle := "=== Меню управления операциями ==="
	subMenu := NewMenu(subMenuTitle)
	registerOperationSubMenu(subMenu, operationHandler)

	operationItem := &MenuItem{
		Key:     len(menu.Items) + 1,
		Title:   "Управлять операциями",
		Handler: nil,
		SubMenu: subMenu,
	}
	menu.AddItem(operationItem)
}

func registerOperationSubMenu(subMenu *Menu, operationHandler *handlers.OperationHandler) {
	createOperationItem := &MenuItem{
		Key:     0,
		Title:   "Создать операцию",
		Handler: operationHandler.CreateOperation,
		SubMenu: nil,
	}
	subMenu.AddItem(createOperationItem)

	getOperationItem := &MenuItem{
		Key:     1,
		Title:   "Найти операцию",
		Handler: operationHandler.GetOperation,
		SubMenu: nil,
	}
	subMenu.AddItem(getOperationItem)

	updateOperationItem := &MenuItem{
		Key:     2,
		Title:   "Изменить операцию",
		Handler: operationHandler.UpdateOperation,
		SubMenu: nil,
	}
	subMenu.AddItem(updateOperationItem)

	deleteOperationItem := &MenuItem{
		Key:     3,
		Title:   "Удалить операцию",
		Handler: operationHandler.DeleteOperation,
		SubMenu: nil,
	}
	subMenu.AddItem(deleteOperationItem)

	getAllOperationItem := &MenuItem{
		Key:     4,
		Title:   "Найти все опарции",
		Handler: operationHandler.GetAllOperations,
		SubMenu: nil,
	}
	subMenu.AddItem(getAllOperationItem)
}
