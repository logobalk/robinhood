package api

import (
	"net/http"
	"robinhood/app/api/models"
	"robinhood/repo"
	"robinhood/usecase"

	"github.com/gin-gonic/gin"
)

type UserProfileInput struct {
	Id string `form:"id"`
}

func GetUserProfile(c *gin.Context) {
	input := new(UserProfileInput)
	if err := c.ShouldBindQuery(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	appointmentRepo := repo.NewAppointmentRepoDdb("appointment")
	commentRepo := repo.NewCommentRepoDdb("appointment")
	userProfileRepo := repo.NewUserProfileRepoDdb("user_profile")
	newUsecase := usecase.New(appointmentRepo, commentRepo, userProfileRepo)

	getUserProfileOutput, err := newUsecase.GetUserProfileById(input.Id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error getUserProfileOutput": err.Error()})
		return
	}
	output := models.NewUserProfile(getUserProfileOutput)

	c.IndentedJSON(http.StatusOK, output)
}
