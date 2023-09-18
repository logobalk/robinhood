package domain

type UseProfile struct {
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

func (u *UseProfile) Clone() *UseProfile {
	newC := *u

	return &newC
}
