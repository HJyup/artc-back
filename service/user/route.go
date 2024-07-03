package user

import (
	"github.com/gorilla/mux"
	"net/http"
)

type Handler struct {
	store any
}

func NewHandler(store any) *Handler {
	return &Handler{store: store}
}

func (handler *Handler) RegisterRouters(router *mux.Router) {
	router.HandleFunc("/login", handler.HandleLogin).Methods("POST")
	router.HandleFunc("/register", handler.HandleRegister).Methods("POST")
}

func (handler *Handler) HandleLogin(_ http.ResponseWriter, _ *http.Request) {
	// handle login
}

func (handler *Handler) HandleRegister(_ http.ResponseWriter, _ *http.Request) {
	// handle register
}
