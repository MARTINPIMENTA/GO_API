package main

import (
	lambdaFuncs "github.com/MARTINPIMENTA/pimen_rest_api_go/src/api/domain/lambda"
	handlerRead "github.com/MARTINPIMENTA/pimen_rest_api_go/src/api/handler/reader"
	handlerWrite "github.com/MARTINPIMENTA/pimen_rest_api_go/src/api/handler/writer"

	"github.com/aws/aws-lambda-go/lambda"
)

// main func.
func main() {
	handlerRead.HandleReadRequests()
	handlerWrite.HandleWriteRequests()

	// Initialization for lambda functions.
	lambda.Start(lambdaFuncs.HandlerLambda)
}
