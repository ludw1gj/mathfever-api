package main

import (
	"github.com/ludw1gj/mathfever/util"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/ludw1gj/mathfever/mathematics"
)

func Handler(request events.APIGatewayProxyRequest) (util.Response, error) {
	return util.HandleAPI(request, mathematics.DecimalToBinaryAPI{})
}

func main() {
	lambda.Start(Handler)
}
