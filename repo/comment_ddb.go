package repo

import (
	"context"
	"robinhood/domain"
	"robinhood/utils"
	dynamoddb "robinhood/utils/dynamoDdb"

	"github.com/aws/aws-sdk-go-v2/feature/dynamodb/expression"
)

const (
	commentHkValue = "comment"
)

type CommentRepoDdb struct {
	tableName string
}

func NewCommentRepoDdb(tableName string) *CommentRepoDdb {
	return &CommentRepoDdb{
		tableName: tableName,
	}
}

func (db *CommentRepoDdb) SaveComment(comment *domain.Comment) error {
	commentClone := comment.Clone()
	item := utils.MustMarshalMap(commentClone, map[string]any{
		hk:     commentHkValue,
		rk:     commentClone.Id,
		gsi1Rk: commentClone.AppId,
	})

	if err := dynamoddb.Save(context.Background(), &dynamoddb.SaveInput{
		TableName: db.tableName,
		Item:      item,
	}); err != nil {
		return err
	}
	return nil
}

func (db *CommentRepoDdb) GetAllCommentByAppId(appId string) ([]*domain.Comment, error) {
	expr, err := expression.NewBuilder().WithKeyCondition(expression.Key(hk).Equal(expression.Value(commentHkValue)).
		And(expression.Key(gsi1Rk).Equal(expression.Value(appId)))).Build()
	if err != nil {
		return nil, err
	}

	result := make([]*domain.Comment, 0, 20)
	if err := dynamoddb.Query(context.Background(), &dynamoddb.QueryInput{
		TableName:                 db.tableName,
		KeyConditionExpression:    expr.KeyCondition(),
		ExpressionAttributeNames:  expr.Names(),
		ExpressionAttributeValues: expr.Values(),
		SortDesc:                  true,
		IndexName:                 gsi1,
		All:                       true,
	}, &result); err != nil {
		return nil, err
	}

	return result, nil
}
