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

type CreateChannel struct {
	ChannelID string `json:"ChannelID"`
}

func Handler(ctx context.Context, req events.APIGatewayWebsocketProxyRequest) (events.APIGatewayProxyResponse, error) {

	// decode channelID
	data := &CreateChannel{}
	err := json.Unmarshal([]byte(req.Body), &data)
	fmt.Println(data)
	if err != nil {
		return events.APIGatewayProxyResponse{StatusCode: 400}, err
	}
	// check if channel already exists
	exists, err := ddbClient.CheckChannelExists(data.ChannelID)
	if err != nil {
		return events.APIGatewayProxyResponse{StatusCode: 400}, err
	}
	fmt.Println("check successful")
	// if the channel exists, return an appropriate error message
	if exists {
		return utils.ConstructResponse(400, map[string]interface{}{
			"message": "channel already exists",
		})
	}
	// otherwise create the channel
	fmt.Println("creating channel")
	fmt.Println("req.RequestContext.ConnectionID", req.RequestContext.ConnectionID)
	err = ddbClient.CreateChannel(data.ChannelID, req.RequestContext.ConnectionID)
	if err != nil {
		return events.APIGatewayProxyResponse{StatusCode: 400}, err
	}
	return events.APIGatewayProxyResponse{StatusCode: 200}, nil
}

func main() {
	lambda.Start(Handler)
}
