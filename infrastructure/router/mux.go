package router

import (
	"net/http"

	"github.com/gorilla/mux"
)

type Router struct {
	router *mux.Router
}

func NewMuxRouter() *Router {
	return &Router{mux.NewRouter()}
}

func (r *Router) GET(uri string, f func(w http.ResponseWriter, r *http.Request)) {
	r.router.HandleFunc(uri, f).Methods("GET")
}

func (r *Router) POST(uri string, f func(w http.ResponseWriter, r *http.Request)) {
	r.router.HandleFunc(uri, f).Methods("POST")
}

func (r *Router) PUT(uri string, f func(w http.ResponseWriter, r *http.Request)) {
	r.router.HandleFunc(uri, f).Methods("PUT")
}

func (r *Router) DELETE(uri string, f func(w http.ResponseWriter, r *http.Request)) {
	r.router.HandleFunc(uri, f).Methods("DELETE")
}

func (r *Router) SERVE(port string) {
	http.ListenAndServe(port, r.router)
}