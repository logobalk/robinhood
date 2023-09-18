package main

import (
	"context"
	"robinhood/app/api"
	"robinhood/repo"
	"robinhood/usecase"
	dynamoddb "robinhood/utils/dynamoDdb"
)

func main() {

	dynamoddb.DefaultClient = dynamoddb.NewLocal()
	// newRepo := repo.NewAppointmentRepoDdb("appointment")
	schema := repo.AppointmentRepoDdbSchema()
	userprofileSchema := repo.UserProfileRepoDdbSchema()
	dynamoddb.MustDeleteTable(context.Background(), schema)
	dynamoddb.MustDeleteTable(context.Background(), userprofileSchema)
	dynamoddb.MustCreateTable(context.Background(), schema)
	dynamoddb.MustCreateTable(context.Background(), userprofileSchema)

	appointmentRepo := repo.NewAppointmentRepoDdb("appointment")
	commentRepo := repo.NewCommentRepoDdb("appointment")
	userProfileRepo := repo.NewUserProfileRepoDdb("user_profile")
	newUsecase := usecase.New(appointmentRepo, commentRepo, userProfileRepo)
	router := api.New(newUsecase)

	router.Run(":8081")

}
