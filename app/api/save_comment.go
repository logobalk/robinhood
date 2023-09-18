package api

import (
	"fmt"
	"net/http"
	"robinhood/usecase"

	"github.com/gin-gonic/gin"
)

type SaveCommentInput struct {
	Id            *string `json:"id"`
	AppId         string  `json:"appId"`
	Message       string  `json:"message"`
	CreatedBy     *string `json:"createdBy"`
	UpdatedBy     *string `json:"updatedBy"`
	UserReference string  `json:"userReference"`
}

func (a *ApiInput) SaveComment(c *gin.Context) {
	input := new(SaveCommentInput)
	if err := c.ShouldBindJSON(&input); err != nil {
		c.IndentedJSON(http.StatusBadRequest, "body failed")
		return
	}

	newUsecase := a.useCase

	err := newUsecase.SaveAppointmentDetail(&usecase.CreateAppointmentDetailInput{
		Id:            input.Id,
		AppId:         input.AppId,
		Message:       input.Message,
		CreatedBy:     input.CreatedBy,
		UpdatedBy:     input.UpdatedBy,
		UserReference: input.UserReference,
	})
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error SaveAppointmentDetail": err.Error()})
		return
	}
	fmt.Println("err==>", err)

	c.IndentedJSON(http.StatusOK, err)
}
