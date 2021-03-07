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
			router.Get("/", handler.GetIssue)       // GET /issues/123
			router.Put("/", handler.UpdateIssue)    // PUT /issues/123
			router.Delete("/", handler.DeleteIssue) // DELETE /issues/123
		})
	})
	return router
}
