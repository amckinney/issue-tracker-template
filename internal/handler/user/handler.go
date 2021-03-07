package user

import (
	"net/http"

	"go.uber.org/zap"
)

// Handler defines the user endpoints.
type Handler interface {
	CreateUser(http.ResponseWriter, *http.Request)
	GetUser(http.ResponseWriter, *http.Request)
	UpdateUser(http.ResponseWriter, *http.Request)
	DeleteUser(http.ResponseWriter, *http.Request)
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

// CreateUser creates an user according to the given request.
func (h *handler) CreateUser(writer http.ResponseWriter, request *http.Request) {
	// TODO
	return
}

// GetUser gets an user according to the given request.
func (h *handler) GetUser(writer http.ResponseWriter, request *http.Request) {
	// TODO
	return
}

// UpdateUser updates an user according to the given request.
func (h *handler) UpdateUser(writer http.ResponseWriter, request *http.Request) {
	// TODO
	return
}

// DeleteUser deletes an user according to the given request.
func (h *handler) DeleteUser(writer http.ResponseWriter, request *http.Request) {
	// TODO
	return
}
