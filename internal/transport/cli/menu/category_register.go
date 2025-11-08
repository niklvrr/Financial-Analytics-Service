package menu

func registerCategoryCommands(menu *Menu, commands []Command) {
	subMenuTitle := "=== Меню управляения категориями ==="
	subMenu := NewMenu(subMenuTitle)
	for i := 0; i < len(commands); i++ {
		command := commands[i]

		item := &MenuItem{
			Key:     len(subMenu.Items) + 1,
			Title:   command.Title(),
			Command: command,
			SubMenu: nil,
		}
		subMenu.AddItem(item)
	}

	categoryItem := &MenuItem{
		Key:     len(menu.Items) + 1,
		Title:   "Управлять категориями",
		Command: nil,
		SubMenu: subMenu,
	}

	menu.AddItem(categoryItem)
}
