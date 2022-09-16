package main

import (
	"AWS_Challenge/controller"
	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/awslabs/aws-lambda-go-api-proxy/core"
	"github.com/awslabs/aws-lambda-go-api-proxy/gorillamux"
	"github.com/gorilla/mux"
)

func main() {
	lambda.Start(Handler)
}

func Handler(request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	router := mux.NewRouter()
	router.HandleFunc("/devices", controller.CreateDevice).Methods("POST")
	router.HandleFunc("/devices/{id}", controller.GetDevice).Methods("GET")
	app, _ := gorillamux.New(router).Proxy(*core.NewSwitchableAPIGatewayRequestV1(&request))
	return *app.Version1(), nil
}
