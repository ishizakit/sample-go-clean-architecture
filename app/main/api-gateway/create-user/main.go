package main

import (
	"context"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"

	controller "github.com/ishizakit/sample-go-clean-architecture/app/adapter/controller"
	dbAccess "github.com/ishizakit/sample-go-clean-architecture/app/adapter/dataaccess/database"
	vm "github.com/ishizakit/sample-go-clean-architecture/app/adapter/view-model"
	infra "github.com/ishizakit/sample-go-clean-architecture/app/infrastructure"
	interactor "github.com/ishizakit/sample-go-clean-architecture/app/usecase/interactor"
)

func main() {
	lambda.Start(exec)
}

func exec(ctx context.Context, request events.APIGatewayProxyRequest) (*vm.CreateUserAPIGateway, error) {
	// infrastructures生成
	db, err := infra.NewInMemorySQLiteDB()
	if err != nil {
		return nil, err
	}

	// dataaccess生成
	userDataAccess := dbAccess.NewUser(db)

	// interactor生成
	createUserinteractor := interactor.NewCreateUser(userDataAccess)

	// controller生成
	userController := controller.NewUserAPIGateway(createUserinteractor, nil)

	// 実行
	return userController.Create(ctx, request)
}
