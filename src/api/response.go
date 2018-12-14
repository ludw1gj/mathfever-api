package api

import (
	"encoding/json"
	"fmt"
	"log"

	"github.com/aws/aws-lambda-go/events"
)

type Response events.APIGatewayProxyResponse

func HandleMathAPI(request events.APIGatewayProxyRequest, input Mathematics) (Response,
	error) {
	type CalculationResponse struct {
		Content string `json:"answer"`
	}

	payload, err := input.ExecuteMath(request.Body)
	if err != nil {
		return HandleErrorResponse(err, "ExecuteMath fault", 400)
	}

	respBody, err := json.Marshal(CalculationResponse{Content: payload})
	if err != nil {
		return HandleErrorResponse(err, "failed to marshal payload", 500)
	}
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
	jsonErrResp, err := json.Marshal(ErrorJSONResponse{Error: newErr.Error()})
	if err != nil {
		resp := Response{
			IsBase64Encoded: true,
			StatusCode:      500,
			Body:            "failed to marshal error response",
		}
		return resp, nil
	}

	resp := Response{
		IsBase64Encoded: true,
		StatusCode:      code,
		Body:            string(jsonErrResp),
	}
	return resp, nil
}
