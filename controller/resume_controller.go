package controller

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
	. "github.com/marcelocollyer/resume_go_app/dao"
	. "github.com/marcelocollyer/resume_go_app/model"
	"gopkg.in/mgo.v2/bson"
)

// ResumeController to expose endpoints
type ResumeController struct {
}

var dao = ResumeDAO{}

func (m *ResumeController) InitEndPoints(router *mux.Router) {
	// Route handles & endpoints
	router.HandleFunc("/resumes", GetResumes).Methods("GET")
	router.HandleFunc("/resumes", createResume).Methods("POST")
	router.HandleFunc("/resumes/{id}", getResume).Methods("GET")
}

// Get all resumes
func GetResumes(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	resumes, err := dao.FindAll()
	if err != nil {
		RespondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}
	RespondWithJson(w, http.StatusOK, resumes)
}

// GET resume by id
func getResume(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	resume, err := dao.FindById(params["id"])
	if err != nil {
		RespondWithError(w, http.StatusBadRequest, "Invalid Resume ID")
		return
	}
	RespondWithJson(w, http.StatusOK, resume)
}

// POST a new resume
func createResume(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	var resume Resume
	if err := json.NewDecoder(r.Body).Decode(&resume); err != nil {
		RespondWithError(w, http.StatusBadRequest, "Invalid request payload")
		return
	}
	resume.ID = bson.NewObjectId()
	if err := dao.Insert(resume); err != nil {
		RespondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}
	RespondWithJson(w, http.StatusCreated, resume)
}

func RespondWithJson(w http.ResponseWriter, code int, payload interface{}) {
	response, _ := json.Marshal(payload)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	w.Write(response)
}

func RespondWithError(w http.ResponseWriter, code int, msg string) {
	RespondWithJson(w, code, map[string]string{"error": msg})
}
