package main

import (
	"errors"
	"strings"

	"github.com/dgrijalva/jwt-go"
	"github.com/jamiedavenport/go-auth0-jwt"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
)

type request = events.APIGatewayCustomAuthorizerRequest
type response = events.APIGatewayCustomAuthorizerResponse

func tokenFromRequest(req request) (string, error) {
	bearerToken := req.AuthorizationToken
	tokenParts := strings.Split(bearerToken, " ")
	if len(tokenParts) < 2 {
		return "", errors.New("invalid token format")
	}

	return tokenParts[1], nil
}

func handler(req request) (response, error) {
	tokenString, err := tokenFromRequest(req) // 1
	if err != nil {
		return response{}, err
	}

	a := auth0.Parser{ // 2
		Audience: "https://demo-serverless-go-auth0.jamiedavenport.dev",
		Domain:   "https://demo-serverless-go-auth0.eu.auth0.com/",
	}
	token, err := a.Parse(tokenString)
	if err != nil {
		return response{}, err
	}

	principalID, ok := token.Claims.(jwt.MapClaims)["sub"].(string) // 3
	if !ok {
		return response{}, errors.New("invalid token")
	}

	policy := generatePolicy("Allow") // 4

	return response{ // 5
		PrincipalID:    principalID,
		PolicyDocument: policy,
	}, nil
}

func main() {
	lambda.Start(handler)
}
