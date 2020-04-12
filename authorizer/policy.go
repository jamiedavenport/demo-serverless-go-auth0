package main

import "github.com/aws/aws-lambda-go/events"

type policy = events.APIGatewayCustomAuthorizerPolicy

func generatePolicy(effect string) policy {
	if effect != "" {
		return policy{
			Version: "2012-10-17",
			Statement: []events.IAMPolicyStatement{
				{
					Action:   []string{"execute-api:Invoke"},
					Effect:   effect,
					Resource: []string{"arn:aws:execute-api:eu-west-2:249071699039:vmlk9hzwt9/*"},
				},
			},
		}
	}

	return policy{}
}
