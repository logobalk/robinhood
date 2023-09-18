package api

import (
	"fmt"
	"net/http"
	"robinhood/usecase"

	"github.com/gin-gonic/gin"
)

type SaveUserProfileInput struct {
	UserName  string `json:"userName"`
	Name      string `json:"name"`
	CreatedBy string `json:"createdBy"`
	Email     string `json:"email"`
}

func (a *ApiInput) SaveUserProfile(c *gin.Context) {
	input := new(SaveUserProfileInput)
	if err := c.ShouldBindJSON(&input); err != nil {
		c.IndentedJSON(http.StatusBadRequest, "body failed")
		return
	}

	newUsecase := a.useCase

	id, err := newUsecase.SaveUserProfile(&usecase.UserProfileInput{
		UserName:  input.UserName,
		Name:      input.Name,
		CreatedBy: input.CreatedBy,
		Email:     input.Email,
	})
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error CreateUserProfile": err.Error()})
		return
	}
	fmt.Println("err==>", err)

	c.IndentedJSON(http.StatusOK, id)
}
