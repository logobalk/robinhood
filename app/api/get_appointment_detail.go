package api

import (
	"net/http"
	"robinhood/app/api/models"

	"github.com/gin-gonic/gin"
)

type AppointmentDetailInput struct {
	AppId string `form:"appId"`
}

func (a *ApiInput) GetAppointmentDetail(c *gin.Context) {
	input := new(AppointmentDetailInput)
	if err := c.ShouldBindQuery(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	newUsecase := a.useCase

	getAppointmentDetailOutput, err := newUsecase.GetAppointmentDetail(input.AppId)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error GetAppointmentList": err.Error()})
		return
	}
	output := models.NewAppointmentDetail(getAppointmentDetailOutput)

	c.IndentedJSON(http.StatusOK, output)
}
