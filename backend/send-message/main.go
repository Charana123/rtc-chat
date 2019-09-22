package main

import (
	"context"
	"encoding/json"

	utils "github.com/Charana123/rtc-chat/backend/utils"
	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
)

var (
	snsClient *utils.SNSClient
)

func init() {
	snsClient = utils.NewSNSClient()
}

type SendMessage struct {
	Message string `json:"Message"`
}

func Handler(ctx context.Context, req events.APIGatewayWebsocketProxyRequest) (events.APIGatewayProxyResponse, error) {

	data := &SendMessage{}
	err := json.Unmarshal([]byte(req.Body), &data)
	if err != nil {
		return events.APIGatewayProxyResponse{StatusCode: 400}, err
	}
	// check if channel already exists
	err = snsClient.SendMessage(data.Message)
	if err != nil {
		return events.APIGatewayProxyResponse{StatusCode: 400}, err
	}
	return events.APIGatewayProxyResponse{StatusCode: 200}, nil
}

func main() {
	lambda.Start(Handler)
}
