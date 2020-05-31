package interactor

import (
	dai "github.com/ishizakit/sample-go-clean-architecture/app/usecase/dataaccess-interface"
	ii "github.com/ishizakit/sample-go-clean-architecture/app/usecase/interactor-interface"
	ds "github.com/ishizakit/sample-go-clean-architecture/app/usecase/io-structure"
)

type getUser struct {
	user dai.User
}

func NewGetUser(user dai.User) ii.GetUser {
	return &getUser{
		user: user,
	}
}

func (g *getUser) GetUser(input *ds.GetUserInputData) (*ds.GetUserOutputData, error) {
	user, err := g.user.Get(input.ID)
	if err != nil {
		return nil, err
	}

	if user == nil {
		return nil, nil
	}

	return ds.NewGetUserOutputData(user.ID, user.Email, user.Name), nil
}
