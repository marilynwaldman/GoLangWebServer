package main

import (
	// Standard library packages
	"net/http"

	// Third party packages
	"github.com/julienschmidt/httprouter"
	"github.com/user/webapp/controllers"
)

func main() {
	// Instantiate a new router
	r := httprouter.New()

	// Get a  DogparkController instance
	uc := controllers.NewDogparkController()

	// Get a  dogpark resource
	r.GET("/dogpark/", uc.GetDogparks)

	// Get a  dogpark resource
	r.GET("/dogpark/:id", uc.GetDogpark)

	r.POST("/dogpark", uc.CreateDogpark)

	r.DELETE("/dogpark/:id", uc.RemoveDogpark)

	// Fire up the server
	http.ListenAndServe("localhost:3000", r)
}