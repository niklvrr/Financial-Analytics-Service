package menu

func registerBankAccountCommands(menu *Menu, commands []Command) {
	subMenuTitle := "=== Меню управления счетами ==="
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

	bankAccountItem := &MenuItem{
		Key:     len(menu.Items) + 1,
		Title:   "Упавлять банковскими счетами",
		Command: nil,
		SubMenu: subMenu,
	}
	menu.AddItem(bankAccountItem)
}
