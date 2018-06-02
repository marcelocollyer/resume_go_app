package main

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	. "github.com/marcelocollyer/resume_go_app/model"
)

// Init resumes var as a slice Resume struct
var resumes []Resume

// Get all resumes
func getResumes(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(resumes)
}

// Get resume by id
func getResume(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r) // Gets params
	// Loop through resumes and find one with the id from the params
	for _, item := range resumes {
		if item.ID == params["id"] {
			json.NewEncoder(w).Encode(item)
			return
		}
	}
	err := json.NewEncoder(w).Encode(&Resume{})
	if err != nil {
		log.Println("Error: ")
	}
}

// Main function
func main() {
	// Init router
	r := mux.NewRouter()

	// Hardcoded data - @todo: add it to mongodb
	resumes = append(resumes,
		Resume{
			ID:         "1",
			Name:       "Marcelo Collyer",
			CareerDesc: "+11 years Senior Full Stack Developer always aiming to learn new technologies to provide clients greatest solutions.",
			RolesDesc:  "Software Architect, Technical Leader, Senior Software Developer",
			Experiences: []Experience{
				{Title: "Full Stack Developer"},
				{Title: "Analyst"},
			}})

	// Route handles & endpoints
	r.HandleFunc("/resumes", getResumes).Methods("GET")
	r.HandleFunc("/resumes/{id}", getResume).Methods("GET")

	// Start server
	if err := http.ListenAndServe(":8000", r); err != nil {
		log.Fatal(err)
	}
}
