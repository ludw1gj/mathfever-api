package main

import (
	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"

	"github.com/ludw1gj/mathfever/mathematics"
	"github.com/ludw1gj/mathfever/util"
)

func Handler(request events.APIGatewayProxyRequest) (util.Response, error) {
	return util.HandleAPI(request, mathematics.PercentageChangeAPI{})
}

func main() {
	lambda.Start(Handler)
}
