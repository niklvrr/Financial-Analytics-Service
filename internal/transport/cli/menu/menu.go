package menu

import (
	"bufio"
	"context"
	"errors"
	"fmt"
	"os"
	"strconv"
	"strings"
)

var (
	incorrectCommandError = errors.New("Неизвестная команда")
)

type Command interface {
	Execute(ctx context.Context) error
}

type HandlerFunc func(ctx context.Context) error

type MenuItem struct {
	Key     int
	Title   string
	Command Command
	SubMenu *Menu
}

type Menu struct {
	Title string
	Items []*MenuItem
	In    *bufio.Reader
}

func NewMenu(title string) *Menu {
	return &Menu{
		Title: title,
		Items: []*MenuItem{},
		In:    bufio.NewReader(os.Stdin),
	}
}

func (m *Menu) Build(bankAccountCommands, categoryCommands, operationCommands []Command) {
	bankAccountCommandsTitles := []string{
		"Создать банковский счет",
		"Найти банковкий счет",
		"Изменить банковский счет",
		"Удалить банковский счет",
		"Найти все банковсие счета",
	}
	registerBankAccountCommands(m, bankAccountCommands, bankAccountCommandsTitles)

	categoryCommandsTitles := []string{
		"Создать категорию",
		"Найти категорию",
		"Изменить категорию",
		"Удалить категорию",
		"Все категории",
	}
	registerCategoryCommands(m, categoryCommands, categoryCommandsTitles)

	operationCommandsTitles := []string{
		"Создать операцию",
		"Найти операцию",
		"Изменить операцию",
		"Удалить операцию",
		"Найти все опарции",
	}
	registerOperationCommands(m, operationCommands, operationCommandsTitles)
}

func (m *Menu) AddItem(item *MenuItem) {
	m.Items = append(m.Items, item)
}

func (m *Menu) Run(ctx context.Context) error {
	for {
		itemsCount := len(m.Items)

		fmt.Printf("%s\n", m.Title)
		for i := 0; i < itemsCount; i++ {
			fmt.Printf("%d) %s\n", m.Items[i].Key, m.Items[i].Title)
		}
		fmt.Printf("%d) Выход\n", itemsCount+1)

		fmt.Print("Выберете команду: ")
		c, err := inputCommand(m.In, itemsCount+1)
		if err != nil {
			fmt.Printf("Ошибка: %v\n", err)
			continue
		}

		if c == itemsCount+1 {
			return nil
		}

		item := m.Items[c-1]
		if item.SubMenu != nil {
			err := item.SubMenu.Run(ctx)
			if err != nil {
				return err
			}
		} else if item.Command != nil {
			err := item.Command.Execute(ctx)
			if err != nil {
				return err
			}
		}
	}
}

func inputCommand(in *bufio.Reader, itemsCount int) (int, error) {
	input, err := in.ReadString('\n')
	if err != nil {
		return 0, err
	}

	input = strings.TrimSpace(input)
	num, err := strconv.Atoi(input)
	if err != nil {
		return 0, err
	}

	if num <= 0 || num > itemsCount {
		return 0, incorrectCommandError
	}

	return num, nil
}
