package main

import (
	"context"
	"robinhood/app/api"
	"robinhood/repo"
	dynamoddb "robinhood/utils/dynamoDdb"
)

func main() {
	router := api.New()

	dynamoddb.DefaultClient = dynamoddb.NewLocal()
	// newRepo := repo.NewAppointmentRepoDdb("appointment")
	schema := repo.AppointmentRepoDdbSchema()
	userprofileSchema := repo.UserProfileRepoDdbSchema()
	dynamoddb.MustDeleteTable(context.Background(), schema)
	dynamoddb.MustDeleteTable(context.Background(), userprofileSchema)
	dynamoddb.MustCreateTable(context.Background(), schema)
	dynamoddb.MustCreateTable(context.Background(), userprofileSchema)

	router.Run(":8081")

}
