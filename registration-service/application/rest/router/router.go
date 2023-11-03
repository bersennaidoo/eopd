package router

import (
	"github.com/bersennaidoo/eopd/registration-service/application/rest/handler"
	"github.com/gorilla/mux"
)

type Route struct {
	handler *handler.Handler
}

func New(h *handler.Handler) *Route {

	return &Route{
		handler: h,
	}
}

func (r *Route) Router() *mux.Router {

	rr := mux.NewRouter()
	router := rr.PathPrefix("/eopd/patient/").Subrouter()

	router.HandleFunc("/", r.handler.HandleTest).Methods("GET")
	router.HandleFunc("/register", r.handler.HandleRegister).Methods("POST")
	router.HandleFunc("/update", r.handler.HandleUpdate).Methods("PUT")
	router.HandleFunc("/view/{id}", r.handler.HandleView).Methods("GET")

	return router
}
