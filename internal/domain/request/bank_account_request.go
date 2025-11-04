package request

// Модель запроса для получения счета
type GetBankAccountsRequest struct {
	Id int64
}

// Модель запросы для создания счета
type CreateBankAccountRequest struct {
	Name    string
	Balance float64
}

// Модель запроса для изменения счета
type UpdateBankAccountRequest struct {
	Id      int64
	Name    string
	Balance float64
}

// Модель запроса для удаления счета
type DeleteBankAccountRequest struct {
	Id int64
}
