package dynamoddb

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb/types"
)

type SaveInput struct {
	TableName                 string
	Item                      map[string]types.AttributeValue
	ConditionExpression       *string
	ExpressionAttributeNames  map[string]string
	ExpressionAttributeValues map[string]types.AttributeValue
}

func (db *DynamoDdb) Save(ctx context.Context, input *SaveInput) error {
	_, err := db.Client.PutItem(ctx, &dynamodb.PutItemInput{
		TableName:                 aws.String(input.TableName),
		Item:                      input.Item,
		ConditionExpression:       input.ConditionExpression,
		ExpressionAttributeNames:  input.ExpressionAttributeNames,
		ExpressionAttributeValues: input.ExpressionAttributeValues,
	})

	return err
}
