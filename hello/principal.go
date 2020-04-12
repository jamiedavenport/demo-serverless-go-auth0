package main

import (
	"errors"

	"github.com/aws/aws-lambda-go/events"
)

func principalFromRequest(req events.APIGatewayProxyRequest) (string, error) {
	principalIDRaw, exists := req.RequestContext.Authorizer["principalId"]
	if !exists {
		return "", errors.New("missing principal id")
	}

	principalID, ok := principalIDRaw.(string)
	if !ok {
		return "", errors.New("principal id was not a string")
	}

	return principalID, nil
}

