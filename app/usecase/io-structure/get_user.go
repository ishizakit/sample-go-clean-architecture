package iostructure

/*
 * UseCaseInteractorのInput用のデータストラクチャ
 */

type GetUserInputData struct {
	ID int64
}

func NewGetUserInputData(id int64) *GetUserInputData {
	return &GetUserInputData{
		ID: id,
	}
}

/*
 * UseCaseInteractorのOutput用のデータストラクチャ
 */

type GetUserOutputData struct {
	ID    int64
	Email string
	Name  string
}

func NewGetUserOutputData(id int64, email, name string) *GetUserOutputData {
	return &GetUserOutputData{
		ID:    id,
		Email: email,
		Name:  name,
	}
}
