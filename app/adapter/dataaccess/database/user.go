package database

import (
	"database/sql"

	"github.com/ishizakit/sample-go-clean-architecture/app/entity"
	repository "github.com/ishizakit/sample-go-clean-architecture/app/usecase/dataaccess-interface"
	"github.com/jmoiron/sqlx"
)

type user struct {
	db *sqlx.DB
}

func NewUser(db *sqlx.DB) repository.User {
	return &user{
		db: db,
	}
}

func (u *user) Create(user *entity.User) (*entity.User, error) {
	result, err := u.db.NamedExec(`
	INSERT INTO
		users (name,email)
	VALUES
		(:name,:email)
	`, *user)
	if err != nil {
		return nil, err
	}

	id, err := result.LastInsertId()
	if err != nil {
		return nil, err
	}

	return &entity.User{
		ID:    id,
		Email: user.Email,
		Name:  user.Name,
	}, nil
}

func (u *user) Get(id int64) (*entity.User, error) {
	var user entity.User
	err := u.db.Get(&user, `
	SELECT
		*
	FROM
		users
	WHERE
		id = ?
	`, id)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return nil, err
	}
	return &user, nil
}
