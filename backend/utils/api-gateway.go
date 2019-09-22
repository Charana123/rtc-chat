package utils

import (
	"bytes"
	"encoding/json"
	"fmt"

	"github.com/aws/aws-sdk-go/aws"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/apigatewaymanagementapi"
)

type APIGateway struct {
	svc *apigatewaymanagementapi.ApiGatewayManagementApi
}

func NewAPIGateway() *APIGateway {
	sess := session.Must(session.NewSessionWithOptions(
		session.Options{
			SharedConfigState: session.SharedConfigEnable,
		},
	))
	return &APIGateway{
		svc: apigatewaymanagementapi.New(
			sess,
			aws.NewConfig().
				WithEndpoint("https://drhscexf31.execute-api.us-east-1.amazonaws.com/dev"),
		),
	}
}

func (apic *APIGateway) SendMessageToClient(connectionID string, message interface{}) error {
	body, err := json.Marshal(map[string]interface{}{
		"message": message,
	})
	if err != nil {
		return err
	}

	input := &apigatewaymanagementapi.PostToConnectionInput{
		ConnectionId: aws.String(connectionID),
		Data:         body,
	}
	_, err = apic.svc.PostToConnection(input)
	return err
}

func ConstructResponse(statusCode int, bodyMap map[string]interface{}) (events.APIGatewayProxyResponse, error) {
	body, err := json.Marshal(bodyMap)
	fmt.Println("here1")
	if err != nil {
		return events.APIGatewayProxyResponse{StatusCode: 400}, err
	}
	fmt.Println("here2")
	buf := &bytes.Buffer{}
	buf.Write(body)
	// json.HTMLEscape(buf, body)

	return events.APIGatewayProxyResponse{
		StatusCode:      statusCode,
		IsBase64Encoded: false,
		Body:            buf.String(),
		Headers: map[string]string{
			"Content-Type": "application/json",
		},
	}, nil
}
