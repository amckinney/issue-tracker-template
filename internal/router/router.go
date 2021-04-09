package router

import (
	"github.com/amckinney/issue-tracker/internal/handler/issue"
	"github.com/go-chi/chi"
)

// New constructs a new router for the Issue Tracker application.
func New(issueHandler issue.Handler) *chi.Mux {
	router := chi.NewRouter()
	router.Route("/issues", func(router chi.Router) {
		router.Post("/", issueHandler.CreateIssue) // POST /issues

		router.Route("/{issueID}", func(router chi.Router) {
			router.Get("/", issueHandler.GetIssue)       // GET /issues/63d4876c-9d85-4c07-a28b-3c1c56a93153
			router.Put("/", issueHandler.UpdateIssue)    // PUT /issues/63d4876c-9d85-4c07-a28b-3c1c56a93153
			router.Delete("/", issueHandler.DeleteIssue) // DELETE /issues/63d4876c-9d85-4c07-a28b-3c1c56a93153
		})
	})
	return router
}
