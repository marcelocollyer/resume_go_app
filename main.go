package main

import (
	"log"
	"net/http"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	. "github.com/marcelocollyer/resume_go_app/controller"
	_ "github.com/qodrorid/godaemon"
)

var resumeController = ResumeController{}

// Main function
func main() {
	// Init router
	r := mux.NewRouter()

	//set endpoints

	resumeController.InitEndPoints(r)

	corsObj := handlers.AllowedOrigins([]string{"*"})

	// Start server
	if err := http.ListenAndServe(":8000", handlers.CORS(corsObj)(r)); err != nil {
		log.Fatal(err)
	}
}
