package api

import (
	"robinhood/usecase"

	"github.com/gin-gonic/gin"
)

type ApiInput struct {
	useCase usecase.UseCase
}

type Api interface {
	GetAppointmentDetail(c *gin.Context)
	GetAppointmentList(c *gin.Context)
	GetUserProfile(c *gin.Context)
	MasterData(c *gin.Context)
	SaveAppointment(c *gin.Context)
	SaveComment(c *gin.Context)
	SaveUserProfile(c *gin.Context)
}

func NewApi(useCase usecase.UseCase) *ApiInput {
	return &ApiInput{
		useCase: useCase,
	}
}
