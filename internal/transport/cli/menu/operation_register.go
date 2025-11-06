package menu

func registerOperationCommands(menu *Menu, commands []Command, titles []string) {
	subMenuTitle := "=== Меню управления операциями ==="
	subMenu := NewMenu(subMenuTitle)
	for i := 0; i < len(commands); i++ {
		command := commands[i]
		title := titles[i]

		item := &MenuItem{
			Key:     len(subMenu.Items) + 1,
			Title:   title,
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
