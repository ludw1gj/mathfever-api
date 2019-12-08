package main

import (
	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/ludw1gj/mathfever-api/src/handler"
	"github.com/ludw1gj/mathfever-api/src/router"
)

func Handler(request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	r := router.NewRouter(&request, handler.NotFoundHandler)
	r.AddRoute("/categories", handler.CategoriesHandler)
	r.AddRoute("/calculation/{slug}", handler.CalculationHandler)
	return r.Handle(request.Path)
}

func main() {
	lambda.Start(Handler)
}
