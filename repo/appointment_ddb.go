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
	hk                 = "hk"
	rk                 = "rk"
	appointmentHkValue = "appointment"
	gsi1               = "gsi1"
	gsi1Rk             = "gsi1Rk"
)

type AppointmentRepoDdb struct {
	tableName string
}

func NewAppointmentRepoDdb(tableName string) *AppointmentRepoDdb {
	return &AppointmentRepoDdb{
		tableName: tableName,
	}
}

func AppointmentRepoDdbSchema() *dynamodb.CreateTableInput {
	return &dynamodb.CreateTableInput{
		TableName: aws.String("appointment"),
		AttributeDefinitions: []types.AttributeDefinition{
			{
				AttributeName: aws.String(hk),
				AttributeType: types.ScalarAttributeTypeS,
			},
			{
				AttributeName: aws.String(rk),
				AttributeType: types.ScalarAttributeTypeS,
			},
			{
				AttributeName: aws.String(gsi1Rk),
				AttributeType: types.ScalarAttributeTypeS,
			},
		},
		KeySchema: []types.KeySchemaElement{
			{
				AttributeName: aws.String(hk),
				KeyType:       types.KeyTypeHash,
			},
			{
				AttributeName: aws.String(rk),
				KeyType:       types.KeyTypeRange,
			},
		},

		BillingMode: types.BillingModePayPerRequest,
		GlobalSecondaryIndexes: []types.GlobalSecondaryIndex{
			{
				IndexName: aws.String(gsi1),
				KeySchema: []types.KeySchemaElement{
					{
						AttributeName: aws.String(hk),
						KeyType:       types.KeyTypeHash,
					},
					{
						AttributeName: aws.String(gsi1Rk),
						KeyType:       types.KeyTypeRange,
					},
				},
				Projection: &types.Projection{
					ProjectionType: types.ProjectionTypeAll,
				},
			},
		},
	}
}

func (db *AppointmentRepoDdb) SaveAppointment(appointment *domain.Appointment) error {
	appointmentClone := appointment.Clone()

	item := utils.MustMarshalMap(appointmentClone, map[string]any{
		hk:     appointmentHkValue,
		rk:     appointmentClone.AppId,
		gsi1Rk: appointmentClone.AppId,
	})

	if err := dynamoddb.Save(context.Background(), &dynamoddb.SaveInput{
		TableName: db.tableName,
		Item:      item,
	}); err != nil {
		return err
	}
	return nil
}

func (db *AppointmentRepoDdb) GetAllItemByLastKey(lastKey string, limit int32) ([]*domain.Appointment, error) {
	expr, err := expression.NewBuilder().
		WithKeyCondition(expression.Key(hk).Equal((expression.Value(appointmentHkValue))).
			And(expression.Key(gsi1Rk).GreaterThan(expression.Value(lastKey)))).
		Build()
	if err != nil {
		return nil, err
	}

	result := make([]*domain.Appointment, 0, limit)
	if err = dynamoddb.Query(context.Background(), &dynamoddb.QueryInput{
		TableName:                 db.tableName,
		KeyConditionExpression:    expr.KeyCondition(),
		ExpressionAttributeNames:  expr.Names(),
		ExpressionAttributeValues: expr.Values(),
		IndexName:                 gsi1,
		Limit:                     limit,
	}, &result); err != nil {
		return nil, err
	}

	return result, nil
}

func (db *AppointmentRepoDdb) GetAppointmentByAppId(appId string) (*domain.Appointment, error) {
	result, err := dynamoddb.Get[domain.Appointment](context.Background(), &dynamoddb.GetInput{
		TableName: db.tableName,
		Key: map[string]any{
			hk: appointmentHkValue,
			rk: appId,
		},
	})
	if err != nil {
		return nil, err
	}

	return result, nil
}
