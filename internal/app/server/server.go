package server

import (
	"net/http"
	"os"

	"github.com/LWich/chat-rest-api/internal/app/config"
	_ "github.com/lib/pq" // ...
	"github.com/sirupsen/logrus"
)

// Server ...
type Server struct {
	Handler http.Handler
	Logger  *logrus.Logger
}

// New ...
func New(handler http.Handler) *Server {
	return &Server{
		Handler: handler,
		Logger:  logrus.New(),
	}
}

// Run ...
func (s *Server) Run(cfg *config.ServerConfig) {
	server := http.Server{
		Addr:    cfg.BindAddr,
		Handler: s.Handler,
	}

	if err := server.ListenAndServe(); err != nil {
		s.Logger.Fatal(err)
		os.Exit(-1)
	}
}
