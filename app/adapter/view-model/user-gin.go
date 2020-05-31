package vm

// CreateUserGIN gin用のUserControllerの戻り値の型
type CreateUserGIN struct {
	ID    int64
	Email string
	Name  string
}

// GetUserGIN gin用のUserControllerの戻り値の型
type GetUserGIN struct {
	ID    int64
	Email string
	Name  string
}
