package main

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	. "github.com/marcelocollyer/resume_go_app/config"
	. "github.com/marcelocollyer/resume_go_app/dao"
	. "github.com/marcelocollyer/resume_go_app/model"
	"gopkg.in/mgo.v2/bson"
)

// Init resumes var as a slice Resume struct
var resumes []Resume
var config = Config{}
var dao = ResumeDAO{}

// Get all resumes
func getResumes(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	resumes, err := dao.FindAll()
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}
	respondWithJson(w, http.StatusOK, resumes)
}

// GET resume by id
func getResume(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	resume, err := dao.FindById(params["id"])
	if err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid Resume ID")
		return
	}
	respondWithJson(w, http.StatusOK, resume)
}

// POST a new resume
func createResume(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	var resume Resume
	if err := json.NewDecoder(r.Body).Decode(&resume); err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid request payload")
		return
	}
	resume.ID = bson.NewObjectId()
	if err := dao.Insert(resume); err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}
	respondWithJson(w, http.StatusCreated, resume)
}

func respondWithJson(w http.ResponseWriter, code int, payload interface{}) {
	response, _ := json.Marshal(payload)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	w.Write(response)
}

func respondWithError(w http.ResponseWriter, code int, msg string) {
	respondWithJson(w, code, map[string]string{"error": msg})
}

// Parse the configuration file 'config.toml', and establish a connection to DB
func init() {
	config.Read()

	dao.Server = config.Server
	dao.Database = config.Database
	dao.Connect()
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
	r.HandleFunc("/resumes", createResume).Methods("POST")
	r.HandleFunc("/resumes/{id}", getResume).Methods("GET")

	// Start server
	if err := http.ListenAndServe(":8000", r); err != nil {
		log.Fatal(err)
	}
}
