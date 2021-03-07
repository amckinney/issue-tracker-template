package issues

import (
	"net/http"

	"go.uber.org/zap"
)

// Handler defines the issues endpoints.
type Handler interface {
	CreateIssue(http.ResponseWriter, *http.Request)
	GetIssue(http.ResponseWriter, *http.Request)
	UpdateIssue(http.ResponseWriter, *http.Request)
	DeleteIssue(http.ResponseWriter, *http.Request)
}

type handler struct {
	logger *zap.Logger
}

// New returns a new Handler.
func New(logger *zap.Logger) Handler {
	return newHandler(logger)
}

func newHandler(logger *zap.Logger) *handler {
	return &handler{
		logger: logger,
	}
}

// CreateIssue creates an issue according to the given request.
func (h *handler) CreateIssue(writer http.ResponseWriter, request *http.Request) {
	// TODO
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
