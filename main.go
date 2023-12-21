// main.go
package main

import (
	"nethgateapi/db"
	"nethgateapi/handlers"
	"nethgateapi/nethgateapi"

	"github.com/gin-gonic/gin"
)

func main() {
	db, err := db.SetupDatabase() // Your function to set up the database
	if err != nil {
		panic(err)
	}
	svc := nethgateapi.NewNethGateService(db)

	userHandler := handlers.NewUserHandler(svc)

	r := gin.Default()

	// Set up your routes and handlers
	r.POST("/users", userHandler.CreateUser)
	// Set up other routes...

	r.Run(":8080")
}
