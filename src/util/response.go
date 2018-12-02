package util

import (
	"encoding/json"

	"github.com/aws/aws-lambda-go/events"
	"github.com/ludw1gj/mathfever/src/mathematics"
)

type Response events.APIGatewayProxyResponse

type CalculationResponse struct {
	Content string `json:"answer"`
}

type ErrorJSONResponse struct {
	Error string `json:"error"`
}

func HandleAPI(request events.APIGatewayProxyRequest, input mathematics.Mathematics) (Response, error) {
	payload, err := input.ExecuteMath(request.Body)
	if err != nil {
		errResp, _ := json.Marshal(ErrorJSONResponse{Error: err.Error()})
		return Response{StatusCode: 400, Body: string(errResp)}, err
	}

	respBody, _ := json.Marshal(CalculationResponse{Content: payload})
	resp := Response{
		StatusCode: 200,
		Body:       string(respBody),
		Headers: map[string]string{
			"Content-Type":                     "application/json",
			"Access-Control-Allow-Origin":      "*",
			"Access-Control-Allow-Credentials": "true",
		},
	}
	return resp, nil
}
