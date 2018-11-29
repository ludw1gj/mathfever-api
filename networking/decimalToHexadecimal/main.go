package main

import (
	"encoding/json"

	"github.com/ludw1gj/mathfever/math"

	"github.com/ludw1gj/mathfever/util"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
)

type DecimalToHexadecimalAPI struct {
	Decimal int `json:"decimal" name:"Decimal"`
}

func Handler(request events.APIGatewayProxyRequest) (util.Response, error) {
	var input DecimalToHexadecimalAPI

	if err := json.Unmarshal([]byte(request.Body), &input); err != nil {
		return util.Response{StatusCode: 400}, nil
	}

	output, err := math.DecimalToHexadecimal(input.Decimal)
	if err != nil {
		return util.Response{StatusCode: 400}, err
	}

	respBody, _ := json.Marshal(util.JSONResponse{output})
	resp := util.Response{
		StatusCode:      200,
		IsBase64Encoded: false,
		Body:            string(respBody),
		Headers: map[string]string{
			"Content-Type": "application/json",
		},
	}
	return resp, nil
}

func main() {
	lambda.Start(Handler)
}
