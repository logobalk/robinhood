package domain

type Status string

const (
	StatusTodo       Status = "todo"
	StatusInProgress Status = "in_progress"
	StatusDone       Status = "done"
)

type Appointment struct {
	AppId           *string
	Title           string
	Description     string
	Status          Status
	CreatedBy       string
	CreateDateTime  string
	UpdatedBy       *string
	UpdatedDateTime string
	IsActive        bool
	Email           string
	UserReference   string
}

func (a *Appointment) Clone() *Appointment {
	newC := *a

	return &newC
}
