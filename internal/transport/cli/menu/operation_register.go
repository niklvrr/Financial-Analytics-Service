package menu

func registerOperationCommands(menu *Menu, commands []Command) {
	subMenuTitle := "=== Меню управления операциями ==="
	subMenu := NewMenu(subMenuTitle)
	for i := 0; i < len(commands); i++ {
		command := commands[i]

		item := &MenuItem{
			Key:     len(subMenu.Items) + 1,
			Title:   command.Title(),
			Command: command,
			SubMenu: subMenu,
		}
		subMenu.AddItem(item)
	}

	operationItem := &MenuItem{
		Key:     len(menu.Items) + 1,
		Title:   "Управлять операциями",
		Command: nil,
		SubMenu: subMenu,
	}
	menu.AddItem(operationItem)
}
