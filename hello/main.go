package main

import (
	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
)

type response events.APIGatewayProxyResponse

func handler(req events.APIGatewayProxyRequest) (response, error) {
	id, err := principalFromRequest(req)
	if err != nil {
		return response{}, err
	}

	return response{
		StatusCode:      200,
		IsBase64Encoded: false,
		Body:            id,
	}, nil
}

func main() {
	lambda.Start(handler)
}
