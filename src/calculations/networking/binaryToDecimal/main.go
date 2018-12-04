package main

import (
	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/ludw1gj/mathfever/src/mathematics"
	"github.com/ludw1gj/mathfever/src/util"
)

func Handler(request events.APIGatewayProxyRequest) (util.Response, error) {
	return util.HandleMathAPI(request, mathematics.BinaryToDecimalAPI{})
}

func main() {
	lambda.Start(Handler)
}
