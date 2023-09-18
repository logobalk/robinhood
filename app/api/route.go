package api

import "github.com/gin-gonic/gin"

func New() *gin.Engine {
	router := gin.Default()
	router.GET("/appointment/list", GetAppointmentList)
	router.POST("/appointment/save", SaveAppointment)
	router.GET("/appointment/detail", GetAppointmentDetail)
	router.POST("/appointment/comment/save", SaveComment)
	router.POST("/userprofile/save", SaveUserProfile)
	router.GET("/userprofile/detail", GetUserProfile)
	router.GET("/master-data/status", MasterData)

	return router
}
