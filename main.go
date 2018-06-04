package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	. "github.com/marcelocollyer/resume_go_app/controller"
)

var resumeController = ResumeController{}

// Main function
func main() {
	// Init router
	r := mux.NewRouter()

	//set endpoints

	resumeController.InitEndPoints(r)

	// Start server
	if err := http.ListenAndServe(":8000", r); err != nil {
		log.Fatal(err)
	}
}
