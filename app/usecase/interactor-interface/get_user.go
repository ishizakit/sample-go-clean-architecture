package user

import (
	ds "github.com/ishizakit/sample-go-clean-architecture/app/usecase/io-structure"
)

type GetUser interface {
	GetUser(input *ds.GetUserInputData) (*ds.GetUserOutputData, error)
}
