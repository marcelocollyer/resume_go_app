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

// InitEndPoints initializes endpoints
func (m *ResumeController) InitEndPoints(router *mux.Router) {

	router.HandleFunc("/resumes", GetResumes).Methods("GET")
	router.HandleFunc("/resumes", CreateResume).Methods("POST")
	router.HandleFunc("/resumes", UpdateResumeEndPoint).Methods("PUT")
	router.HandleFunc("/resumes", DeleteResumeEndPoint).Methods("DELETE")
	router.HandleFunc("/resumes/{id}", DeleteResumeByIDEndPoint).Methods("DELETE")
	router.HandleFunc("/resumes/{id}", GetResume).Methods("GET")
}

// GetResumes - get all resumes
func GetResumes(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	resumes, err := dao.FindAll()
	if err != nil {
		RespondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}
	RespondWithJSON(w, http.StatusOK, resumes)
}

// GetResume - get resume by id
func GetResume(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	resume, err := dao.FindByID(params["id"])
	if err != nil {
		RespondWithError(w, http.StatusBadRequest, "Invalid Resume ID")
		return
	}
	RespondWithJSON(w, http.StatusOK, resume)
}

// CreateResume - creates a new resume entry
func CreateResume(w http.ResponseWriter, r *http.Request) {
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
	RespondWithJSON(w, http.StatusCreated, resume)
}

// UpdateResumeEndPoint - updates an existing resume entry
func UpdateResumeEndPoint(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	var resume Resume
	if err := json.NewDecoder(r.Body).Decode(&resume); err != nil {
		RespondWithError(w, http.StatusBadRequest, "Invalid request payload")
		return
	}
	if err := dao.Update(resume); err != nil {
		RespondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}
	RespondWithJSON(w, http.StatusOK, map[string]string{"result": "success"})
}

// DeleteResumeEndPoint - Deletes a resume entry
func DeleteResumeEndPoint(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	var resume Resume
	if err := json.NewDecoder(r.Body).Decode(&resume); err != nil {
		RespondWithError(w, http.StatusBadRequest, "Invalid request payload")
		return
	}
	if err := dao.Delete(resume); err != nil {
		RespondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}
	RespondWithJSON(w, http.StatusOK, map[string]string{"result": "success"})
}

// DeleteResumeEndPoint - Deletes a resume entry by given ID
func DeleteResumeByIDEndPoint(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	params := mux.Vars(r)

	if err := dao.DeleteByID(params["id"]); err != nil {
		RespondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}
	RespondWithJSON(w, http.StatusOK, map[string]string{"result": "success"})
}

// RespondWithJSON - writes JSON format response
func RespondWithJSON(w http.ResponseWriter, code int, payload interface{}) {
	response, _ := json.Marshal(payload)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	w.Write(response)
}

// RespondWithError - handles errors
func RespondWithError(w http.ResponseWriter, code int, msg string) {
	RespondWithJSON(w, code, map[string]string{"error": msg})
}
