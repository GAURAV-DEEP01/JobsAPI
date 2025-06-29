package routes

import (
	"github.com/gaurav-deep01/jobboard-api/internal/handler"
	"github.com/go-chi/chi/v5"
)

func Jobs(r chi.Router) {
	r.Route("/", func(r chi.Router) {
		r.Get("/", handler.AllJobs)
		r.Post("/", handler.AddJob)
		r.Get("/{id}", handler.CompanyJob)
		r.Delete("/{id}", handler.RemoveJob)
	})
}
