package repo

import (
	"context"
	"robinhood/domain"
	dynamoddb "robinhood/utils/dynamoDdb"
	"testing"

	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	"github.com/aws/smithy-go/ptr"
	"github.com/stretchr/testify/suite"
)

type CommentRepoTestSuite struct {
	suite.Suite
	ctx    context.Context
	schema *dynamodb.CreateTableInput
	repo   *CommentRepoDdb
}

func TestCommentRepo(t *testing.T) {
	suite.Run(t, new(CommentRepoTestSuite))
}

func (t *CommentRepoTestSuite) SetupSuite() {
	dynamoddb.DefaultClient = dynamoddb.NewLocalStack()

	t.ctx = context.Background()
	t.schema = AppointmentRepoDdbSchema()
	t.repo = NewCommentRepoDdb(*t.schema.TableName)

	dynamoddb.MustCreateTable(t.ctx, t.schema)
}

func (t *CommentRepoTestSuite) TearDownTest() {
	dynamoddb.MustDeleteTable(t.ctx, t.schema)
}

func (t *CommentRepoTestSuite) TestSaveAndGet() {
	input := &domain.Comment{
		Id:          ptr.String("0"),
		AppId:       "1",
		Message:     "test",
		CreatedBy:   ptr.String("a"),
		CreatedDate: "10/2/2566",
		UpdatedBy:   ptr.String("d"),
		UpdatedDate: "11/2/2566",
		IsActive:    true,
	}

	err := t.repo.SaveComment(input)
	t.NoError(err)

	input2 := &domain.Comment{
		Id:          ptr.String("1"),
		AppId:       "1",
		Message:     "test",
		CreatedBy:   ptr.String("a"),
		CreatedDate: "10/2/2566",
		UpdatedBy:   ptr.String("d"),
		UpdatedDate: "11/2/2566",
		IsActive:    true,
	}

	err = t.repo.SaveComment(input2)
	t.NoError(err)

	input2.Message = "test update message"
	err = t.repo.SaveComment(input2)
	t.NoError(err)

	res, err := t.repo.GetAllCommentByAppId("1")
	t.NoError(err)
	t.Equal(res[0], input)
	t.Equal(res[1], input2)
	t.Equal(res[1].Message, "test update message")
}
