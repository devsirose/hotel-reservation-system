package gapi

import (
	"github.com/devsirose/simplebank/config"
	db "github.com/devsirose/simplebank/db/sqlc"
	"github.com/devsirose/simplebank/pb"
)

// Server serves gRPC requests for our banking service.
type Server struct {
	pb.UnimplementedAccountServiceServer
	config config.Config
	store  db.Store
}

// NewServer creates a new gRPC server.
func NewServer(config config.Config, store db.Store) *Server {
	return &Server{
		config: config,
		store:  store,
	}
}
