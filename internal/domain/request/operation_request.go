package request

import "time"

// Модель запроса для получения операции
type GetOperationRequest struct {
	Id int64
}

// Модель для создания операции
type CreateOperationRequest struct {
	Kind          string
	BankAccountId int64
	Amount        float64
	Description   string
	CategoryId    int64
}

// Модель для изменения операции
type UpdateOperationRequest struct {
	Id            int64
	Kind          string
	BankAccountId int64
	Amount        float64
	Date          time.Time
	Description   string
	CategoryId    int64
}

// Модель для удаления операци
type DeleteOperationRequest struct {
	Id int64
}
