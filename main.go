package main

import (
	"log"
	"net/http"

	lambdaFuncs "github.com/MARTINPIMENTA/pimen_rest_api_go/src/api/domain/lambda"
	handlerRead "github.com/MARTINPIMENTA/pimen_rest_api_go/src/api/handler/reader"
	handlerWrite "github.com/MARTINPIMENTA/pimen_rest_api_go/src/api/handler/writer"
	"github.com/gorilla/mux"

	"github.com/aws/aws-lambda-go/lambda"
)

// main func.
func main() {
	myRouter := mux.NewRouter().StrictSlash(true)

	handlerRead.HandleReadRequests(myRouter)
	handlerWrite.HandleWriteRequests(myRouter)
	log.Fatal(http.ListenAndServe(":8081", myRouter))

	// Initialization for lambda functions.
	lambda.Start(lambdaFuncs.HandlerLambda)
}
