package user

import (
	"github.com/ishizakit/sample-go-clean-architecture/app/entity"
)

// User UserのDataAccessのInterface
type User interface {
	Create(user *entity.User) (*entity.User, error)
	Get(id int64) (*entity.User, error)
}
