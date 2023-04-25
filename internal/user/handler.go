package user

import (
	"net/http"

	"github.com/RAprogramm/rest-api-tutorial/internal/handlers"
	"github.com/RAprogramm/rest-api-tutorial/pkg/logging"
	"github.com/julienschmidt/httprouter"
)

var _ handlers.Handler = &handler{}

const (
	usersURL = "/users"
	userURL  = "/users/:uuid"
)

type handler struct {
	logger logging.Logger
}

func NewHandler(logger logging.Logger) handlers.Handler {
	return &handler{
		logger: logger,
	}
}

func (h *handler) Register(router *httprouter.Router) {
	router.HandlerFunc(http.MethodGet, usersURL, h.GetList)
	router.HandlerFunc(http.MethodGet, userURL, h.GetUserByUUID)
	router.HandlerFunc(http.MethodPost, usersURL, h.CreateUser)
	router.HandlerFunc(http.MethodPut, userURL, h.UpdateUser)
	router.HandlerFunc(http.MethodPatch, userURL, h.PartiallyUpdateUser)
	router.HandlerFunc(http.MethodDelete, userURL, h.DeleteUser)
}

func (h *handler) GetList(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(200)
	w.Write([]byte("This is the list of users"))
}

func (h *handler) GetUserByUUID(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(200)
	w.Write([]byte("This is get user by uuid"))
}

func (h *handler) CreateUser(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(201)
	w.Write([]byte("This is create user"))
}

func (h *handler) UpdateUser(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(204)
	w.Write([]byte("This is update"))
}

func (h *handler) PartiallyUpdateUser(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(204)
	w.Write([]byte("This is partially update user"))
}

func (h *handler) DeleteUser(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(204)
	w.Write([]byte("This is delete user"))
}
