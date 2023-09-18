package models

import "robinhood/usecase"

type UserProfile struct {
	Id              string `json:"id"`
	UserName        string `json:"userName"`
	Name            string `json:"name"`
	CreatedBy       string `json:"createdBy"`
	CreatedDateTime string `json:"createdDateTime"`
	UpdatedBy       string `json:"updatedBy"`
	UpdatedDateTime string `json:"updatedDateTime"`
	IsActive        bool   `json:"isActive"`
	Email           string `json:"email"`
}

func NewUserProfile(u *usecase.UserProfileOutput) *UserProfile {
	if u == nil {
		u = &usecase.UserProfileOutput{}
	}

	return &UserProfile{
		Id:              u.Id,
		UserName:        u.UserName,
		Name:            u.Name,
		CreatedBy:       u.CreatedBy,
		CreatedDateTime: u.CreatedDateTime,
		UpdatedBy:       u.UpdatedBy,
		UpdatedDateTime: u.CreatedDateTime,
		IsActive:        u.IsActive,
		Email:           u.Email,
	}
}
