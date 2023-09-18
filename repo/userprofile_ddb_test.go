package repo

import (
	"context"
	"fmt"
	"robinhood/domain"
	dynamoddb "robinhood/utils/dynamoDdb"
	"testing"

	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	"github.com/stretchr/testify/suite"
)

type UserProfileRepoTestSuite struct {
	suite.Suite
	ctx    context.Context
	schema *dynamodb.CreateTableInput
	repo   *UserProfileRepoDdb
}

func TestUserProfileRepoTestSuiteRepo(t *testing.T) {
	suite.Run(t, new(UserProfileRepoTestSuite))
}

func (t *UserProfileRepoTestSuite) SetupSuite() {
	dynamoddb.DefaultClient = dynamoddb.NewLocalStack()
	t.ctx = context.Background()
	t.schema = UserProfileRepoDdbSchema()
	t.repo = NewUserProfileRepoDdb(*t.schema.TableName)

	dynamoddb.MustCreateTable(t.ctx, t.schema)
}

func (t *UserProfileRepoTestSuite) TearDownTest() {
	dynamoddb.MustDeleteTable(t.ctx, t.schema)
}

func (t *UserProfileRepoTestSuite) TestSaveAndGet() {
	input := &domain.UseProfile{
		Id:              "1",
		UserName:        "aa",
		Name:            "robin",
		CreatedBy:       "d",
		CreatedDateTime: "10/10/2566",
		UpdatedBy:       "d",
		UpdatedDateTime: "",
		IsActive:        true,
		Email:           "d@d.com",
	}

	err := t.repo.SaveUser(input)
	t.NoError(err)

	input.Email = "test@t.com"
	err = t.repo.SaveUser(input)
	t.NoError(err)

	res, err := t.repo.GetUserById("1")
	t.NoError(err)
	t.Equal(res.Email, "test@t.com")

	res, err = t.repo.GetUserById("0")
	t.NoError(err)
	t.Nil(res)

	res3, err := t.repo.GetAllUser()
	fmt.Println("res3==>", res3)
	t.NoError(err)
}
