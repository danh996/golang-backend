package api

import (
	db "github.com/danh996/go-school/db/sqlc"
	"github.com/gin-gonic/gin"
)

type Server struct {
	store  *db.Store
	router *gin.Engine
}

func NewServer(store *db.Store) *Server {
	server := &Server{store: store}
	router := gin.Default()

	//add routes to router
	router.POST("/accounts", server.createAccount)

	server.router = router
	return server
}

func (server *Server) Start(address string) error {
	err := server.router.Run(address)
	if err != nil {
		return err
	}
	return nil
}

func errorResponse(err error) gin.H {
	return gin.H{
		"error": err.Error(),
	}
}
