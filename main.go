package main

import (
	lambdaFuncs "github.com/MARTINPIMENTA/pimen_rest_api_go/src/api/domain/lambda"
	handler "github.com/MARTINPIMENTA/pimen_rest_api_go/src/api/handler"

	"github.com/aws/aws-lambda-go/lambda"
)

// main func.
func main() {
	handler.HandleRequests()

	// Initialization for lambda functions.
	lambda.Start(lambdaFuncs.HandlerLambda)
}
