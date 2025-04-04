package gapi

import (
	"fmt"

	db "github.com/PP-lab1023/Go-bank/db/sqlc"
	"github.com/PP-lab1023/Go-bank/pb"
	"github.com/PP-lab1023/Go-bank/token"
	"github.com/PP-lab1023/Go-bank/util"
)

// Server serves gRPC requests for our banking service.
type Server struct {
	pb.UnimplementedGoBankServer
	config     util.Config
	store      db.Store
	tokenMaker token.Maker
}

// NewServer creates a new gRPC server.
func NewServer(config util.Config, store db.Store) (*Server, error) {
	tokenMaker, err := token.NewPasetoMaker(config.TokenSymmetricKey)
	if err != nil {
		return nil, fmt.Errorf("cannot create token maker: %w", err)
	}

	server := &Server{
		config:     config,
		store:      store,
		tokenMaker: tokenMaker,
	}

	return server, nil
}
