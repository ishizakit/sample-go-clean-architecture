package vm

import (
	"github.com/aws/aws-lambda-go/events"
)

// CreateUserAPIGateway APIGateway用のUserControllerの戻り値の型
type CreateUserAPIGateway events.APIGatewayProxyResponse

type CreateUserBody struct {
	ID    int64  `json:"id"`
	Email string `json:"email"`
	Name  string `json:"name"`
}

// GetUserAPIGateway APIGateway用のUserControllerの戻り値の型
type GetUserAPIGateway events.APIGatewayProxyResponse

type GetUserBody struct {
	ID    int64  `json:"id"`
	Email string `json:"email"`
	Name  string `json:"name"`
}
