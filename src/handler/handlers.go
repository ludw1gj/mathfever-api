package handler

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/aws/aws-lambda-go/events"
	"github.com/ludw1gj/mathfever-api/src/logic"
	"strings"
)

type CalculationResponse struct {
	Content string `json:"answer"`
}

type ErrorJSONResponse struct {
	Error string `json:"error"`
}

func CategoriesHandler(_ *events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	return handleResponse(logic.GetCategoriesJson())
}

func CalculationHandler(request *events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	calculationSlug := strings.Replace(request.Path, "/calculation/", "", 1)
	calculation, err := logic.FindMath(calculationSlug)
	if err != nil {
		return handleErrorResponse(err, "invalid slug", 404)
	}

	payload, err := calculation.ExecuteMath(request.Body)
	if err != nil {
		return handleErrorResponse(err, "ExecuteMath fault", 400)
	}

	respBody, err := json.Marshal(CalculationResponse{Content: payload})
	if err != nil {
		return handleErrorResponse(err, "failed to marshal payload", 500)
	}
	return handleResponse(string(respBody))
}

func NotFoundHandler(_ *events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	err := errors.New("invalid path")
	return handleErrorResponse(err, "please request valid route", 404)
}

func handleResponse(body string) (events.APIGatewayProxyResponse, error) {
	resp := events.APIGatewayProxyResponse{
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

func handleErrorResponse(err error, msg string, code int) (events.APIGatewayProxyResponse, error) {
	newErr := fmt.Errorf("%s - %v", msg, err.Error())
	jsonErrResp, err := json.Marshal(ErrorJSONResponse{Error: newErr.Error()})
	if err != nil {
		resp := events.APIGatewayProxyResponse{
			IsBase64Encoded: true,
			StatusCode:      500,
			Body:            `{"error": "Failed to marshal error response on server."}`,
			Headers: map[string]string{
				"Content-Type":                     "application/json",
				"Access-Control-Allow-Origin":      "*",
				"Access-Control-Allow-Credentials": "true",
			},
		}
		return resp, nil
	}

	resp := events.APIGatewayProxyResponse{
		IsBase64Encoded: true,
		StatusCode:      code,
		Body:            string(jsonErrResp),
		Headers: map[string]string{
			"Content-Type":                     "application/json",
			"Access-Control-Allow-Origin":      "*",
			"Access-Control-Allow-Credentials": "true",
		},
	}
	return resp, nil
}
