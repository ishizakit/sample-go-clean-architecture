package controller

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"

	vm "github.com/ishizakit/sample-go-clean-architecture/app/adapter/view-model"
	ii "github.com/ishizakit/sample-go-clean-architecture/app/usecase/interactor-interface"
	ds "github.com/ishizakit/sample-go-clean-architecture/app/usecase/io-structure"
)

type UserGINAPI struct {
	create ii.CreateUser
	get    ii.GetUser
}

func NewUserGINAPI(create ii.CreateUser, get ii.GetUser) *UserGINAPI {
	return &UserGINAPI{
		create: create,
		get:    get,
	}
}

func (u *UserGINAPI) Create(ctx *gin.Context) {
	type inputBody struct {
		Name  string `json:"name"`
		Email string `json:"email"`
	}
	var body inputBody
	if err := ctx.Bind(&body); err != nil {
		fmt.Println(err)
		// TODO: エラーハンドリング
		return
	}

	input := ds.NewCreateUserInputData(body.Email, body.Email)
	output, err := u.create.CreateUser(input)
	u.createResponse(ctx, output, err)
}

// createResponse Createメソッドのレスポンス生成 Presenterに相当
func (u *UserGINAPI) createResponse(ctx *gin.Context, output *ds.CreateUserOutputData, err error) {
	errResponseBody := `{"message":"internal server error."}`

	if err != nil {
		fmt.Println(err)
		ctx.JSON(http.StatusInternalServerError, errResponseBody)
		return
	}

	ctx.JSON(http.StatusOK, vm.CreateUserGIN{
		ID:    output.ID,
		Email: output.Email,
		Name:  output.Name,
	})
}

func (u *UserGINAPI) Get(ctx *gin.Context) {
	id, err := strconv.ParseInt(ctx.Param("id"), 10, 64)
	if err != nil {
		// TODO: エラーハンドリング
	}
	input := ds.NewGetUserInputData(id)
	output, err := u.get.GetUser(input)
	u.getResponse(ctx, output, err)
}

// getResponse Getメソッドのレスポンス生成 Presenterに相当
func (u *UserGINAPI) getResponse(ctx *gin.Context, output *ds.GetUserOutputData, err error) {
	if err != nil {
		fmt.Println(err)
		ctx.JSON(http.StatusInternalServerError, `{"message":"internal server error."}`)
		return
	}

	if output == nil {
		fmt.Println(err)
		ctx.JSON(http.StatusNotFound, `{"message":"not found."}`)
		return
	}

	ctx.JSON(http.StatusOK, vm.CreateUserGIN{
		ID:    output.ID,
		Email: output.Email,
		Name:  output.Name,
	})
}
