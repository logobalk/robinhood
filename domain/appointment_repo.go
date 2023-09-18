package domain

//go:generate mockery --name=AppointmentRepo
type AppointmentRepo interface {
	SaveAppointment(appointment *Appointment) error
	GetAllItemByLastKey(lastKey string, limit int32) ([]*Appointment, error)
	GetAppointmentByAppId(appId string) (*Appointment, error)
}
