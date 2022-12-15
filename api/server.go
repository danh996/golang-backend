package api

import (
	"fmt"

	db "github.com/danh996/go-school/db/sqlc"
	"github.com/danh996/go-school/token"
	"github.com/danh996/go-school/util"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/validator/v10"
)

type Server struct {
	store      db.Store
	tokenMaker token.Maker
	router     *gin.Engine
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

	if v, ok := binding.Validator.Engine().(*validator.Validate); ok {
		v.RegisterValidation("currentcy", validCurrency)
	}

	server.setupRouter()

	return server, nil
}

func (server *Server) Start(address string) error {
	err := server.router.Run(address)
	if err != nil {
		return err
	}
	return nil
}

func (s *Server) setupRouter() {
	router := gin.Default()

	router.POST("/users", s.createUser)
	router.POST("/users/login", s.loginUser)

	authRoutes := router.Group("/").Use(authMiddleware(s.tokenMaker))

	authRoutes.POST("/accounts", s.createAccount)
	authRoutes.GET("/accounts/:id", s.getAccount)
	authRoutes.GET("/accounts", s.listAccount)

	authRoutes.POST("/transfers", s.createTransfer)

	s.router = router

}

func errorResponse(err error) gin.H {
	return gin.H{
		"error": err.Error(),
	}
}
