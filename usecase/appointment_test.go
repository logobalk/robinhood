package usecase

import (
	"errors"
	"robinhood/domain"
	"robinhood/domain/mocks"
	"robinhood/utils/uuidx"
	"testing"
	"time"

	"github.com/aws/smithy-go/ptr"
	"github.com/stretchr/testify/suite"
)

type GetAppointmentListTestSuite struct {
	suite.Suite
	usecase         *useCase
	appointmentRepo *mocks.AppointmentRepo
	userProfileRepo *mocks.UserProfileRepo
}

func TestGetAppointmentList(t *testing.T) {
	suite.Run(t, new(GetAppointmentListTestSuite))
}

func (t *GetAppointmentListTestSuite) SetupSuite() {
	uuidx.UUIDX = &uuidx.Fixed{Value: "123"}
}

func (t *GetAppointmentListTestSuite) SetupTest() {
	t.appointmentRepo = mocks.NewAppointmentRepo(t.T())
	commentRepo := mocks.NewCommentRepo(t.T())
	t.userProfileRepo = mocks.NewUserProfileRepo(t.T())
	t.usecase = New(t.appointmentRepo, commentRepo, t.userProfileRepo)
}

func (t *GetAppointmentListTestSuite) TestGetAppointmentListSuccess() {
	mockAppointmentListOutput := []*domain.Appointment{{
		AppId:           ptr.String("1"),
		Title:           "a",
		Description:     "b",
		Status:          domain.StatusTodo,
		CreatedBy:       "a",
		CreateDateTime:  "10/10/2566",
		UpdatedBy:       ptr.String("d"),
		UpdatedDateTime: "11/11/2566",
		IsActive:        true,
		Email:           "d@d.com",
		UserReference:   "1",
	}, {
		AppId:          ptr.String("2"),
		Title:          "a",
		Description:    "b",
		Status:         domain.StatusTodo,
		CreatedBy:      "a",
		CreateDateTime: "10/10/2566",
		IsActive:       false,
		Email:          "d@d.com",
		UserReference:  "1",
	}}
	mockUserProfileOutput := &domain.UseProfile{
		Id:              "1",
		UserName:        "d",
		Name:            "test",
		CreatedBy:       "admin",
		CreatedDateTime: "10/05/2566",
		UpdatedBy:       "",
		UpdatedDateTime: "",
		IsActive:        true,
		Email:           "da@da.com",
	}

	expect := &GetAppointmentListOutput{
		LastKey: "1",
		Limit:   10,
		Data: []*Item{{
			AppId:           ptr.String("1"),
			Title:           "a",
			Description:     "b",
			Status:          domain.StatusTodo,
			CreatedBy:       ptr.String("test"),
			CreateDateTime:  "10/10/2566",
			UpdatedBy:       ptr.String("d"),
			UpdatedDateTime: "11/11/2566",
			IsActive:        true,
			Email:           "da@da.com",
			UserReference:   "1",
		}},
	}

	t.appointmentRepo.On("GetAllItemByLastKey", "1", int32(10)).Return(mockAppointmentListOutput, nil)
	t.userProfileRepo.On("GetUserById", "1").Return(mockUserProfileOutput, nil)

	result, err := t.usecase.GetAppointmentList("1", int32(10))
	t.NoError(err)
	t.Equal(result, expect)

	t.appointmentRepo.AssertExpectations(t.T())
}

func (t *GetAppointmentListTestSuite) TestGetAppointmentListSuccess_WhenAppointmentListNotFound() {
	t.appointmentRepo.On("GetAllItemByLastKey", "1", int32(10)).Return(nil, nil)

	result, err := t.usecase.GetAppointmentList("1", int32(10))
	t.NoError(err)
	t.Nil(result)

	t.appointmentRepo.AssertExpectations(t.T())
}

func (t *GetAppointmentListTestSuite) TestCreateAppointmentSuccess() {
	currentTime := time.Now()
	mockInput := &SaveAppointmentInput{
		Title:         "a",
		Description:   "a",
		CreatedBy:     "a",
		Email:         "d@d.com",
		UserReference: "1",
	}

	mockCreateInput := &domain.Appointment{
		AppId:          ptr.String(uuidx.UUIDX.New()),
		Title:          mockInput.Title,
		Description:    mockInput.Description,
		Status:         domain.StatusTodo,
		CreatedBy:      mockInput.CreatedBy,
		CreateDateTime: currentTime.Format("2006-01-02 15:04:05"),
		IsActive:       true,
		Email:          mockInput.Email,
		UserReference:  mockInput.UserReference,
	}

	mockUserProfileOutput := &domain.UseProfile{
		Id:              "1",
		UserName:        "d",
		Name:            "test",
		CreatedBy:       "admin",
		CreatedDateTime: "10/05/2566",
		UpdatedBy:       "",
		UpdatedDateTime: "",
		IsActive:        true,
		Email:           "da@da.com",
	}

	t.userProfileRepo.On("GetUserById", "1").Return(mockUserProfileOutput, nil)
	t.appointmentRepo.On("SaveAppointment", mockCreateInput).Return(nil)

	err := t.usecase.SaveAppointment(mockInput)
	t.NoError(err)

	t.appointmentRepo.AssertExpectations(t.T())
}

