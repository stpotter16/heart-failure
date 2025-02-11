package handlers

func addRoutes(s Server) {
	mux := s.router

	mux.HandleFunc("GET /{$}", s.indexGet())
}
