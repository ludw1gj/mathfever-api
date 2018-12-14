package main

import (
	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/ludw1gj/mathfever/src/api"
)

func Handler(request events.APIGatewayProxyRequest) (api.Response, error) {
	return api.HandleMathAPI(request, api.TsaPythagoreanTheoremAPI{})
}

func main() {
	lambda.Start(Handler)
}
