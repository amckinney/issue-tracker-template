package router

import (
	"github.com/amckinney/issue-tracker/internal/handler/issues"
	"github.com/go-chi/chi"
)

// New constructs a new router for the
// Issue Tracker application.
func New(handler issues.Handler) *chi.Mux {
	router := chi.NewRouter()
	router.Route("/issues", func(router chi.Router) {
		router.Post("/", handler.CreateIssue) // POST /issues

		router.Route("/{issueID}", func(router chi.Router) {
			router.Get("/", handler.GetIssue)       // GET /issues/63d4876c-9d85-4c07-a28b-3c1c56a93153
			router.Put("/", handler.UpdateIssue)    // PUT /issues/63d4876c-9d85-4c07-a28b-3c1c56a93153
			router.Delete("/", handler.DeleteIssue) // DELETE /issues/63d4876c-9d85-4c07-a28b-3c1c56a93153
		})
	})
	return router
}
