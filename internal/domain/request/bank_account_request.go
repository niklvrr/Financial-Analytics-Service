package request

// Модель запроса для получения счета
type GetBankAccountsRequest struct {
	Id int64 `csv:"id" json:"id" yaml:"id"`
}

// Модель запросы для создания счета
type CreateBankAccountRequest struct {
	Name    string  `csv:"name" json:"name" yaml:"name"`
	Balance float64 `csv:"balance" json:"balance" yaml:"balance"`
}

// Модель запроса для изменения счета
type UpdateBankAccountRequest struct {
	Id      int64   `csv:"id" json:"id" yaml:"id"`
	Name    string  `csv:"name" json:"name" yaml:"name"`
	Balance float64 `csv:"balance" json:"balance" yaml:"balance"`
}

// Модель запроса для удаления счета
type DeleteBankAccountRequest struct {
	Id int64 `csv:"id" json:"id" yaml:"id"`
}
