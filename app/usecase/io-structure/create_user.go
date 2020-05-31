package iostructure

/*
 * UseCaseInteractorのInput用のデータストラクチャ
 */

type CreateUserInputData struct {
	Email string
	Name  string
}

func NewCreateUserInputData(email, name string) *CreateUserInputData {
	return &CreateUserInputData{
		Email: email,
		Name:  name,
	}
}

/*
 * UseCaseInteractorのOutput用のデータストラクチャ
 */

type CreateUserOutputData struct {
	ID    int64
	Email string
	Name  string
}

func NewCreateUserOutputData(id int64, email, name string) *CreateUserOutputData {
	return &CreateUserOutputData{
		ID:    id,
		Email: email,
		Name:  name,
	}
}
