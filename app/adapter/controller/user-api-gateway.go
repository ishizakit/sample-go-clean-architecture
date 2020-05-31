package controller

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"strconv"

	"github.com/aws/aws-lambda-go/events"

	vm "github.com/ishizakit/sample-go-clean-architecture/app/adapter/view-model"
	ii "github.com/ishizakit/sample-go-clean-architecture/app/usecase/interactor-interface"
	ds "github.com/ishizakit/sample-go-clean-architecture/app/usecase/io-structure"
)

type UserAPIGateway struct {
	create ii.CreateUser
	get    ii.GetUser
}

func NewUserAPIGateway(create ii.CreateUser, get ii.GetUser) *UserAPIGateway {
	return &UserAPIGateway{
		create: create,
		get:    get,
	}
}

func (u *UserAPIGateway) Create(ctx context.Context, request events.APIGatewayProxyRequest) (*vm.CreateUserAPIGateway, error) {
	input := ds.NewCreateUserInputData(request.QueryStringParameters["email"], request.QueryStringParameters["userName"])
	output, err := u.create.CreateUser(input)
	return u.createResponse(output, err), nil
}

// createResponse Createメソッドのレスポンス生成 Presenterに相当
func (u *UserAPIGateway) createResponse(output *ds.CreateUserOutputData, err error) *vm.CreateUserAPIGateway {
	errResponse := &vm.CreateUserAPIGateway{
		StatusCode: http.StatusInternalServerError,
		Body:       `{"message":"internal server error."}`,
	}

	if err != nil {
		return errResponse
	}

	json, err := json.Marshal(vm.CreateUserBody{
		ID:    output.ID,
		Email: output.Email,
		Name:  output.Name,
	})
	if err != nil {
		return errResponse
	}

	return &vm.CreateUserAPIGateway{
		StatusCode: http.StatusOK,
		Body:       string(json),
	}
}

func (u *UserAPIGateway) Get(ctx context.Context, request events.APIGatewayProxyRequest) (*vm.GetUserAPIGateway, error) {
	id, ok := request.PathParameters["id"]
	if !ok {
		err := errors.New("URLからidが取得できませんでした")
		fmt.Println(err)
		// TODO: エラーハンドリング
	}
	idInt, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
		fmt.Println(err)
		// TODO: エラーハンドリング
	}
	input := ds.NewGetUserInputData(idInt)
	output, err := u.get.GetUser(input)
	return u.getResponse(output, err), nil
}

// getResponse Getメソッドのレスポンス生成 Presenterに相当
func (u *UserAPIGateway) getResponse(output *ds.GetUserOutputData, err error) *vm.GetUserAPIGateway {
	errResponse := &vm.GetUserAPIGateway{
		StatusCode: http.StatusInternalServerError,
		Body:       `{"message":"internal server error."}`,
	}

	if err != nil {
		fmt.Println(err)
		return errResponse
	}

	if output == nil {
		return &vm.GetUserAPIGateway{
			StatusCode: http.StatusNotFound,
			Body:       `{"message":"not found."}`,
		}
	}

	json, err := json.Marshal(vm.GetUserBody{
		ID:    output.ID,
		Email: output.Email,
		Name:  output.Name,
	})
	if err != nil {
		fmt.Println(err)
		return errResponse
	}

	fmt.Println(string(json))

	return &vm.GetUserAPIGateway{
		StatusCode: http.StatusOK,
		Body:       string(json),
	}
}
