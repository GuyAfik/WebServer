package httpApi

import (
	"fmt"
	"net/http"

	"github.com/WebServer/server/pkg/logger"
)

// Server -.
type Server struct {
	httpServer *http.Server
	logger logger.LoggerService
}

func New(handler http.Handler, logger logger.LoggerService, options ...Option) {
	
	srvr := &Server{
		httpServer: &http.Server{
			Handler: handler,
		},
		logger: logger,
	}

	for _, opt := range options {
		opt(srvr)
	}
	srvr.start()
}

func (s *Server) start() {
	s.logger.Info("Starting server at address %s", fmt.Sprintf("localhost%s", s.httpServer.Addr))
	defer s.httpServer.Close()
	s.httpServer.ListenAndServe()
}