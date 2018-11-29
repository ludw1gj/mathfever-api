package main

import (
	"encoding/json"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"

	"github.com/ludw1gj/mathfever/math"
	"github.com/ludw1gj/mathfever/util"
)

type binaryToDecimalInput struct {
	Binary string `json:"binary"`
}

func Handler(request events.APIGatewayProxyRequest) (util.Response, error) {
	var input binaryToDecimalInput

	if err := json.Unmarshal([]byte(request.Body), &input); err != nil {
		return util.Response{StatusCode: 400}, nil
	}

	output, err := math.BinaryToDecimal(input.Binary)
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
