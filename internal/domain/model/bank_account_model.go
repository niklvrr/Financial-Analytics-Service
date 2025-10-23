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

func (b *BankAccount) ID() int64 {
	return b.id
}

func (b *BankAccount) Name() string {
	return b.name
}

func (b *BankAccount) SetName(name string) {
	b.name = name
}

func (b *BankAccount) Balance() float64 {
	return b.balance
}

func (b *BankAccount) SetBalance(balance float64) {
	b.balance = balance
}
