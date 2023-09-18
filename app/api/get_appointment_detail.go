package api

import (
	"fmt"
	"net/http"
	"robinhood/app/api/models"
	"robinhood/repo"
	"robinhood/usecase"

	"github.com/gin-gonic/gin"
)

type AppointmentDetailInput struct {
	AppId string `form:"appId"`
}

func GetAppointmentDetail(c *gin.Context) {
	input := new(AppointmentDetailInput)
	if err := c.ShouldBindQuery(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	fmt.Println("input.AppId==>", input.AppId)

	appointmentRepo := repo.NewAppointmentRepoDdb("appointment")
	commentRepo := repo.NewCommentRepoDdb("appointment")
	userProfileRepo := repo.NewUserProfileRepoDdb("user_profile")
	newUsecase := usecase.New(appointmentRepo, commentRepo, userProfileRepo)

	getAppointmentDetailOutput, err := newUsecase.GetAppointmentDetail(input.AppId)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error GetAppointmentList": err.Error()})
		return
	}
	output := models.NewAppointmentDetail(getAppointmentDetailOutput)

	c.IndentedJSON(http.StatusOK, output)
}
