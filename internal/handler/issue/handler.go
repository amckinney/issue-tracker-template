package issue

import (
	"encoding/json"
	"net/http"

	"github.com/amckinney/issue-tracker/api"
	"github.com/amckinney/issue-tracker/internal/controller/issue"
	"github.com/amckinney/issue-tracker/internal/entity"
	"github.com/amckinney/issue-tracker/internal/mapper"
	"github.com/amckinney/issue-tracker/internal/persist"
	"go.uber.org/zap"
)

// Handler defines the issue endpoints.
type Handler interface {
	CreateIssue(http.ResponseWriter, *http.Request)
	GetIssue(http.ResponseWriter, *http.Request)
	UpdateIssue(http.ResponseWriter, *http.Request)
	DeleteIssue(http.ResponseWriter, *http.Request)
}

type handler struct {
	logger     *zap.Logger
	store      persist.Store
	controller issue.Controller
}

// New returns a new Handler.
func New(logger *zap.Logger, store persist.Store, controller issue.Controller) Handler {
	return newHandler(logger, store, controller)
}

func newHandler(logger *zap.Logger, store persist.Store, controller issue.Controller) *handler {
	return &handler{
		logger:     logger,
		store:      store,
		controller: controller,
	}
}

// CreateIssue creates an issue according to the given request.
func (h *handler) CreateIssue(writer http.ResponseWriter, request *http.Request) {
	issue, err := issueFromRequest(request)
	if err != nil {
		http.Error(writer, err.Error(), http.StatusBadRequest)
	}
	// Create the transaction.
	created, err := h.controller.CreateIssue(request.Context(), issue)
	if err != nil {
		http.Error(writer, err.Error(), http.StatusInternalServerError)
		return
	}
	if err := writeIssue(writer, created); err != nil {
		http.Error(writer, err.Error(), http.StatusInternalServerError)
		return
	}
	return
}

// GetIssue gets an issue according to the given request.
func (h *handler) GetIssue(writer http.ResponseWriter, request *http.Request) {
	// TODO
	return
}

// UpdateIssue updates an issue according to the given request.
func (h *handler) UpdateIssue(writer http.ResponseWriter, request *http.Request) {
	// TODO
	return
}

// DeleteIssue deletes an issue according to the given request.
func (h *handler) DeleteIssue(writer http.ResponseWriter, request *http.Request) {
	// TODO
	return
}

// issueFromRequest maps the given *http.Request into
// an *entity.Issue.
func issueFromRequest(request *http.Request) (*entity.Issue, error) {
	var apiIssue *api.Issue
	if err := json.NewDecoder(request.Body).Decode(apiIssue); err != nil {
		return nil, err
	}
	return mapper.IssueFromAPI(apiIssue)
}

// writeIssue writes the given *entity.Issue to the given http.ResponseWriter.
func writeIssue(writer http.ResponseWriter, issue *entity.Issue) error {
	apiIssue, err := mapper.IssueToAPI(issue)
	if err != nil {
		return err
	}
	if err := json.NewEncoder(writer).Encode(apiIssue); err != nil {
		return err
	}
	writer.Header().Set("Content-Type", "application/json")
	writer.WriteHeader(http.StatusOK)
	return nil
}
