package response

import "time"

// Модель для ответа для счета
type BankAccountResponse struct {
	Id      int64
	Name    string
	Balance float64
}

func NewBankAccountResponse(
	id int64,
	name string,
	balance float64,
) *BankAccountResponse {
	return &BankAccountResponse{
		Id:      id,
		Name:    name,
		Balance: balance,
	}
}

type CategoryResponse struct {
	Id   int64
	Kind string
	Name string
}

func NewCategoryResponse(
	id int64,
	kind string,
	name string,
) *CategoryResponse {
	return &CategoryResponse{
		Id:   id,
		Kind: kind,
		Name: name,
	}
}

type OperationResponse struct {
	Id            int64
	Kind          string
	BankAccountId int64
	Amount        float64
	Date          time.Time
	Description   string
	CategoryId    int64
}

func NewOperationResponse(
	id int64,
	kind string,
	bankAccountId int64,
	amount float64,
	date time.Time,
	description string,
	categoryId int64,
) *OperationResponse {
	return &OperationResponse{
		Id:            id,
		Kind:          kind,
		BankAccountId: bankAccountId,
		Amount:        amount,
		Date:          date,
		Description:   description,
		CategoryId:    categoryId,
	}
}
