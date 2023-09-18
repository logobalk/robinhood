package api

import (
	"net/http"
	"robinhood/domain"

	"github.com/gin-gonic/gin"
)

type Status struct {
	ID    int           `json:"id"`
	Label string        `json:"label"`
	Value domain.Status `json:"value"`
}

func MasterData(c *gin.Context) {

	status := []Status{
		{ID: 1, Label: "To Do", Value: domain.StatusTodo},
		{ID: 2, Label: "In Progress", Value: domain.StatusInProgress},
		{ID: 3, Label: "Done", Value: domain.StatusDone},
	}

	c.JSON(http.StatusOK, status)
}
