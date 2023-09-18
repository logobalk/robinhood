package usecase

import (
	"errors"
	"robinhood/domain"
	"robinhood/domain/mocks"
	"robinhood/utils/uuidx"
	"testing"
	"time"

	"github.com/stretchr/testify/suite"
)

type UserProfileTestSuite struct {
	suite.Suite
	usecase         *useCase
	appointmentRepo *mocks.AppointmentRepo
	userProfileRepo *mocks.UserProfileRepo
}

func TestUserProfile(t *testing.T) {
	suite.Run(t, new(UserProfileTestSuite))
}

func (t *UserProfileTestSuite) SetupSuite() {
	uuidx.UUIDX = &uuidx.Fixed{Value: "123"}
}

func (t *UserProfileTestSuite) SetupTest() {
	t.appointmentRepo = mocks.NewAppointmentRepo(t.T())
	commentRepo := mocks.NewCommentRepo(t.T())
	t.userProfileRepo = mocks.NewUserProfileRepo(t.T())
	t.usecase = New(t.appointmentRepo, commentRepo, t.userProfileRepo)
}

func (t *UserProfileTestSuite) TestCreateUserProfileSuccess() {
	currentTime := time.Now()
	mockInput := &UserProfileInput{
		UserName:  "dd",
		Name:      "robin",
		CreatedBy: "admin",
		UpdatedBy: "",
		Email:     "d@d.com",
	}
	mockUserProfileInput := &domain.UseProfile{
		Id:              uuidx.New(),
		UserName:        "dd",
		Name:            "robin",
		CreatedBy:       "admin",
		CreatedDateTime: currentTime.Format("2006-01-02 15:04:05"),
		UpdatedBy:       "",
		UpdatedDateTime: "",
		IsActive:        true,
		Email:           "d@d.com",
	}
	t.userProfileRepo.On("SaveUser", mockUserProfileInput).Return(nil)

	id, err := t.usecase.SaveUserProfile(mockInput)
	t.NoError(err)
	t.Equal(id, mockUserProfileInput.Id)

	t.userProfileRepo.AssertExpectations(t.T())
}

func (t *UserProfileTestSuite) TestCreateUserProfileFail_WhenSaveUserError() {
	currentTime := time.Now()
	mockInput := &UserProfileInput{
		UserName:  "dd",
		Name:      "robin",
		CreatedBy: "admin",
		UpdatedBy: "",
		Email:     "d@d.com",
	}
	mockUserProfileInput := &domain.UseProfile{
		Id:              uuidx.New(),
		UserName:        "dd",
		Name:            "robin",
		CreatedBy:       "admin",
		CreatedDateTime: currentTime.Format("2006-01-02 15:04:05"),
		UpdatedBy:       "",
		UpdatedDateTime: "",
		IsActive:        true,
		Email:           "d@d.com",
	}
	t.userProfileRepo.On("SaveUser", mockUserProfileInput).Return(errors.New("test error"))

	id, err := t.usecase.SaveUserProfile(mockInput)
	t.Error(err)
	t.Equal(id, "")

	t.userProfileRepo.AssertExpectations(t.T())
}

func (t *UserProfileTestSuite) TestGetUserProfileByIdSuccess() {
	mockUserProfileOutput := &domain.UseProfile{
		Id:              "1",
		UserName:        "dd",
		Name:            "robin",
		CreatedBy:       "admin",
		CreatedDateTime: "10/10/2566",
		UpdatedBy:       "",
		UpdatedDateTime: "",
		IsActive:        true,
		Email:           "d@d.com",
	}
	expect := &UserProfileOutput{
		Id:              "1",
		UserName:        "dd",
		Name:            "robin",
		CreatedBy:       "admin",
		CreatedDateTime: "10/10/2566",
		UpdatedBy:       "",
		UpdatedDateTime: "",
		IsActive:        true,
		Email:           "d@d.com",
	}
	t.userProfileRepo.On("GetUserById", "1").Return(mockUserProfileOutput, nil)

	result, err := t.usecase.GetUserProfileById("1")
	t.NoError(err)
	t.Equal(expect, result)

	t.userProfileRepo.AssertExpectations(t.T())
}

func (t *UserProfileTestSuite) TestGetUserProfileByIdSuccess_WhenUserProfileIsNil() {
	t.userProfileRepo.On("GetUserById", "0").Return(nil, nil)

	result, err := t.usecase.GetUserProfileById("0")
	t.NoError(err)
	t.Nil(result)

	t.userProfileRepo.AssertExpectations(t.T())
}

func (t *UserProfileTestSuite) TestGetUserProfileByIdFail_WhenGetUserProfileByIdError() {
	t.userProfileRepo.On("GetUserById", "1").Return(nil, errors.New("error"))

	result, err := t.usecase.GetUserProfileById("1")
	t.Error(err)
	t.Nil(result)

	t.userProfileRepo.AssertExpectations(t.T())
}
