package usecase

import (
	"errors"
	"robinhood/domain"
	"robinhood/utils/uuidx"
	"time"

	"github.com/aws/smithy-go/ptr"
)

type GetAppointmentListOutput struct {
	LastKey string
	Limit   int32
	Data    []*Item
}

type Item struct {
	AppId           *string
	Title           string
	Description     string
	Status          domain.Status
	CreatedBy       *string
	CreateDateTime  string
	UpdatedBy       *string
	UpdatedDateTime string
	Email           string
	UserReference   string
	IsActive        bool
}

type SaveAppointmentInput struct {
	AppId         *string
	Title         string
	Description   string
	CreatedBy     string
	UpdatedBy     *string
	Email         string
	UserReference string
	Status        domain.Status
	IsActive      bool
}

func (u *useCase) GetAppointmentList(lastKey string, limit int32) (*GetAppointmentListOutput, error) {
	appointmentList, err := u.appointmentRepo.GetAllItemByLastKey(lastKey, limit)
	if err != nil {
		return nil, err
	}

	if len(appointmentList) == 0 {
		return nil, nil
	}

	ensureAppointmentList := make([]*Item, 0, len(appointmentList))
	for _, val := range appointmentList {
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
			ensureAppointmentList = append(ensureAppointmentList, &Item{
				AppId:           val.AppId,
				Title:           val.Title,
				Description:     val.Description,
				Status:          val.Status,
				CreatedBy:       ptr.String(userProfile.Name),
				CreateDateTime:  val.CreateDateTime,
				UpdatedBy:       val.UpdatedBy,
				UpdatedDateTime: val.UpdatedDateTime,
				Email:           userProfile.Email,
				UserReference:   val.UserReference,
				IsActive:        val.IsActive,
			})
		}
	}

	return &GetAppointmentListOutput{
		LastKey: *ensureAppointmentList[len(ensureAppointmentList)-1].AppId,
		Limit:   limit,
		Data:    ensureAppointmentList,
	}, nil
}

func (u *useCase) SaveAppointment(input *SaveAppointmentInput) error {
	userProfile, err := u.GetUserProfileById(input.UserReference)
	if err != nil {
		return err
	}
	if userProfile == nil {
		return errors.New("fail to create cannot find user reference")
	}

	currentTime := time.Now()
	appointment := &domain.Appointment{
		AppId:         ptr.String(uuidx.New()),
		Title:         input.Title,
		Description:   input.Description,
		Status:        domain.StatusTodo,
		IsActive:      true,
		Email:         input.Email,
		UserReference: input.UserReference,
	}
	if input.AppId != nil {
		appointment.AppId = input.AppId
		appointment.UpdatedDateTime = currentTime.Format("2006-01-02 15:04:05")
		appointment.Status = input.Status
		appointment.IsActive = input.IsActive
	} else {
		appointment.CreateDateTime = currentTime.Format("2006-01-02 15:04:05")
		appointment.CreatedBy = input.CreatedBy
	}

	if input.UpdatedBy != nil {
		appointment.UpdatedBy = input.UpdatedBy
	}

	return u.appointmentRepo.SaveAppointment(appointment)
}
