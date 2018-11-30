package util

import (
	"encoding/json"

	"github.com/ludw1gj/mathfever/mathematics"

	"github.com/aws/aws-lambda-go/events"
)

// Response is of type APIGatewayProxyResponse since we're leveraging the
// AWS Lambda Proxy Request functionality (default behavior)
//
// https://serverless.com/framework/docs/providers/aws/events/apigateway/#lambda-proxy-integration
type Response events.APIGatewayProxyResponse

type JSONResponse struct {
	Content string `json:"content"`
}

func HandleAPI(request events.APIGatewayProxyRequest, input mathematics.Mathematics) (Response, error) {
	if err := json.Unmarshal([]byte(request.Body), &input); err != nil {
		return Response{StatusCode: 400}, nil
	}

	payload, err := input.ExecuteMath()
	if err != nil {
		return Response{StatusCode: 400}, err
	}

	respBody, _ := json.Marshal(JSONResponse{payload})
	resp := Response{
		StatusCode:      200,
		IsBase64Encoded: false,
		Body:            string(respBody),
		Headers: map[string]string{
			"Content-Type": "application/json",
		},
	}
	return resp, nil
}
