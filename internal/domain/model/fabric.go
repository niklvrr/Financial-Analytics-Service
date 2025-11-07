package model

import (
	"github.com/niklvrr/Financial-Analytics-Service/pkg/utils"
	"time"
)

type DomainFabric struct {
}

func NewDomainFabric() *DomainFabric {
	return &DomainFabric{}
}

func (f *DomainFabric) BuildBankAccount(id int64, name string, balance float64) (*BankAccount, error) {
	err := utils.ValidateInt64(id)
	if err != nil {
		return nil, err
	}

	err = utils.ValidateString(name)
	if err != nil {
		return nil, err
	}

	err = utils.ValidateFloat(balance)
	if err != nil {
		return nil, err
	}

	return NewBankAccount(id, name, balance), nil
}

func (f *DomainFabric) BuildCategory(id int64, kind, name string) (*Category, error) {
	err := utils.ValidateInt64(id)
	if err != nil {
		return nil, err
	}

	err = utils.ValidateString(name)
	if err != nil {
		return nil, err
	}

	err = utils.ValidateKind(kind)
	if err != nil {
		return nil, err
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

	err := utils.ValidateInt64(id)
	if err != nil {
		return nil, err
	}

	err = utils.ValidateKind(kind)
	if err != nil {
		return nil, err
	}

	err = utils.ValidateInt64(id)
	if err != nil {
		return nil, err
	}

	err = utils.ValidateFloat(amount)
	if err != nil {
		return nil, err
	}

	err = utils.ValidateInt64(id)
	if err != nil {
		return nil, err
	}

	return NewOperation(id, kind, bankAccountId, amount, date, description, categoryId), nil
}
