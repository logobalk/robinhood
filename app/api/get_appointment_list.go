package api

import (
	"net/http"
	"robinhood/app/api/models"
	"robinhood/repo"
	"robinhood/usecase"

	"github.com/aws/smithy-go/ptr"
	"github.com/gin-gonic/gin"
)

type AppointmentListInput struct {
	Lastkey *string `form:"lastKey"`
	Limit   *int32  `form:"limit"`
	Status  *string `form:"status"`
}

func GetAppointmentList(c *gin.Context) {
	input := new(AppointmentListInput)
	if err := c.ShouldBindQuery(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	appointmentRepo := repo.NewAppointmentRepoDdb("appointment")
	commentRepo := repo.NewCommentRepoDdb("appointment")
	userProfileRepo := repo.NewUserProfileRepoDdb("user_profile")
	newUsecase := usecase.New(appointmentRepo, commentRepo, userProfileRepo)

	var ensureLastKey *string
	ensureLastKey = input.Lastkey
	if input.Lastkey == nil {
		ensureLastKey = ptr.String("0")
	}
	var ensureLimit *int32
	ensureLimit = input.Limit
	if input.Limit == nil {
		ensureLimit = ptr.Int32(0)
	}

	getAppointmentOutput, err := newUsecase.GetAppointmentList(*ensureLastKey, *ensureLimit)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error GetAppointmentList": err.Error()})
		return
	}
	output := models.NewAppointmentList(getAppointmentOutput)

	c.IndentedJSON(http.StatusOK, output)
}
