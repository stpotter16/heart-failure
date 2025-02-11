package handlers

import "net/http"

type Server struct {
	router *http.ServeMux
}

func (s Server) Handler() http.Handler {
	return s.router
}

func New() Server {
	s := Server{
		router: http.NewServeMux(),
	}

	addRoutes(s)

	return s
}
