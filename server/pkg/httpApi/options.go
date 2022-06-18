package httpApi

import "net"

type Option func(*Server)

// Port
func Port(port string) Option {
	return func(s *Server) {
		s.httpServer.Addr = net.JoinHostPort("", port)
	}
}