package model

import "time"

type Operation struct {
	id            int64
	kind          string
	bankAccountId int64
	amount        float64
	date          time.Time
	description   string
	categoryId    int64
}

func NewOperation(
	id int64,
	kind string,
	bankAccountId int64,
	amount float64,
	date time.Time,
	description string,
	categoryId int64,
) *Operation {
	return &Operation{
		id:            id,
		kind:          kind,
		bankAccountId: bankAccountId,
		amount:        amount,
		date:          date,
		description:   description,
		categoryId:    categoryId,
	}
}

func (o *Operation) ID() int64 {
	return o.id
}

func (o *Operation) SetID(id int64) {
	o.id = id
}

func (o *Operation) Kind() string {
	return o.kind
}

func (o *Operation) SetKind(kind string) {
	o.kind = kind
}

func (o *Operation) BankAccountId() int64 {
	return o.bankAccountId
}

func (o *Operation) SetBankAccountId(bankAccountId int64) {
	o.bankAccountId = bankAccountId
}

func (o *Operation) Amount() float64 {
	return o.amount
}

func (o *Operation) SetAmount(amount float64) {
	o.amount = amount
}

func (o *Operation) Date() time.Time {
	return o.date
}

func (o *Operation) SetDate(date time.Time) {
	o.date = date
}

func (o *Operation) Description() string {
	return o.description
}

func (o *Operation) SetDescription(description string) {
	o.description = description
}

func (o *Operation) CategoryId() int64 {
	return o.categoryId
}

func (o *Operation) SetCategoryId(categoryId int64) {
	o.categoryId = categoryId
}
