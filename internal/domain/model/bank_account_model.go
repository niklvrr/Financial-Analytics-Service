package model

type BankAccount struct {
	id      int64
	name    string
	balance float64
}

func NewBankAccount(id int64, name string, balance float64) *BankAccount {
	return &BankAccount{
		id:      id,
		name:    name,
		balance: balance,
	}
}
