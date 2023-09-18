package api

import (
	"net/http"
	"robinhood/app/api/models"

	"github.com/gin-gonic/gin"
)

type UserProfileInput struct {
	Id string `form:"id"`
}

func (a *ApiInput) GetUserProfile(c *gin.Context) {
	input := new(UserProfileInput)
	if err := c.ShouldBindQuery(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	newUsecase := a.useCase

	getUserProfileOutput, err := newUsecase.GetUserProfileById(input.Id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error getUserProfileOutput": err.Error()})
		return
	}
	output := models.NewUserProfile(getUserProfileOutput)

	c.IndentedJSON(http.StatusOK, output)
}
