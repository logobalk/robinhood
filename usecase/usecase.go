package usecase

import (
	"robinhood/domain"
)

type UseCase interface {
	GetAppointmentList(lastKey string, limit int32) (*GetAppointmentListOutput, error)
	SaveAppointment(input *SaveAppointmentInput) error
	GetAppointmentDetail(appId string) (*GetAppointmentDetailOutput, error)
	SaveAppointmentDetail(input *CreateAppointmentDetailInput) error
	SaveUserProfile(input *UserProfileInput) (string, error)
	GetUserProfileById(id string) (*UserProfileOutput, error)
}

type useCase struct {
	appointmentRepo domain.AppointmentRepo
	commentRepo     domain.CommentRepo
	userprofileRepo domain.UserProfileRepo
}

func New(appointmentRepo domain.AppointmentRepo, commentRepo domain.CommentRepo, userprofileRepo domain.UserProfileRepo) *useCase {
	return &useCase{
		appointmentRepo: appointmentRepo,
		commentRepo:     commentRepo,
		userprofileRepo: userprofileRepo,
	}
}
