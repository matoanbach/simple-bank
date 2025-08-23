package gapi

import (
	"fmt"

	db "github.com/matoanbach/simple-bank/db/sqlc"
	"github.com/matoanbach/simple-bank/db/util"
	"github.com/matoanbach/simple-bank/pb"
	"github.com/matoanbach/simple-bank/token"
)

// Server serves HTTP requests for our banking service
type Server struct {
	pb.UnimplementedSimpleBankServer
	config     util.Config
	store      db.Store
	tokenMaker token.Maker
}

// NewServer creates a new GRPC server
func NewServer(config util.Config, store db.Store) (*Server, error) {
	tokenMaker, err := token.NewPasetoMaker(config.TokenSymmetric)
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
