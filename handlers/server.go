package handlers

import "net/http"

type Server struct {
    router *http.ServeMux
}

func New() Server {
    s := Server{
        router: http.NewServeMux(),
    }

    return s
}
