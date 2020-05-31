package interactor

import (
	dai "github.com/ishizakit/sample-go-clean-architecture/app/usecase/dataaccess-interface"
	ii "github.com/ishizakit/sample-go-clean-architecture/app/usecase/interactor-interface"
	ds "github.com/ishizakit/sample-go-clean-architecture/app/usecase/io-structure"
)

type getUserMock struct {
	user dai.User
}

func NewGetUserMock() ii.GetUser {
	return &getUserMock{}
}

func (g *getUserMock) GetUser(_ *ds.GetUserInputData) (*ds.GetUserOutputData, error) {
	return &ds.GetUserOutputData{
		ID:    1,
		Email: "mock@example.com",
		Name:  "mock",
	}, nil
}
