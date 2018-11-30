package main

import (
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/aws/aws-lambda-go/events"
	"github.com/ludw1gj/mathfever/mathematics"
	"github.com/ludw1gj/mathfever/util"
)

func Handler(request events.APIGatewayProxyRequest) (util.Response, error) {
	return util.HandleAPI(request, mathematics.HexadecimalToBinaryAPI{})
}

func main() {
	lambda.Start(Handler)
}
