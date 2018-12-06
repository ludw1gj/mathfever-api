package util

import (
	"encoding/json"
	"fmt"
	"log"

	"github.com/aws/aws-lambda-go/events"
	"github.com/ludw1gj/mathfever/src/mathematics"
)

type Response events.APIGatewayProxyResponse

func HandleMathAPI(request events.APIGatewayProxyRequest, input mathematics.Mathematics) (Response,
	error) {
	type CalculationResponse struct {
		Content string `json:"answer"`
	}

	payload, err := input.ExecuteMath(request.Body)
	if err != nil {
		return HandleErrorResponse(err, "ExecuteMath fault", 400)
	}

	respBody, _ := json.Marshal(CalculationResponse{Content: payload})
	return HandleResponse(string(respBody))
}

func HandleResponse(body string) (Response, error) {
	resp := Response{
		IsBase64Encoded: true,
		StatusCode:      200,
		Body:            body,
		Headers: map[string]string{
			"Content-Type":                     "application/json",
			"Access-Control-Allow-Origin":      "*",
			"Access-Control-Allow-Credentials": "true",
		},
	}
	return resp, nil
}

func HandleErrorResponse(err error, msg string, code int) (Response, error) {
	log.Printf("message: %v\nerror: %v\n", msg, err.Error())

	type ErrorJSONResponse struct {
		Error string `json:"error"`
	}

	newErr := fmt.Errorf("%v, %v", msg, err.Error())
	jsonErrResp, _ := json.Marshal(ErrorJSONResponse{Error: newErr.Error()})
	return Response{
		IsBase64Encoded: true,
		StatusCode:      code,
		Body:            string(jsonErrResp),
	}, nil
}
