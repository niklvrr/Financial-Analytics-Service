package model

import (
	"errors"
	"time"
)

var (
	incorrectIdError      = errors.New("Ошибка некорректный уникальный номер")
	incorrectNameError    = errors.New("Ошибка некорректное имя")
	incorrectBalanceError = errors.New("Ошибка некорректный баланс")
	incorrectKindError    = errors.New("Ошибка некорректный тип")
	incorrectAmountError  = errors.New("Ошибка некорректная сумма")
)

const (
	incomeKindRus      = "доход"
	incomeKindEng      = "income"
	expenditureKindRus = "расход"
	expenditureKindEng = "expenditure"
)

type DomainFabric struct {
}

func NewDomainFabric() *DomainFabric {
	return &DomainFabric{}
}

func (f *DomainFabric) BuildBankAccount(id int64, name string, balance float64) (*BankAccount, error) {
	if id < 0 {
		return nil, incorrectIdError
	}

	if len(name) == 0 {
		return nil, incorrectNameError
	}

	if balance < 0 {
		return nil, incorrectBalanceError
	}

	return NewBankAccount(id, name, balance), nil
}

func (f *DomainFabric) BuildCategory(id int64, kind, name string) (*Category, error) {
	if id < 0 {
		return nil, incorrectIdError
	}

	if len(name) == 0 {
		return nil, incorrectNameError
	}

	if kind != incomeKindRus && kind != expenditureKindRus && kind != incomeKindEng && kind != expenditureKindEng {
		return nil, incorrectKindError
	}

	return NewCategory(id, name, kind), nil
}

func (f *DomainFabric) BuildOperation(
	id int64,
	kind string,
	bankAccountId int64,
	amount float64,
	date time.Time,
	description string,
	categoryId int64) (*Operation, error) {
	if id < 0 {
		return nil, incorrectIdError
	}

	if kind != incomeKindRus && kind != expenditureKindRus && kind != incomeKindEng && kind != expenditureKindEng {
		return nil, incorrectKindError
	}

	if bankAccountId < 0 {
		return nil, incorrectIdError
	}

	if amount < 0 {
		return nil, incorrectAmountError
	}

	if categoryId < 0 {
		return nil, incorrectIdError
	}

	return NewOperation(id, kind, bankAccountId, amount, date, description, categoryId), nil
}
