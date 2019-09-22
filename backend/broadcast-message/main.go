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
	ddbClient        *utils.DDBClient
	apiGatewayClient *utils.APIGateway
)

func init() {
	ddbClient = utils.NewDDBClient()
	apiGatewayClient = utils.NewAPIGateway()
}

type BroadcastMessage struct {
	ChannelID string `json:"ChannelID"`
	Message   string `json:"Message"`
}

func Handler(ctx context.Context, snsEvent events.SNSEvent) error {

	for _, event := range snsEvent.Records {
		// Obtain json channelID and message published to sns topic
		fmt.Println("hello world")
		data := &BroadcastMessage{}
		json.Unmarshal([]byte(event.SNS.Message), data)
		fmt.Println(data)

		// query ddb for connectionIDs subscribed to the channel
		connectionIDs, err := ddbClient.ListSubscribersToChannel(data.ChannelID)
		if err != nil {
			return err
		}
		// send message to each subscribed connection
		for _, connectionID := range connectionIDs {
			fmt.Println(connectionID)
			fmt.Println(data.Message)
			err := apiGatewayClient.SendMessageToClient(connectionID, data.Message)
			if err != nil {
				// awsErr := err.(awserr.Error)
				// switch awsErr.Code() {
				// // Ignore if connection doesn't exist (i.e. user has disconnected)
				// case apigatewaymanagementapi.ErrCodeGoneException:
				// 	continue
				// default:
				// 	return err
				// }
				continue
			}
		}
	}
	return nil
}

func main() {
	lambda.Start(Handler)
}
