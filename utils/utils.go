package utils

import (
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/feature/dynamodb/attributevalue"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb/types"
)

func MustMarshalMap(data any, merge ...map[string]any) map[string]types.AttributeValue {
	if d, ok := data.(map[string]types.AttributeValue); ok {
		return d
	}

	result, err := attributevalue.MarshalMap(data)
	if err != nil {
		panic(err)
	}

	if len(merge) > 0 {
		mergeItem := MustMarshalMap(merge[0])
		for k, v := range mergeItem {
			result[k] = v
		}
	}

	return result
}

func EmptyStringPtr(val string) *string {
	if val == "" {
		return nil
	}

	return aws.String(val)
}
