package mock

import (
	"reflect"

	"github.com/ishizakit/sample-go-clean-architecture/app/entity"
	repository "github.com/ishizakit/sample-go-clean-architecture/app/usecase/dataaccess-interface"
)

type userMock struct {
	data UserData
}

// UserData 各メソッドの期待する引数と戻り値のデータ
type UserData struct {
	Create []struct {
		Input struct {
			User *entity.User
		}
		Output struct {
			User *entity.User
			Err  error
		}
	}
	Get []struct {
		Input struct {
			ID int64
		}
		Output struct {
			User *entity.User
			Err  error
		}
	}
}

func NewUser(userData UserData) repository.User {
	return &userMock{
		data: userData,
	}
}

func (u *userMock) Create(user *entity.User) (*entity.User, error) {
	// dataが設定されていれば、inputに対応するoutputを返す
	if u.data.Create != nil {
		for _, data := range u.data.Create {
			if reflect.DeepEqual(data.Input.User, user) {
				return data.Output.User, data.Output.Err
			}
		}
	}

	// dataが設定されていない、またはinputに対応するoutputがない
	return &entity.User{
		ID:    1,
		Name:  "mock",
		Email: "mock@example.com",
	}, nil
}

func (u *userMock) Get(id int64) (*entity.User, error) {
	// dataが設定されていれば、inputに対応するoutputを返す
	if u.data.Get != nil {
		for _, data := range u.data.Get {
			if reflect.DeepEqual(data.Input.ID, id) {
				return data.Output.User, data.Output.Err
			}
		}
	}

	// dataが設定されていない、またはinputに対応するoutputがない
	return &entity.User{
		ID:    id,
		Name:  "mock",
		Email: "mock@example.com",
	}, nil
}
