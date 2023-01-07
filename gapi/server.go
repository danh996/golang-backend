package gapi

import (
	"fmt"

	db "github.com/danh996/golang-backend/db/sqlc"
	"github.com/danh996/golang-backend/pb"
	"github.com/danh996/golang-backend/token"
	"github.com/danh996/golang-backend/util"
)

type Server struct {
	pb.UnimplementedSimpleBankServer
	store      db.Store
	tokenMaker token.Maker
	config     util.Config
}

func NewServer(config util.Config, store db.Store) (*Server, error) {
	tokenMaker, err := token.NewPasetoMaker(config.TokenSymmetricKey)
	if err != nil {
		return nil, fmt.Errorf("cannot create token maker: %w", err)
	}

	server := &Server{
		store:      store,
		tokenMaker: tokenMaker,
		config:     config,
	}

	return server, nil
}
