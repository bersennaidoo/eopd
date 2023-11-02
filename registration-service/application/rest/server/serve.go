package server

import (
	"fmt"
	"net/http"
	"time"

	"github.com/bersennaidoo/eopd/registration-service/application/rest/router"
)

type Server struct {
	route *router.Route

	addr string
}

func New(addr string, route *router.Route) *Server {

	return &Server{
		route: route,
		addr:  addr,
	}
}

func (s *Server) ListenAndServe() {

	srv := &http.Server{
		Addr:           s.addr,
		Handler:        s.route.Router(),
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}

	fmt.Println("Listening on port ", s.addr)
	srv.ListenAndServe()
}
