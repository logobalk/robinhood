package models

import "robinhood/usecase"

type AppointmentDetail struct {
	CommentData       []*Detail `json:"commentData"`
	AppointmentDetail *Item     `json:"appointmentDetail"`
}

type Detail struct {
	AppId       string  `json:"appId"`
	Message     string  `json:"message"`
	CreatedBy   string  `json:"createdBy"`
	CreatedDate string  `json:"createdDate"`
	UpdatedBy   *string `json:"updatedBy"`
	UpdatedDate string  `json:"updatedDate"`
}

func NewAppointmentDetail(gad *usecase.GetAppointmentDetailOutput) *AppointmentDetail {
	if gad == nil {
		gad = &usecase.GetAppointmentDetailOutput{}
	}
	var newItems []*Detail
	if gad != nil && len(gad.CommentData) > 0 {
		newItems = make([]*Detail, 0, len(gad.CommentData))
		for _, value := range gad.CommentData {
			newItems = append(newItems, &Detail{
				AppId:       value.AppId,
				Message:     value.Message,
				CreatedBy:   value.CreatedBy,
				CreatedDate: value.CreatedDate,
				UpdatedBy:   value.UpdatedBy,
				UpdatedDate: value.UpdatedDate,
			})
		}
	}
	newAppointmentDetail := &Item{
		AppId:           gad.AppointmentDetail.AppId,
		Title:           gad.AppointmentDetail.Title,
		Description:     gad.AppointmentDetail.Description,
		Status:          gad.AppointmentDetail.Status,
		CreatedBy:       gad.AppointmentDetail.CreatedBy,
		CreateDateTime:  gad.AppointmentDetail.CreateDateTime,
		UpdatedBy:       gad.AppointmentDetail.UpdatedBy,
		UpdatedDateTime: gad.AppointmentDetail.UpdatedDateTime,
		IsActive:        gad.AppointmentDetail.IsActive,
		Email:           gad.AppointmentDetail.Email,
		UserReference:   gad.AppointmentDetail.UserReference,
	}

	return &AppointmentDetail{
		CommentData:       newItems,
		AppointmentDetail: newAppointmentDetail,
	}
}
