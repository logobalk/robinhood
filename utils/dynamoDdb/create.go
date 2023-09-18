package dynamoddb

import (
	"context"
	"errors"
	"time"

	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb/types"
)

func (db *DynamoDdb) MustCreateTable(ctx context.Context, schema *dynamodb.CreateTableInput) {
	_, err := db.Client.CreateTable(ctx, (*dynamodb.CreateTableInput)(schema))
	if err != nil {
		var xerr *types.ResourceInUseException
		if errors.As(err, &xerr) {
			return
		} else {
			panic(err)
		}
	}

	params := &dynamodb.DescribeTableInput{
		TableName: schema.TableName,
	}

	waiter := dynamodb.NewTableExistsWaiter(db)
	err = waiter.Wait(ctx, params, 5*time.Minute)
	if err != nil {
		panic(err)
	}
}

func (db *DynamoDdb) MustDeleteTable(ctx context.Context, schema *dynamodb.CreateTableInput) {
	_, err := db.Client.DeleteTable(ctx, &dynamodb.DeleteTableInput{
		TableName: schema.TableName,
	})
	if err != nil {
		var xerr *types.ResourceNotFoundException
		if errors.As(err, &xerr) {
			return
		} else {
			panic(err)
		}
	}

	params := &dynamodb.DescribeTableInput{
		TableName: schema.TableName,
	}

	waiter := dynamodb.NewTableNotExistsWaiter(db)
	err = waiter.Wait(ctx, params, 5*time.Minute)
	if err != nil {
		panic(err)
	}
}
