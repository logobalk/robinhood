package dynamoddb

import (
	"context"
	"robinhood/utils"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/feature/dynamodb/attributevalue"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb/types"
)

type QueryInput struct {
	TableName                 string
	KeyConditionExpression    *string
	AttributesToGet           []string
	ConsistentRead            *bool
	ExpressionAttributeNames  map[string]string
	ExpressionAttributeValues map[string]types.AttributeValue
	FilterExpression          *string
	IndexName                 string
	Limit                     int32
	SortDesc                  bool
	Select                    types.Select
	All                       bool
}

func (db *DynamoDdb) Query(ctx context.Context, input *QueryInput, data any) error {
	var nextKey map[string]types.AttributeValue

	limit := input.Limit
	if limit == 0 {
		limit = 10
	}

	var result []map[string]types.AttributeValue
	if input.Limit > 0 {
		result = make([]map[string]types.AttributeValue, 0, input.Limit)
	} else {
		result = []map[string]types.AttributeValue{}
	}

	for {
		out, err := db.Client.Query(ctx, &dynamodb.QueryInput{
			TableName:                 aws.String(input.TableName),
			AttributesToGet:           input.AttributesToGet,
			ConsistentRead:            input.ConsistentRead,
			ExclusiveStartKey:         nextKey,
			ExpressionAttributeNames:  input.ExpressionAttributeNames,
			ExpressionAttributeValues: input.ExpressionAttributeValues,
			FilterExpression:          input.FilterExpression,
			IndexName:                 utils.EmptyStringPtr(input.IndexName),
			KeyConditionExpression:    input.KeyConditionExpression,
			Limit:                     aws.Int32(limit),
			ScanIndexForward:          aws.Bool(!input.SortDesc),
			Select:                    input.Select,
		})
		if err != nil {
			return err
		}

		if len(out.Items) == 0 {
			break
		}

		result = append(result, out.Items...)
		if !input.All {
			break
		}

		if out.LastEvaluatedKey != nil {

			nextKey = out.LastEvaluatedKey
		} else {
			break
		}
	}
	err := attributevalue.UnmarshalListOfMaps(result, data)
	if err != nil {
		return err
	}

	return nil
}
