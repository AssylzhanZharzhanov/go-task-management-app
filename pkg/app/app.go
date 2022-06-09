package app

import (
	"context"
	"net/http"
	"time"
)

type Server struct {
	Handlers   http.Handler
	Port       string
	httpServer *http.Server
}

func NewServer(port string) *Server {
	return &Server{
		//Handlers: handlers,
		Port: port,
	}
}

func (s *Server) Run() error {
	s.httpServer = &http.Server{
		Addr: ":" + s.Port,
		//Handler: s.Handlers,
		MaxHeaderBytes: 1 << 20,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
	}

	return s.httpServer.ListenAndServe()
}

func (s *Server) Shutdown(ctx context.Context) error {
	return s.httpServer.Shutdown(ctx)
}
