package dynamoddb

import (
	"context"
	"robinhood/utils"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/feature/dynamodb/attributevalue"
	"github.com/aws/aws-sdk-go-v2/feature/dynamodb/expression"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb/types"
)

type GetInput struct {
	TableName            string
	IndexName            string
	Key                  map[string]any
	ConsistentRead       *bool
	AttributesToGet      []string
	ProjectionExpression string
	Select               types.Select
}

func (db *DynamoDdb) Get(ctx context.Context, input *GetInput, data any) (error, bool) {
	if input.IndexName != "" {
		return db.getByIndex(ctx, input, data)
	}

	out, err := db.Client.GetItem(ctx, &dynamodb.GetItemInput{
		Key:                  utils.MustMarshalMap(input.Key),
		TableName:            aws.String(input.TableName),
		ConsistentRead:       input.ConsistentRead,
		ProjectionExpression: utils.EmptyStringPtr(input.ProjectionExpression),
		AttributesToGet:      input.AttributesToGet,
	})
	if err != nil {
		return err, false
	}

	if out.Item == nil {
		return nil, false
	}

	err = attributevalue.UnmarshalMap(out.Item, data)
	if err != nil {
		return err, false
	}

	return nil, true
}

func (db *DynamoDdb) getByIndex(ctx context.Context, input *GetInput, data any) (error, bool) {
	builder := expression.NewBuilder()
	var keyBuilder expression.KeyConditionBuilder
	for key, val := range input.Key {
		if !keyBuilder.IsSet() {
			keyBuilder = expression.KeyEqual(expression.Key(key), expression.Value(val))
		} else {
			keyBuilder = keyBuilder.And(expression.KeyEqual(expression.Key(key), expression.Value(val)))
		}
	}

	expr, err := builder.WithKeyCondition(keyBuilder).Build()
	if err != nil {
		return err, false
	}

	out, err := db.Client.Query(ctx, &dynamodb.QueryInput{
		TableName:                 aws.String(input.TableName),
		IndexName:                 aws.String(input.IndexName),
		ConsistentRead:            input.ConsistentRead,
		ExpressionAttributeNames:  expr.Names(),
		ExpressionAttributeValues: expr.Values(),
		KeyConditionExpression:    expr.KeyCondition(),
		Limit:                     aws.Int32(1),
		AttributesToGet:           input.AttributesToGet,
		ProjectionExpression:      utils.EmptyStringPtr(input.ProjectionExpression),
		Select:                    input.Select,
	})
	if err != nil {
		return err, false
	}

	if out.Items == nil || len(out.Items) == 0 {
		return nil, false
	}

	err = attributevalue.UnmarshalMap(out.Items[0], data)
	if err != nil {
		return err, false
	}

	return nil, true
}
