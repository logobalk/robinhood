package domain

type Comment struct {
	Id            *string
	AppId         string
	Message       string
	CreatedBy     *string
	CreatedDate   string
	UpdatedBy     *string
	UpdatedDate   string
	IsActive      bool
	UserReference string
}

func (a *Comment) Clone() *Comment {
	newC := *a

	return &newC
}
