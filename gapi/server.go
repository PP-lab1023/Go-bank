package gapi

import (
	"fmt"

	db "github.com/PP-lab1023/Go-bank/db/sqlc"
	"github.com/PP-lab1023/Go-bank/pb"
	"github.com/PP-lab1023/Go-bank/token"
	"github.com/PP-lab1023/Go-bank/util"
	"github.com/PP-lab1023/Go-bank/worker"
)

// Server serves gRPC requests for our banking service.
type Server struct {
	pb.UnimplementedGoBankServer
	config          util.Config
	store           db.Store
	tokenMaker      token.Maker
	taskDistributor worker.TaskDistributor
}

// NewServer creates a new gRPC server.
func NewServer(config util.Config, store db.Store, taskDistributor worker.TaskDistributor) (*Server, error) {
	tokenMaker, err := token.NewPasetoMaker(config.TokenSymmetricKey)
	if err != nil {
		return nil, fmt.Errorf("cannot create token maker: %w", err)
	}

	server := &Server{
		config:          config,
		store:           store,
		tokenMaker:      tokenMaker,
		taskDistributor: taskDistributor,
	}

	return server, nil
}
