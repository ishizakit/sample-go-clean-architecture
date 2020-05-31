package interactor

import (
	"github.com/ishizakit/sample-go-clean-architecture/app/entity"
	dai "github.com/ishizakit/sample-go-clean-architecture/app/usecase/dataaccess-interface"
	ii "github.com/ishizakit/sample-go-clean-architecture/app/usecase/interactor-interface"
	ds "github.com/ishizakit/sample-go-clean-architecture/app/usecase/io-structure"
)

type createUser struct {
	user dai.User
}

func NewCreateUser(user dai.User) ii.CreateUser {
	return &createUser{
		user: user,
	}
}

func (c *createUser) CreateUser(input *ds.CreateUserInputData) (*ds.CreateUserOutputData, error) {
	model := entity.NewUser(input.Email, input.Name)
	user, err := c.user.Create(model)
	if err != nil {
		return nil, err
	}

	return &ds.CreateUserOutputData{
		ID:    user.ID,
		Email: user.Email,
		Name:  user.Name,
	}, nil
}
