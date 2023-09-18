package models

import (
	"robinhood/domain"
	"robinhood/usecase"
)

type AppointmentList struct {
	LastKey string  `json:"lastKey"`
	Limit   int32   `json:"limit"`
	Data    []*Item `json:"data"`
}

type Item struct {
	AppId           *string       `json:"appId"`
	Title           string        `json:"title"`
	Description     string        `json:"description"`
	Status          domain.Status `json:"status"`
	CreatedBy       *string       `json:"createdBy"`
	CreateDateTime  string        `json:"createDateTime"`
	UpdatedBy       *string       `json:"updatedBy"`
	UpdatedDateTime string        `json:"updatedDateTime"`
	IsActive        bool          `json:"isActive"`
	Email           string        `json:"email"`
	UserReference   string        `json:"userReference"`
}

func NewAppointmentList(ga *usecase.GetAppointmentListOutput) *AppointmentList {
	if ga == nil {
		ga = &usecase.GetAppointmentListOutput{}
	}
	var newItems []*Item
	if ga != nil && len(ga.Data) > 0 {
		newItems = make([]*Item, 0, len(ga.Data))
		for _, value := range ga.Data {
			newItems = append(newItems, &Item{
				AppId:           value.AppId,
				Title:           value.Title,
				Description:     value.Description,
				Status:          value.Status,
				CreatedBy:       value.CreatedBy,
				CreateDateTime:  value.CreateDateTime,
				UpdatedBy:       value.UpdatedBy,
				UpdatedDateTime: value.UpdatedDateTime,
				IsActive:        value.IsActive,
				Email:           value.Email,
				UserReference:   value.UserReference,
			})
		}
	}

	return &AppointmentList{
		LastKey: ga.LastKey,
		Limit:   ga.Limit,
		Data:    newItems,
	}
}
