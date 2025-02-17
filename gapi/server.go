package gapi

import (
	"fmt"

	db "github.com/joekingsleyMukundi/bank/db/sqlc"
	"github.com/joekingsleyMukundi/bank/pb"
	"github.com/joekingsleyMukundi/bank/tokens"
	"github.com/joekingsleyMukundi/bank/util"
)

type Server struct {
	pb.UnimplementedSimpleBankServer
	config     util.Config
	store      db.Store
	tokenMaker tokens.Maker
	// taskDistributor worker.TaskDistributor
}

func NewServer(config util.Config, store db.Store) (*Server, error) {
	tokenMaker, err := tokens.NewPasetoMaker(config.TokenSymmetricKey)
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
