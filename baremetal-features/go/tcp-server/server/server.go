// The server file responsble for  Listening on a socket, accepting connections, and parsing messages. 
// It doesn't care where the port number came from; it just uses what it's given.

package server

import (
	"bufio"
	"fmt"
	"net"
	"log/slog"
)

type Server struct {
	Addr string
	Logger *slog.Logger
} 

func New(addr string, logger *slog.Logger) *Server{
	return &Server{
		Addr: addr,
		Logger: logger,
	}
}

func (s *Server) Start() error {
	listener, err := net.Listen("tcp", s.Addr)
	if err != nil {
		return fmt.Errorf("failed to listen on %s: %w", s.Addr, err)
	}
	defer listener.Close()

	s.Logger.Info("Server started", "address", s.Addr)

	for {
		conn, err := listener.Accept()
		if err != nil {
			s.Logger.Error("Accept error", "error", err)
			continue
		}

		go s.handleConnection(conn)
	}
}

func(s *Server) handleConnection(conn net.Conn) {
	defer conn.Close()

	reader := bufio.NewReader(conn)

	for {
		message, err := reader.ReadString('\n')
		if err != nil {
			s.Logger.Debug("Connection closed", "remote_addr", conn.RemoteAddr())
			return
		}

		s.Logger.Info("Message received", "content", message, "from", conn.RemoteAddr())

		response := "Echo: " + message
		conn.Write([]byte(response))
	}
}