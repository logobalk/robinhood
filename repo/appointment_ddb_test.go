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

type AppointmentRepoTestSuite struct {
	suite.Suite
	ctx    context.Context
	schema *dynamodb.CreateTableInput
	repo   *AppointmentRepoDdb
}

func TestAppointmentRepo(t *testing.T) {
	suite.Run(t, new(AppointmentRepoTestSuite))
}

func (t *AppointmentRepoTestSuite) SetupSuite() {
	dynamoddb.DefaultClient = dynamoddb.NewLocalStack()

	t.ctx = context.Background()
	t.schema = AppointmentRepoDdbSchema()
	t.repo = NewAppointmentRepoDdb(*t.schema.TableName)

	dynamoddb.MustCreateTable(t.ctx, t.schema)
}

func (t *AppointmentRepoTestSuite) TearDownTest() {
	dynamoddb.MustDeleteTable(t.ctx, t.schema)
}

func (t *AppointmentRepoTestSuite) TestSaveAndGet() {
	// createTableInput := AppointmentRepoDdbSchema()

	// // Marshal the CreateTableInput structure to JSON with indentation
	// jsonData, err2 := json.MarshalIndent(createTableInput, "", "  ")
	// if err2 != nil {
	// 	fmt.Println("Error:", err2)
	// 	return
	// }

	// // Print the JSON data
	// fmt.Println(string(jsonData))

	input := &domain.Appointment{
		AppId:           ptr.String("10100"),
		Title:           "1",
		Description:     "test",
		Status:          "todo",
		CreatedBy:       "a",
		CreateDateTime:  "10/2/2566",
		UpdatedBy:       ptr.String("a"),
		UpdatedDateTime: "11/2/2566",
		IsActive:        true,
		UserReference:   "1",
	}
	err := t.repo.SaveAppointment(input)
	t.NoError(err)

	input2 := &domain.Appointment{
		AppId:           ptr.String("10102"),
		Title:           "2",
		Description:     "test",
		Status:          "todo",
		CreatedBy:       "a",
		CreateDateTime:  "11/2/2566",
		UpdatedBy:       ptr.String("a"),
		UpdatedDateTime: "11/2/2566",
		IsActive:        true,
		UserReference:   "1",
	}
	err = t.repo.SaveAppointment(input2)
	t.NoError(err)

	input3 := &domain.Appointment{
		AppId:           ptr.String("10103"),
		Title:           "3",
		Description:     "test",
		Status:          "todo",
		CreatedBy:       "a",
		CreateDateTime:  "12/2/2566",
		UpdatedBy:       ptr.String("a"),
		UpdatedDateTime: "11/2/2566",
		IsActive:        true,
		UserReference:   "1",
	}
	err = t.repo.SaveAppointment(input3)
	t.NoError(err)

	input4 := &domain.Appointment{
		AppId:           ptr.String("10104"),
		Title:           "4",
		Description:     "test",
		Status:          "todo",
		CreatedBy:       "a",
		CreateDateTime:  "13/2/2566",
		UpdatedBy:       ptr.String("a"),
		UpdatedDateTime: "11/2/2566",
		IsActive:        true,
		UserReference:   "1",
	}
	err = t.repo.SaveAppointment(input4)
	t.NoError(err)

	input4.Description = "new Desc"
	err = t.repo.SaveAppointment(input4)
	t.NoError(err)

	res, err := t.repo.GetAllItemByLastKey("0", 0)
	t.NoError(err)

	res2, err := t.repo.GetAllItemByLastKey("10102", 2)
	t.NoError(err)
	t.Equal(res[0], input)
	t.Equal(res[1], input2)
	t.Equal(res[2], input3)
	t.Equal(res2[0], input3)
	t.Equal(res2[1], input4)
	t.Equal(res2[1].Description, "new Desc")

	res3, err := t.repo.GetAppointmentByAppId("10103")
	t.NoError(err)
	t.Equal(res3, input3)
}
