package issue

import (
	"encoding/json"
	"net/http"

	"github.com/amckinney/issue-tracker/api"
	"github.com/amckinney/issue-tracker/internal/controller/issue"
	"github.com/amckinney/issue-tracker/internal/entity"
	"github.com/amckinney/issue-tracker/internal/mapper"
	"github.com/amckinney/issue-tracker/internal/store"
	"github.com/go-chi/chi"
	"github.com/gofrs/uuid"
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
	store      store.Store
	controller issue.Controller
}

// New returns a new Handler.
func New(logger *zap.Logger, store store.Store, controller issue.Controller) Handler {
	return newHandler(logger, store, controller)
}

func newHandler(logger *zap.Logger, store store.Store, controller issue.Controller) *handler {
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
		return
	}
	var createdIssue *entity.Issue
	if err := h.store.ReadWrite(
		request.Context(),
		func(tx store.ReadWriteTx) error {
			createdIssue, err = h.controller.CreateIssue(tx, issue)
			return err
		},
	); err != nil {
		// TODO: We need to correctly handle error codes here.
		//       Not all errors should be handled as internal
		//       server errors.
		http.Error(writer, err.Error(), http.StatusInternalServerError)
		return
	}
	if err := writeIssue(writer, createdIssue); err != nil {
		http.Error(writer, err.Error(), http.StatusInternalServerError)
		return
	}
	return
}

// GetIssue gets an issue according to the given request.
func (h *handler) GetIssue(writer http.ResponseWriter, request *http.Request) {
	issueID, err := issueIDFromRequest(request)
	if err != nil {
		http.Error(writer, err.Error(), http.StatusBadRequest)
		return
	}
	var gotIssue *entity.Issue
	if err := h.store.ReadWrite(
		request.Context(),
		func(tx store.ReadWriteTx) error {
			gotIssue, err = h.controller.GetIssue(tx, issueID)
			return err
		},
	); err != nil {
		// TODO: We need to correctly handle error codes here.
		//       Not all errors should be handled as internal
		//       server errors.
		http.Error(writer, err.Error(), http.StatusInternalServerError)
		return
	}
	if err := writeIssue(writer, gotIssue); err != nil {
		http.Error(writer, err.Error(), http.StatusInternalServerError)
		return
	}
	return
}

// UpdateIssue updates an issue according to the given request.
func (h *handler) UpdateIssue(writer http.ResponseWriter, request *http.Request) {
	// Grab the issueID from the query param and validate that it
	// matches the issue contained in the body.
	issue, err := issueFromRequest(request)
	if err != nil {
		http.Error(writer, err.Error(), http.StatusBadRequest)
		return
	}
	var updatedIssue *entity.Issue
	if err := h.store.ReadWrite(
		request.Context(),
		func(tx store.ReadWriteTx) error {
			updatedIssue, err = h.controller.UpdateIssue(tx, issue)
			return err
		},
	); err != nil {
		// TODO: We need to correctly handle error codes here.
		//       Not all errors should be handled as internal
		//       server errors.
		http.Error(writer, err.Error(), http.StatusInternalServerError)
		return
	}
	if err := writeIssue(writer, updatedIssue); err != nil {
		http.Error(writer, err.Error(), http.StatusInternalServerError)
		return
	}
	return
}

// DeleteIssue deletes an issue according to the given request.
func (h *handler) DeleteIssue(writer http.ResponseWriter, request *http.Request) {
	issueID, err := issueIDFromRequest(request)
	if err != nil {
		http.Error(writer, err.Error(), http.StatusBadRequest)
		return
	}
	var gotIssue *entity.Issue
	if err := h.store.ReadWrite(
		request.Context(),
		func(tx store.ReadWriteTx) error {
			gotIssue, err = h.controller.DeleteIssue(tx, issueID)
			return err
		},
	); err != nil {
		// TODO: We need to correctly handle error codes here.
		//       Not all errors should be handled as internal
		//       server errors.
		http.Error(writer, err.Error(), http.StatusInternalServerError)
		return
	}
	if err := writeIssue(writer, gotIssue); err != nil {
		http.Error(writer, err.Error(), http.StatusInternalServerError)
		return
	}
	return
}

// issueFromRequest maps the given *http.Request into
// an *entity.Issue.
func issueFromRequest(request *http.Request) (*entity.Issue, error) {
	var apiIssue api.Issue
	if err := json.NewDecoder(request.Body).Decode(&apiIssue); err != nil {
		return nil, err
	}
	return mapper.IssueFromAPI(&apiIssue)
}

// issueIDFromRequest maps the given *http.Request into
// an issueID, represented as a UUID.
func issueIDFromRequest(request *http.Request) (uuid.UUID, error) {
	issueID := chi.URLParam(request, "issueID")
	return uuid.FromString(issueID)
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
	return nil
}
