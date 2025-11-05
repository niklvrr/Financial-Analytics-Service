package menu

import "github.com/niklvrr/Financial-Analytics-Service/internal/transport/cli/handlers"

func registerCategory(menu *Menu, categoryHandler *handlers.CategoryHandler) {
	subMenuTitle := "=== Меню управляения категориями ==="
	subMenu := NewMenu(subMenuTitle)
	registerCategorySubMenu(subMenu, categoryHandler)

	categoryItem := &MenuItem{
		Key:     1,
		Title:   "Управлять категориями",
		Handler: nil,
		SubMenu: subMenu,
	}

	menu.AddItem(categoryItem)
}

func registerCategorySubMenu(subMenu *Menu, categoryHandler *handlers.CategoryHandler) {
	createMenuItem := &MenuItem{
		Key:     0,
		Title:   "Создать категорию",
		Handler: categoryHandler.CreateCategory,
		SubMenu: nil,
	}

	getMenuItem := &MenuItem{
		Key:     1,
		Title:   "Найти категорию",
		Handler: categoryHandler.GetCategory,
		SubMenu: nil,
	}

	updateMenuItem := &MenuItem{
		Key:     2,
		Title:   "Изменить категорию",
		Handler: categoryHandler.UpdateCategory,
		SubMenu: nil,
	}

	deleteMenuItem := &MenuItem{
		Key:     3,
		Title:   "Удалить категорию",
		Handler: categoryHandler.DeleteCategory,
		SubMenu: nil,
	}

	getAllMenuItem := &MenuItem{
		Key:     4,
		Title:   "Все категории",
		Handler: categoryHandler.GetAllCategories,
		SubMenu: nil,
	}

	subMenu.AddItem(createMenuItem)
	subMenu.AddItem(getMenuItem)
	subMenu.AddItem(getAllMenuItem)
	subMenu.AddItem(updateMenuItem)
	subMenu.AddItem(deleteMenuItem)
	subMenu.AddItem(getAllMenuItem)
}
