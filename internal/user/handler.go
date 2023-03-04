package user

import (
	"fmt"
	"github.com/julienschmidt/httprouter"
	"net/http"
	"todo/internal/apperror"
	"todo/internal/handlers"
	"todo/pkg/logging"
)

const (
	usersURl = "/users"
	userURl  = "/users/:uuid"
)

type handler struct {
	logger *logging.Logger
}

func NewHandler(logger *logging.Logger) handlers.Handler {
	return &handler{
		logger: logger,
	}
}

func (h *handler) Register(router *httprouter.Router) {
	router.HandlerFunc(http.MethodGet, usersURl, apperror.Middleware(h.GetList))
	router.HandlerFunc(http.MethodGet, userURl, apperror.Middleware(h.GetUserByUUId))
	router.HandlerFunc(http.MethodPost, usersURl, apperror.Middleware(h.CreateUser))
	router.HandlerFunc(http.MethodPut, userURl, apperror.Middleware(h.UpdateUser))
	router.HandlerFunc(http.MethodPatch, userURl, apperror.Middleware(h.PartiallyUpdateUser))
	router.HandlerFunc(http.MethodDelete, userURl, apperror.Middleware(h.DeleteUser))
}

func (h *handler) GetList(w http.ResponseWriter, r *http.Request) error {
	return apperror.ErrNotFound
}

func (h *handler) GetUserByUUId(w http.ResponseWriter, r *http.Request) error {
	return apperror.NewAppError(nil, "test", "test", "123")
}
func (h *handler) CreateUser(w http.ResponseWriter, r *http.Request) error {
	return fmt.Errorf("this ia  Api error")
}
func (h *handler) UpdateUser(w http.ResponseWriter, r *http.Request) error {
	w.Write([]byte("this is update user"))
	return nil
}
func (h *handler) PartiallyUpdateUser(w http.ResponseWriter, r *http.Request) error {
	w.Write([]byte("this is partially update user"))
	return nil

}
func (h *handler) DeleteUser(w http.ResponseWriter, r *http.Request) error {
	w.Write([]byte("this is delete user"))
	return nil
}
