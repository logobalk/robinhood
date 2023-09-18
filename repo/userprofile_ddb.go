package repo

import (
	"context"
	"robinhood/domain"
	"robinhood/utils"
	dynamoddb "robinhood/utils/dynamoDdb"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/feature/dynamodb/expression"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb/types"
)

const (
	userHk = "hk"
	// userRk = "rk"
)

type UserProfileRepoDdb struct {
	tableName string
}

func NewUserProfileRepoDdb(tableName string) *UserProfileRepoDdb {
	return &UserProfileRepoDdb{
		tableName: tableName,
	}
}

func UserProfileRepoDdbSchema() *dynamodb.CreateTableInput {
	return &dynamodb.CreateTableInput{
		TableName: aws.String("user_profile"),
		AttributeDefinitions: []types.AttributeDefinition{
			{
				AttributeName: aws.String(userHk),
				AttributeType: types.ScalarAttributeTypeS,
			},
		},
		KeySchema: []types.KeySchemaElement{
			{
				AttributeName: aws.String(hk),
				KeyType:       types.KeyTypeHash,
			},
		},
		BillingMode: types.BillingModePayPerRequest,
	}
}

func (db *UserProfileRepoDdb) SaveUser(userprofile *domain.UseProfile) error {
	userprofileClone := userprofile.Clone()
	item := utils.MustMarshalMap(userprofileClone, map[string]any{
		hk: userprofile.Id,
	})

	if err := dynamoddb.Save(context.Background(), &dynamoddb.SaveInput{
		TableName: db.tableName,
		Item:      item,
	}); err != nil {
		return err
	}
	return nil
}

func (db *UserProfileRepoDdb) GetUserById(id string) (*domain.UseProfile, error) {
	result, err := dynamoddb.Get[domain.UseProfile](context.Background(), &dynamoddb.GetInput{
		TableName: db.tableName,
		Key: map[string]any{
			hk: id,
		},
	})
	if err != nil {
		return nil, err
	}
	if result == nil {
		return nil, nil
	}

	return result, nil
}

func (db *UserProfileRepoDdb) GetAllUser() ([]*domain.UseProfile, error) {
	expr, err := expression.NewBuilder().
		WithKeyCondition(expression.Key(hk).Equal((expression.Value(userHk)))).
		Build()
	if err != nil {
		return nil, err
	}

	result := make([]*domain.UseProfile, 0, 20)
	if err = dynamoddb.Query(context.Background(), &dynamoddb.QueryInput{
		TableName:                 db.tableName,
		KeyConditionExpression:    expr.KeyCondition(),
		ExpressionAttributeNames:  expr.Names(),
		ExpressionAttributeValues: expr.Values(),
		All:                       true,
	}, &result); err != nil {
		return nil, err
	}

	return result, nil
}
