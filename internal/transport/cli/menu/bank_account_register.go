package menu

import "github.com/niklvrr/Financial-Analytics-Service/internal/transport/cli/handlers"

func registerBankAccount(menu *Menu, bankAccountHandler *handlers.BankAccountHandler) {
	subMenuTitle := "=== Меню управления счетами ==="
	subMenu := NewMenu(subMenuTitle)
	registerBankAccountSubMenu(subMenu, bankAccountHandler)

	bankAccountItem := &MenuItem{
		Key:     len(menu.Items) + 1,
		Title:   "Упавлять банковскими счетами",
		Handler: nil,
		SubMenu: subMenu,
	}
	menu.AddItem(bankAccountItem)
}

func registerBankAccountSubMenu(subMenu *Menu, bankAccountHandler *handlers.BankAccountHandler) {
	createBankAccountItem := &MenuItem{
		Key:     len(subMenu.Items) + 1,
		Title:   "Создать банковский счет",
		Handler: bankAccountHandler.CreateBankAccount,
		SubMenu: nil,
	}
	subMenu.AddItem(createBankAccountItem)

	getBankAccountItem := &MenuItem{
		Key:     len(subMenu.Items) + 1,
		Title:   "Найти банковкий счет",
		Handler: bankAccountHandler.GetBankAccount,
		SubMenu: nil,
	}
	subMenu.AddItem(getBankAccountItem)

	updateBankAccountItem := &MenuItem{
		Key:     len(subMenu.Items) + 1,
		Title:   "Изменить банковский счет",
		Handler: bankAccountHandler.UpdateBankAccount,
		SubMenu: nil,
	}
	subMenu.AddItem(updateBankAccountItem)

	deleteBankAccountItem := &MenuItem{
		Key:     len(subMenu.Items) + 1,
		Title:   "Удалить банковский счет",
		Handler: bankAccountHandler.DeleteBankAccount,
		SubMenu: nil,
	}
	subMenu.AddItem(deleteBankAccountItem)

	getAllBankAccountsItem := &MenuItem{
		Key:     len(subMenu.Items) + 1,
		Title:   "Найти все банковсие счета",
		Handler: bankAccountHandler.GetAllBankAccounts,
		SubMenu: nil,
	}
	subMenu.AddItem(getAllBankAccountsItem)
}
