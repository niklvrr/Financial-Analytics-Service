package request

// Модель запроса для получения операции
type GetOperationRequest struct {
	Id int64
}

func NewGetOperationRequest(id int64) *GetOperationRequest {
	return &GetOperationRequest{
		Id: id,
	}
}

// Модель для создания операции
type CreateOperationRequest struct {
	Kind          string
	BankAccountId int64
	Amount        float64
	Description   string
	CategoryId    int64
}

func NewCreateOperationRequest(
	kind string,
	bankAccountId int64,
	amount float64,
	description string,
	categoryId int64,
) *CreateOperationRequest {
	return &CreateOperationRequest{
		Kind:          kind,
		BankAccountId: bankAccountId,
		Amount:        amount,
		Description:   description,
		CategoryId:    categoryId,
	}
}

// Модель для изменения операции
type UpdateOperationRequest struct {
	Id            int64
	Kind          string
	BankAccountId int64
	Amount        float64
	Description   string
	CategoryId    int64
}

func NewUpdateOperationRequest(
	id int64,
	kind string,
	bankAccountId int64,
	amount float64,
	description string,
	categoryId int64,
) *UpdateOperationRequest {
	return &UpdateOperationRequest{
		Id:            id,
		Kind:          kind,
		BankAccountId: bankAccountId,
		Amount:        amount,
		Description:   description,
		CategoryId:    categoryId,
	}
}

// Модель для удаления операци
type DeleteOperationRequest struct {
	Id int64
}

func NewDeleteOperationRequest(id int64) *DeleteOperationRequest {
	return &DeleteOperationRequest{
		Id: id,
	}
}
