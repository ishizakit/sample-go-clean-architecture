package user

import (
	ds "github.com/ishizakit/sample-go-clean-architecture/app/usecase/io-structure"
)

type CreateUser interface {
	CreateUser(input *ds.CreateUserInputData) (*ds.CreateUserOutputData, error)
}
