package handler

import (
	"encoding/json"
	"github.com/gaurav-deep01/jobboard-api/internal/db"
	"github.com/go-chi/chi/v5"
	"net/http"
)

func AllJobs(w http.ResponseWriter, r *http.Request) {
	jobs, err := db.GetAllJobs()
	if err != nil {
		http.Error(w, "Failed to get jobs", http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(jobs)
}

func AddJob(w http.ResponseWriter, r *http.Request) {
	var job db.Job
	if err := json.NewDecoder(r.Body).Decode(&job); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}
	if err := db.AddJob(job); err != nil {
		http.Error(w, "Failed to add job", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(map[string]string{"status": "job created"})
}

func CompanyJob(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")

	job, err := db.GetJobByID(id)
	if err != nil {
		http.Error(w, "Job not found", http.StatusNotFound)
		return
	}

	json.NewEncoder(w).Encode(job)
}

func RemoveJob(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")

	if err := db.DeleteJob(id); err != nil {
		http.Error(w, "Failed to delete job", http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(map[string]string{"status": "job deleted"})
}
