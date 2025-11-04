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

type HandlerFunc func(ctx context.Context) error

type MenuItem struct {
	Key     int
	Title   string
	Handler HandlerFunc
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

func (m *Menu) AddItem(item *MenuItem) {
	m.Items = append(m.Items, item)
}

func (m *Menu) Run(ctx context.Context) error {
	for {
		itemsCount := len(m.Items)

		fmt.Printf("=== %s ===\n", m.Title)
		for i := 1; i <= itemsCount; i++ {
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

		item := m.Items[c]
		if item.SubMenu != nil {
			err := item.SubMenu.Run(ctx)
			if err != nil {
				return err
			}
		} else if item.Handler != nil {
			err := item.Handler(ctx)
			if err != nil {
				return err
			}
		}
	}
}

func inputCommand(in *bufio.Reader, itemsCount int) (int, error) {
	fmt.Print("Введите команду: ")
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
