package api

import (
	"fmt"
	"net/http"
	"robinhood/domain"
	"robinhood/usecase"

	"github.com/gin-gonic/gin"
)

type SaveAppointmentInput struct {
	AppId         *string       `json:"appId"`
	Title         string        `json:"title"`
	Description   string        `json:"description"`
	CreatedBy     string        `json:"createdBy"`
	Email         string        `json:"email"`
	UserReference string        `json:"userReference"`
	Status        domain.Status `json:"status"`
	IsActive      bool          `json:"IsActive"`
}

func (a *ApiInput) SaveAppointment(c *gin.Context) {
	input := new(SaveAppointmentInput)
	if err := c.ShouldBindJSON(&input); err != nil {
		c.IndentedJSON(http.StatusBadRequest, "body failed")
		// c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	newUsecase := a.useCase

	err := newUsecase.SaveAppointment(&usecase.SaveAppointmentInput{
		AppId:         input.AppId,
		Title:         input.Title,
		Description:   input.Description,
		CreatedBy:     input.CreatedBy,
		Email:         input.Email,
		UserReference: input.UserReference,
		Status:        input.Status,
		IsActive:      input.IsActive,
	})
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error SaveAppointment": err.Error()})
		return
	}
	fmt.Println("err==>", err)

	c.IndentedJSON(http.StatusOK, err)
}
