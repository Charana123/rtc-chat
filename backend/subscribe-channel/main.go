package main

import (
	"context"
	"encoding/json"
	"fmt"

	utils "github.com/Charana123/rtc-chat/backend/utils"
	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
)

var (
	ddbClient *utils.DDBClient
)

func init() {
	ddbClient = utils.NewDDBClient()
}

type SubscribeChannel struct {
	ChannelID string `json:"ChannelID"`
}

func Handler(ctx context.Context, req events.APIGatewayWebsocketProxyRequest) (events.APIGatewayProxyResponse, error) {

	data := &SubscribeChannel{}
	err := json.Unmarshal([]byte(req.Body), &data)
	fmt.Println(data)
	if err != nil {
		return events.APIGatewayProxyResponse{StatusCode: 400}, err
	}
	// check if channel already exists
	err = ddbClient.SubscribeUserToChannel(req.RequestContext.ConnectionID, data.ChannelID)
	if err != nil {
		return events.APIGatewayProxyResponse{StatusCode: 400}, err
	}
	return events.APIGatewayProxyResponse{StatusCode: 200}, nil
}

func main() {
	lambda.Start(Handler)
}
