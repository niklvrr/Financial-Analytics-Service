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
