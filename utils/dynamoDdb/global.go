package dynamoddb

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
)

var DefaultClient = New()

func Save(ctx context.Context, input *SaveInput) error {
	return DefaultClient.Save(ctx, input)
}

func Query(ctx context.Context, input *QueryInput, data any) error {
	return DefaultClient.Query(ctx, input, &data)
}

func Get[T any](ctx context.Context, input *GetInput) (*T, error) {
	data := new(T)
	err, exist := DefaultClient.Get(ctx, input, data)
	if err != nil {
		return nil, err
	}

	if !exist {
		return nil, nil
	}

	return data, nil
}

func MustCreateTable(ctx context.Context, schema *dynamodb.CreateTableInput) {
	DefaultClient.MustCreateTable(ctx, schema)
}

func MustDeleteTable(ctx context.Context, schema *dynamodb.CreateTableInput) {
	DefaultClient.MustDeleteTable(ctx, schema)
}
