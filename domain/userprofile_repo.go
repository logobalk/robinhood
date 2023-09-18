package domain

//go:generate mockery --name=UserProfileRepo
type UserProfileRepo interface {
	SaveUser(userprofile *UseProfile) error
	GetUserById(id string) (*UseProfile, error)
}
