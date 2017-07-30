package handler

import (
	"github.com/gorilla/mux"
	"net/http"
)

type Handler struct {
	router *mux.Router
}

func New() *Handler {
	return &Handler{}
}

func (h Handler) Router() *mux.Router {
	return h.router
}

func (h *Handler) Regist() {
	h.router = mux.NewRouter()
	h.router.HandleFunc("/logging/test", h.LoggingTest)
}

func (h *Handler) LoggingTest(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello World"))
}
