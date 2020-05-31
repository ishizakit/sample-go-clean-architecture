package main

import (
	"os"

	"github.com/gin-gonic/gin"

	controller "github.com/ishizakit/sample-go-clean-architecture/app/adapter/controller"
	dbAccess "github.com/ishizakit/sample-go-clean-architecture/app/adapter/dataaccess/database"
	mockAccess "github.com/ishizakit/sample-go-clean-architecture/app/adapter/dataaccess/mock"
	infra "github.com/ishizakit/sample-go-clean-architecture/app/infrastructure"
	repository "github.com/ishizakit/sample-go-clean-architecture/app/usecase/dataaccess-interface"
	interactor "github.com/ishizakit/sample-go-clean-architecture/app/usecase/interactor"
)

func main() {
	exec()
}

func exec() {
	// dataaccess生成
	var userDataAccess repository.User
	if os.Getenv("APP_ENV") != "test" {
		// infrastructures生成
		db, err := infra.NewInMemorySQLiteDB()
		if err != nil {
			panic(err)
		}
		userDataAccess = dbAccess.NewUser(db)
	} else {
		userDataAccess = mockAccess.NewUser(mockAccess.UserData{})
	}

	// interactor生成
	createUserinteractor := interactor.NewCreateUser(userDataAccess)
	getUserinteractor := interactor.NewGetUser(userDataAccess)

	// controller生成
	userController := controller.NewUserGINAPI(
		createUserinteractor,
		getUserinteractor,
	)

	router := gin.Default()
	router.POST("/users", userController.Create)
	router.GET("/users/:id", userController.Get)

	router.Run()
}
