package server

import (
	"log"
	"net/http"
)

type Server struct {
	router Router
}

func NewServer(router Router) Server {
	return Server{
		router: router,
	}
}

func (s *Server) InitServer() error {
	server := http.Server{
		Addr:    ":8080",
		Handler: s.router.InitRouter(),
	}

	log.Println("Starting HTTP server at port 8080")
	return server.ListenAndServe()
}
