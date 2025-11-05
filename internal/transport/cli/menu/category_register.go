package menu

import "github.com/niklvrr/Financial-Analytics-Service/internal/transport/cli/handlers"

func registerCategory(menu *Menu, categoryHandler *handlers.CategoryHandler) {
	subMenuTitle := "=== Меню управляения категориями ==="
	subMenu := NewMenu(subMenuTitle)
	registerCategorySubMenu(subMenu, categoryHandler)

	categoryItem := &MenuItem{
		Key:     len(menu.Items) + 1,
		Title:   "Управлять категориями",
		Handler: nil,
		SubMenu: subMenu,
	}

	menu.AddItem(categoryItem)
}

func registerCategorySubMenu(subMenu *Menu, categoryHandler *handlers.CategoryHandler) {
	createMenuItem := &MenuItem{
		Key:     len(subMenu.Items) + 1,
		Title:   "Создать категорию",
		Handler: categoryHandler.CreateCategory,
		SubMenu: nil,
	}
	subMenu.AddItem(createMenuItem)

	getMenuItem := &MenuItem{
		Key:     len(subMenu.Items) + 1,
		Title:   "Найти категорию",
		Handler: categoryHandler.GetCategory,
		SubMenu: nil,
	}
	subMenu.AddItem(getMenuItem)

	updateMenuItem := &MenuItem{
		Key:     len(subMenu.Items) + 1,
		Title:   "Изменить категорию",
		Handler: categoryHandler.UpdateCategory,
		SubMenu: nil,
	}
	subMenu.AddItem(updateMenuItem)

	deleteMenuItem := &MenuItem{
		Key:     len(subMenu.Items) + 1,
		Title:   "Удалить категорию",
		Handler: categoryHandler.DeleteCategory,
		SubMenu: nil,
	}
	subMenu.AddItem(deleteMenuItem)

	getAllMenuItem := &MenuItem{
		Key:     len(subMenu.Items) + 1,
		Title:   "Все категории",
		Handler: categoryHandler.GetAllCategories,
		SubMenu: nil,
	}
	subMenu.AddItem(getAllMenuItem)
}
