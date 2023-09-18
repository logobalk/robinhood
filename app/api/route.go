package api

import (
	"robinhood/usecase"

	"github.com/gin-gonic/gin"
)

func New(useCase usecase.UseCase) *gin.Engine {
	router := gin.Default()
	newApi := NewApi(useCase)

	router.GET("/appointment/list", newApi.GetAppointmentList)
	router.POST("/appointment/save", newApi.SaveAppointment)
	router.GET("/appointment/detail", newApi.GetAppointmentDetail)
	router.POST("/appointment/comment/save", newApi.SaveComment)
	router.POST("/userprofile/save", newApi.SaveUserProfile)
	router.GET("/userprofile/detail", newApi.GetUserProfile)
	router.GET("/master-data/status", newApi.MasterData)

	return router
}
