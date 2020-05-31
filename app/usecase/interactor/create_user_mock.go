package interactor

import (
	dai "github.com/ishizakit/sample-go-clean-architecture/app/usecase/dataaccess-interface"
	ii "github.com/ishizakit/sample-go-clean-architecture/app/usecase/interactor-interface"
	ds "github.com/ishizakit/sample-go-clean-architecture/app/usecase/io-structure"
)

type createUserMock struct {
	user dai.User
}

func NewCreateUserMock() ii.CreateUser {
	return &createUserMock{}
}

func (c *createUserMock) CreateUser(_ *ds.CreateUserInputData) (*ds.CreateUserOutputData, error) {
	return &ds.CreateUserOutputData{
		ID:    1,
		Email: "mock@example.com",
		Name:  "mock",
	}, nil
}
