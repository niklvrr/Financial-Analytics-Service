package request

// Модель запроса для получения счета
type GetBankAccountsRequest struct {
	Id int64
}

func NewGetBankAccountsRequest(id int64) *GetBankAccountsRequest {
	return &GetBankAccountsRequest{
		Id: id,
	}
}

// Модель запросы для создания счета
type CreateBankAccountRequest struct {
	Name string
}

func NewCreateBankAccountRequest(name string) *CreateBankAccountRequest {
	return &CreateBankAccountRequest{
		Name: name,
	}
}

// Модель запроса для изменения счета
type UpdateBankAccountRequest struct {
	Id   int64
	Name string
}

func NewUpdateBankAccountRequest(id int64, name string) *UpdateBankAccountRequest {
	return &UpdateBankAccountRequest{
		Id:   id,
		Name: name,
	}
}

// Модель запроса для удаления счета
type DeleteBankAccountRequest struct {
	Id int64
}

func NewDeleteBankAccountRequest(id int64) *DeleteBankAccountRequest {
	return &DeleteBankAccountRequest{
		Id: id,
	}
}
