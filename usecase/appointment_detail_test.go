package usecase

import (
	"robinhood/domain"
	"robinhood/domain/mocks"
	"robinhood/utils/uuidx"
	"testing"
	"time"

	"github.com/aws/smithy-go/ptr"
	"github.com/stretchr/testify/suite"
)

type GetAppointmentDetailTestSuite struct {
	suite.Suite
	usecase         *useCase
	appointmentRepo *mocks.AppointmentRepo
	commentRepo     *mocks.CommentRepo
	userProfileRepo *mocks.UserProfileRepo
}

func TestGetAppointmentDetail(t *testing.T) {
	suite.Run(t, new(GetAppointmentDetailTestSuite))
}

func (t *GetAppointmentDetailTestSuite) SetupSuite() {
	uuidx.UUIDX = &uuidx.Fixed{Value: "123"}
}

func (t *GetAppointmentDetailTestSuite) SetupTest() {
	t.appointmentRepo = mocks.NewAppointmentRepo(t.T())
	t.commentRepo = mocks.NewCommentRepo(t.T())
	t.userProfileRepo = mocks.NewUserProfileRepo(t.T())
	t.usecase = New(t.appointmentRepo, t.commentRepo, t.userProfileRepo)
}

func (t *GetAppointmentDetailTestSuite) TestCreateAppointmentDetailSuccess() {
	currentTime := time.Now()
	mockInput := &CreateAppointmentDetailInput{
		AppId:         "1",
		Message:       "a",
		CreatedBy:     ptr.String("a"),
		UserReference: "123",
	}

	mockCommentRepoInput := &domain.Comment{
		Id:            ptr.String(uuidx.New()),
		AppId:         "1",
		Message:       "a",
		CreatedBy:     ptr.String("a"),
		CreatedDate:   currentTime.Format("2006-01-02 15:04:05"),
		IsActive:      true,
		UserReference: "123",
	}

	mockUserProfileOutput := &domain.UseProfile{
		Id:              "123",
		UserName:        "d",
		Name:            "test",
		CreatedBy:       "admin",
		CreatedDateTime: "10/05/2566",
		UpdatedBy:       "",
		UpdatedDateTime: "",
		IsActive:        true,
		Email:           "da@da.com",
	}

	t.userProfileRepo.On("GetUserById", "123").Return(mockUserProfileOutput, nil)
	t.commentRepo.On("SaveComment", mockCommentRepoInput).Return(nil)

	err := t.usecase.SaveAppointmentDetail(mockInput)
	t.NoError(err)

	t.commentRepo.AssertExpectations(t.T())
}

func (t *GetAppointmentDetailTestSuite) TestUpdateAppointmentDetailSuccess() {
	currentTime := time.Now()
	mockInput := &CreateAppointmentDetailInput{
		Id:            ptr.String(uuidx.UUIDX.New()),
		AppId:         "1",
		Message:       "a",
		UpdatedBy:     ptr.String("b"),
		UserReference: "123",
	}

	mockCommentRepoInput := &domain.Comment{
		Id:            ptr.String(uuidx.UUIDX.New()),
		AppId:         "1",
		Message:       "a",
		UpdatedBy:     ptr.String("b"),
		UpdatedDate:   currentTime.Format("2006-01-02 15:04:05"),
		IsActive:      true,
		UserReference: "123",
	}
	mockUserProfileOutput := &domain.UseProfile{
		Id:              "123",
		UserName:        "d",
		Name:            "test",
		CreatedBy:       "admin",
		CreatedDateTime: "10/05/2566",
		UpdatedBy:       "",
		UpdatedDateTime: "",
		IsActive:        true,
		Email:           "da@da.com",
	}

	t.userProfileRepo.On("GetUserById", "123").Return(mockUserProfileOutput, nil)
	t.commentRepo.On("SaveComment", mockCommentRepoInput).Return(nil)

	err := t.usecase.SaveAppointmentDetail(mockInput)
	t.NoError(err)

	t.commentRepo.AssertExpectations(t.T())
}

func (t *GetAppointmentDetailTestSuite) TestGetAppointmentDetailSuccess() {
	mockAppointmentOutput := &domain.Appointment{
		AppId:           ptr.String("1"),
		Title:           "a",
		Description:     "aa",
		Status:          domain.StatusTodo,
		CreatedBy:       "a",
		CreateDateTime:  "09/05/2566",
		UpdatedBy:       ptr.String("a"),
		UpdatedDateTime: "",
		IsActive:        true,
		Email:           "d@d.com",
		UserReference:   "123",
	}
	mockCommentListOutput := []*domain.Comment{
		{
			Id:            ptr.String("1"),
			AppId:         "1",
			Message:       "aa",
			CreatedBy:     ptr.String("a"),
			CreatedDate:   "10/05/2566",
			IsActive:      true,
			UserReference: "1",
		},
		{
			Id:            ptr.String("3"),
			AppId:         "1",
			Message:       "aa",
			CreatedBy:     ptr.String("a"),
			CreatedDate:   "10/05/2566",
			IsActive:      false,
			UserReference: "1",
		},
	}

	mockOutput := &GetAppointmentDetailOutput{
		CommentData: []*CommentData{{
			AppId:       "1",
			Message:     "aa",
			CreatedBy:   "test",
			CreatedDate: "10/05/2566",
		}},
		AppointmentDetail: &Item{
			AppId:           ptr.String("1"),
			Title:           "a",
			Description:     "aa",
			Status:          domain.StatusTodo,
			CreatedBy:       ptr.String("test"),
			CreateDateTime:  "09/05/2566",
			UpdatedBy:       ptr.String("a"),
			UpdatedDateTime: "",
			IsActive:        true,
			Email:           "d@d.com",
			UserReference:   "123",
		},
	}

	mockUserProfileOutput := &domain.UseProfile{
		Id:              "123",
		UserName:        "d",
		Name:            "test",
		CreatedBy:       "admin",
		CreatedDateTime: "10/05/2566",
		UpdatedBy:       "",
		UpdatedDateTime: "",
		IsActive:        true,
		Email:           "d@d.com",
	}

	mockAppointmentUserProfileOutput := &domain.UseProfile{
		Id:              "1",
		UserName:        "d",
		Name:            "test",
		CreatedBy:       "admin",
		CreatedDateTime: "10/05/2566",
		UpdatedBy:       "",
		UpdatedDateTime: "",
		IsActive:        true,
		Email:           "d@d.com",
	}

	t.appointmentRepo.On("GetAppointmentByAppId", "1").Return(mockAppointmentOutput, nil)
	t.commentRepo.On("GetAllCommentByAppId", "1").Return(mockCommentListOutput, nil)
	t.userProfileRepo.On("GetUserById", "123").Return(mockAppointmentUserProfileOutput, nil).Once()
	t.userProfileRepo.On("GetUserById", "1").Return(mockUserProfileOutput, nil)

	result, err := t.usecase.GetAppointmentDetail("1")
	t.NoError(err)
	t.Equal(result, mockOutput)

	t.commentRepo.AssertExpectations(t.T())
	t.appointmentRepo.AssertExpectations(t.T())
}
