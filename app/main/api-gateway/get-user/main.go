package main

import (
	"context"
	"os"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"

	controller "github.com/ishizakit/sample-go-clean-architecture/app/adapter/controller"
	dbAccess "github.com/ishizakit/sample-go-clean-architecture/app/adapter/dataaccess/database"
	mockAccess "github.com/ishizakit/sample-go-clean-architecture/app/adapter/dataaccess/mock"
	vm "github.com/ishizakit/sample-go-clean-architecture/app/adapter/view-model"
	infra "github.com/ishizakit/sample-go-clean-architecture/app/infrastructure"
	repository "github.com/ishizakit/sample-go-clean-architecture/app/usecase/dataaccess-interface"
	interactor "github.com/ishizakit/sample-go-clean-architecture/app/usecase/interactor"
)

func main() {
	lambda.Start(exec)
}

func exec(ctx context.Context, request events.APIGatewayProxyRequest) (*vm.CreateUserAPIGateway, error) {
	// dataaccess生成
	var userDataAccess repository.User
	if os.Getenv("APP_ENV") != "test" {
		// infrastructures生成
		mysql, err := infra.NewInMemorySQLiteDB()
		if err != nil {
			return nil, err
		}
		userDataAccess = dbAccess.NewUser(mysql)
	} else {
		userDataAccess = mockAccess.NewUser(mockAccess.UserData{})
	}

	// interactor生成
	getUserinteractor := interactor.NewGetUser(userDataAccess)

	// controller生成
	userController := controller.NewUserAPIGateway(nil, getUserinteractor)

	// 実行
	return userController.Create(ctx, request)
}
