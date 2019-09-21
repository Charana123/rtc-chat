package utils

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/dynamodb"
)

type DDBClient struct {
	svc       *dynamodb.DynamoDB
	TableName string
	GSIName   string
}

func NewDDBClient(svc *dynamodb.DynamoDB) *DDBClient {
	return &DDBClient{
		svc:       svc,
		TableName: "RTCChat",
		GSIName:   "reverseGSI",
	}
}

// CreateChannel creates two items
// {pk: channelID, sk: connectionID} item, which represents the subscription of a user (connection)
// to a channel
// {pk: channelID, sk: info, channelName: channelName} item, which stores channel info and
// allows us to retrieve all the channelIDs in the reverse GSI
func (ddb *DDBClient) CreateChannel(channelID string, connectionID string) error {
	input := &dynamodb.BatchWriteItemInput{
		RequestItems: map[string][]*dynamodb.WriteRequest{
			ddb.TableName: {
				&dynamodb.WriteRequest{
					PutRequest: &dynamodb.PutRequest{
						Item: map[string]*dynamodb.AttributeValue{
							"pk": {
								S: aws.String(channelID),
							},
							"sk": {
								S: aws.String("info"),
							},
						},
					},
				},
				&dynamodb.WriteRequest{
					PutRequest: &dynamodb.PutRequest{
						Item: map[string]*dynamodb.AttributeValue{
							"pk": {
								S: aws.String(channelID),
							},
							"sk": {
								S: aws.String(connectionID),
							},
						},
					},
				},
			},
		},
		ReturnConsumedCapacity: aws.String("NONE"),
	}
	_, err := ddb.svc.BatchWriteItem(input)
	return err
}

func (ddb *DDBClient) CheckChannelExists(channelID string) (bool, error) {
	input := &dynamodb.GetItemInput{
		Key: map[string]*dynamodb.AttributeValue{
			"pk": {
				S: aws.String(channelID),
			},
			"sk": {
				S: aws.String("info"),
			},
		},
		TableName: aws.String(ddb.TableName),
	}
	output, err := ddb.svc.GetItem(input)
	if err != nil {
		return false, err
	}
	if output.Item != nil {
		return true, nil
	}
	return false, nil
}

// ListChannels retrieves all channel info items (i.e. {sk: info}) in the reverse GSI
func (ddb *DDBClient) ListChannels() ([]string, error) {
	input := &dynamodb.QueryInput{
		ExpressionAttributeValues: map[string]*dynamodb.AttributeValue{
			":sk": {
				S: aws.String("info"),
			},
		},
		KeyConditionExpression: aws.String("sk = :sk"),
		Select:                 aws.String("ALL_PROJECTED_ATTRIBUTES"),
		IndexName:              aws.String(ddb.GSIName),
		TableName:              aws.String(ddb.TableName),
	}
	output, err := ddb.svc.Query(input)
	if err != nil {
		return nil, err
	}

	items := make([]string, 0)
	for _, item := range output.Items {
		items = append(items, *item["pk"].S)
	}
	return items, nil
}

// SubscribeUserToChannel creates an {pk: channelID, sk: connectionID} item, which stores the
// subscription of a user to a channel
func (ddb *DDBClient) SubscribeUserToChannel(connectionID string, channelID string) error {
	input := &dynamodb.PutItemInput{
		Item: map[string]*dynamodb.AttributeValue{
			"pk": &dynamodb.AttributeValue{
				S: aws.String(channelID),
			},
			"sk": &dynamodb.AttributeValue{
				S: aws.String(connectionID),
			},
		},
		ReturnConsumedCapacity: aws.String("NONE"),
		TableName:              aws.String(ddb.TableName),
	}
	_, err := ddb.svc.PutItem(input)
	return err
}

// UserConnected subscribed a user to all channels
func (ddb *DDBClient) UserConnected(connectionID string, channelIDs []string) error {
	input := &dynamodb.BatchWriteItemInput{
		RequestItems: map[string][]*dynamodb.WriteRequest{
			ddb.TableName: make([]*dynamodb.WriteRequest, 0),
		},
		ReturnConsumedCapacity: aws.String("NONE"),
	}
	for _, channelID := range channelIDs {
		input.RequestItems[ddb.TableName] = append(input.RequestItems[ddb.TableName],
			&dynamodb.WriteRequest{
				PutRequest: &dynamodb.PutRequest{
					Item: map[string]*dynamodb.AttributeValue{
						"pk": {
							S: aws.String(channelID),
						},
						"sk": {
							S: aws.String(connectionID),
						},
					},
				},
			},
		)
	}
	_, err := ddb.svc.BatchWriteItem(input)
	return err
}

// UserDisconnected removes a users subscription to all channels
func (ddb *DDBClient) UserDisconnected(connectionID string) error {
	// Query for subscribed channels for user
	input := &dynamodb.QueryInput{
		ExpressionAttributeValues: map[string]*dynamodb.AttributeValue{
			":pk": {
				S: aws.String(connectionID),
			},
		},
		KeyConditionExpression: aws.String("pk = :pk"),
		Select:                 aws.String("ALL_PROJECTED_ATTRIBUTES"),
		IndexName:              aws.String(ddb.GSIName),
		TableName:              aws.String(ddb.TableName),
	}
	_, err := ddb.svc.Query(input)
	if err != nil {
		return err
	}

	// Delete subscribed channels for user
	// input := &dynamodb.BatchWriteItemInput{
	// 	RequestItems: map[string][]*dynamodb.WriteRequest{
	// 		ddb.TableName: make([]*dynamodb.WriteRequest, 0),
	// 	},
	// 	ReturnConsumedCapacity: aws.String("NONE"),
	// }
	// for _, channelID := range channelIDs {
	// 	input.RequestItems[ddb.TableName] = append(input.RequestItems[ddb.TableName],
	// 		&dynamodb.WriteRequest{
	// 			PutRequest: &dynamodb.DeleteRequest{
	// 				Key: map[string]*dynamodb.AttributeValue{
	// 					"pk": {
	// 						S: aws.String(channelID),
	// 					},
	// 					"sk": {
	// 						S: aws.String(connectionID),
	// 					},
	// 				},
	// 			},
	// 		},
	// 	)
	// }
	// _, err := ddb.svc.BatchWriteItem(input)
	return err
}
