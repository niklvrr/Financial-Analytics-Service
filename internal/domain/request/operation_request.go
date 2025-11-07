package request

import "time"

// Модель запроса для получения операции
type GetOperationRequest struct {
	Id int64 `csv:"id" json:"id" yaml:"id"`
}

// Модель для создания операции
type CreateOperationRequest struct {
	Kind          string  `csv:"kind" json:"kind" yaml:"kind"`
	BankAccountId int64   `csv:"bank_account_id" json:"bank_account_id" yaml:"bank_account_id"`
	Amount        float64 `csv:"amount" json:"amount" yaml:"amount"`
	Description   string  `csv:"description" json:"description" yaml:"description"`
	CategoryId    int64   `csv:"category_id" json:"category_id" yaml:"category_id"`
}

// Модель для изменения операции
type UpdateOperationRequest struct {
	Id            int64     `csv:"id" json:"id" yaml:"id"`
	Kind          string    `csv:"kind" json:"kind" yaml:"kind"`
	BankAccountId int64     `csv:"bank_account_id" json:"bank_account_id" yaml:"bank_account_id"`
	Amount        float64   `csv:"amount" json:"amount" yaml:"amount"`
	Date          time.Time `csv:"date" json:"date" yaml:"date"`
	Description   string    `csv:"description" json:"description" yaml:"description"`
	CategoryId    int64     `csv:"category_id" json:"category_id" yaml:"category_id"`
}

// Модель для удаления операци
type DeleteOperationRequest struct {
	Id int64 `csv:"id" json:"id" yaml:"id"`
}