func (t *GetAppointmentListTestSuite) TestCreateAppointmentSuccess_WhenGetUserByIdIsNil() {
	mockInput := &SaveAppointmentInput{
		Title:         "a",
		Description:   "a",
		CreatedBy:     "a",
		Email:         "d@d.com",
		UserReference: "1",
	}

	t.userProfileRepo.On("GetUserById", "1").Return(nil, nil)

	err := t.usecase.SaveAppointment(mockInput)
	t.Error(err)
	t.Equal(err.Error(), "fail to create cannot find user reference")

	t.appointmentRepo.AssertExpectations(t.T())
}

func (t *GetAppointmentListTestSuite) TestSaveAppointmentFail_WhenSaveAppointmentError() {
	currentTime := time.Now()
	mockInput := &SaveAppointmentInput{
		Title:         "a",
		Description:   "a",
		CreatedBy:     "a",
		Email:         "d@d.com",
		UserReference: "1",
	}

	mockCreateInput := &domain.Appointment{
		AppId:          ptr.String(uuidx.UUIDX.New()),
		Title:          mockInput.Title,
		Description:    mockInput.Description,
		Status:         domain.StatusTodo,
		CreatedBy:      mockInput.CreatedBy,
		CreateDateTime: currentTime.Format("2006-01-02 15:04:05"),
		IsActive:       true,
		Email:          mockInput.Email,
		UserReference:  mockInput.UserReference,
	}

	mockUserProfileOutput := &domain.UseProfile{
		Id:              "1",
		UserName:        "d",
		Name:            "test",
		CreatedBy:       "admin",
		CreatedDateTime: "10/05/2566",
		UpdatedBy:       "",
		UpdatedDateTime: "",
		IsActive:        true,
		Email:           "da@da.com",
	}

	t.userProfileRepo.On("GetUserById", "1").Return(mockUserProfileOutput, nil)
	t.appointmentRepo.On("SaveAppointment", mockCreateInput).Return(errors.New("test"))

	err := t.usecase.SaveAppointment(mockInput)
	t.Error(err)

	t.appointmentRepo.AssertExpectations(t.T())
}

func (t *GetAppointmentListTestSuite) TestSaveAppointmentFail_WhenGetUserByIdError() {
	mockInput := &SaveAppointmentInput{
		Title:         "a",
		Description:   "a",
		CreatedBy:     "a",
		Email:         "d@d.com",
		UserReference: "1",
	}

	t.userProfileRepo.On("GetUserById", "1").Return(nil, errors.New("test"))

	err := t.usecase.SaveAppointment(mockInput)
	t.Error(err)

	t.appointmentRepo.AssertExpectations(t.T())
}

func (t *GetAppointmentListTestSuite) TestUpdateAppointmentSuccess() {
	currentTime := time.Now()
	mockInput := &SaveAppointmentInput{
		AppId:         ptr.String("1"),
		Title:         "a",
		Description:   "a",
		UpdatedBy:     ptr.String("a"),
		Email:         "d@d.com",
		UserReference: "1",
		Status:        domain.StatusInProgress,
		IsActive:      true,
	}

	mockUpdateInput := &domain.Appointment{
		AppId:           ptr.String("1"),
		Title:           mockInput.Title,
		Description:     mockInput.Description,
		Status:          domain.StatusInProgress,
		UpdatedBy:       ptr.String("a"),
		UpdatedDateTime: currentTime.Format("2006-01-02 15:04:05"),
		IsActive:        true,
		Email:           mockInput.Email,
		UserReference:   mockInput.UserReference,
	}

	mockUserProfileOutput := &domain.UseProfile{
		Id:              "1",
		UserName:        "d",
		Name:            "test",
		CreatedBy:       "admin",
		CreatedDateTime: "10/05/2566",
		UpdatedBy:       "",
		UpdatedDateTime: "",
		IsActive:        true,
		Email:           "da@da.com",
	}

	t.userProfileRepo.On("GetUserById", "1").Return(mockUserProfileOutput, nil)
	t.appointmentRepo.On("SaveAppointment", mockUpdateInput).Return(nil)

	err := t.usecase.SaveAppointment(mockInput)
	t.NoError(err)

	t.appointmentRepo.AssertExpectations(t.T())
}
