package usecase

import (
	"robinhood/domain"
	"robinhood/utils/uuidx"
	"time"
)

type UserProfileInput struct {
	UserName  string
	Name      string
	CreatedBy string
	UpdatedBy string
	Email     string
}

type UserProfileOutput struct {
	Id              string
	UserName        string
	Name            string
	CreatedBy       string
	CreatedDateTime string
	UpdatedBy       string
	UpdatedDateTime string
	IsActive        bool
	Email           string
}

func (u *useCase) SaveUserProfile(input *UserProfileInput) (string, error) {
	currentTime := time.Now()
	userProfile := &domain.UseProfile{
		Id:              uuidx.New(),
		UserName:        input.UserName,
		Name:            input.Name,
		CreatedBy:       input.CreatedBy,
		CreatedDateTime: currentTime.Format("2006-01-02 15:04:05"),
		UpdatedBy:       "",
		UpdatedDateTime: "",
		IsActive:        true,
		Email:           input.Email,
	}

	err := u.userprofileRepo.SaveUser(userProfile)
	if err != nil {
		return "", err
	}

	return userProfile.Id, nil
}

func (u *useCase) GetUserProfileById(id string) (*UserProfileOutput, error) {
	userProfile, err := u.userprofileRepo.GetUserById(id)
	if err != nil {
		return nil, err
	}
	if userProfile == nil {
		return nil, nil
	}

	return &UserProfileOutput{
		Id:              userProfile.Id,
		UserName:        userProfile.UserName,
		Name:            userProfile.Name,
		CreatedBy:       userProfile.CreatedBy,
		CreatedDateTime: userProfile.CreatedDateTime,
		UpdatedBy:       userProfile.UpdatedBy,
		UpdatedDateTime: userProfile.UpdatedDateTime,
		IsActive:        userProfile.IsActive,
		Email:           userProfile.Email,
	}, nil
}
