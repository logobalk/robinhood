package usecase

import (
	"errors"
	"robinhood/domain"
	"robinhood/utils/uuidx"
	"time"

	"github.com/aws/smithy-go/ptr"
)

type GetAppointmentDetailOutput struct {
	CommentData       []*CommentData
	AppointmentDetail *Item
}

type CommentData struct {
	AppId       string
	Message     string
	CreatedBy   string
	CreatedDate string
	UpdatedBy   *string
	UpdatedDate string
}

type CreateAppointmentDetailInput struct {
	Id            *string
	AppId         string
	Message       string
	CreatedBy     *string
	UpdatedBy     *string
	UserReference string
}

func (u *useCase) SaveAppointmentDetail(input *CreateAppointmentDetailInput) error {
	userProfile, err := u.GetUserProfileById(input.UserReference)
	if err != nil {
		return err
	}
	if userProfile == nil {
		return errors.New("fail to create cannot find user reference")
	}

	currentTime := time.Now()
	comment := &domain.Comment{
		Id:            ptr.String(uuidx.New()),
		AppId:         input.AppId,
		Message:       input.Message,
		IsActive:      true,
		UserReference: input.UserReference,
	}
	if input.Id != nil {
		comment.Id = input.Id
		comment.UpdatedDate = currentTime.Format("2006-01-02 15:04:05")
	} else {
		comment.CreatedDate = currentTime.Format("2006-01-02 15:04:05")
		comment.CreatedBy = input.CreatedBy
	}

	if input.UpdatedBy != nil {
		comment.UpdatedBy = input.UpdatedBy
	}

	return u.commentRepo.SaveComment(comment)
}

func (u *useCase) GetAppointmentDetail(appId string) (*GetAppointmentDetailOutput, error) {
	appointment, err := u.appointmentRepo.GetAppointmentByAppId(appId)
	if err != nil {
		return nil, err
	}

	appointmentUserProfile, err := u.GetUserProfileById(appointment.UserReference)
	if err != nil {
		return nil, err
	}
	appointmentDetail := &Item{
		AppId:           appointment.AppId,
		Title:           appointment.Title,
		Description:     appointment.Description,
		Status:          appointment.Status,
		CreatedBy:       ptr.String(appointmentUserProfile.Name),
		CreateDateTime:  appointment.CreateDateTime,
		UpdatedBy:       appointment.UpdatedBy,
		UpdatedDateTime: appointment.UpdatedDateTime,
		Email:           appointmentUserProfile.Email,
		UserReference:   appointment.UserReference,
		IsActive:        appointment.IsActive,
	}

	commentList, err := u.commentRepo.GetAllCommentByAppId(appId)
	if err != nil {
		return nil, err
	}

	if len(commentList) == 0 {
		return &GetAppointmentDetailOutput{
			CommentData:       nil,
			AppointmentDetail: appointmentDetail,
		}, nil
	}

	ensureCommentList := make([]*CommentData, 0, len(commentList))
	for _, val := range commentList {
		if val.IsActive {
			userProfile, err := u.GetUserProfileById(val.UserReference)
			if err != nil {
				return nil, err
			}
			if userProfile == nil {
				userProfile = &UserProfileOutput{
					Name: "",
				}
			}

			ensureCommentList = append(ensureCommentList, &CommentData{
				AppId:       val.AppId,
				Message:     val.Message,
				CreatedBy:   userProfile.Name,
				CreatedDate: val.CreatedDate,
				UpdatedBy:   val.UpdatedBy,
				UpdatedDate: val.UpdatedDate,
			})
		}
	}

	return &GetAppointmentDetailOutput{
		CommentData:       ensureCommentList,
		AppointmentDetail: appointmentDetail,
	}, nil
}
