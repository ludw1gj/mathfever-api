package util

import (
	"encoding/json"
	"fmt"
	"log"

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
		return HandleError(err, "ExecuteMath fault", 400)
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

func HandleError(err error, msg string, code int) (Response, error) {
	log.Printf("message: %v\nerror: %v\n", msg, err.Error())

	newErr := fmt.Errorf("%v, %v", msg, err.Error())
	jsonErrResp, _ := json.Marshal(ErrorJSONResponse{Error: newErr.Error()})
	return Response{StatusCode: code, Body: string(jsonErrResp)}, nil
}
