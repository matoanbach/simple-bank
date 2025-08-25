package api

import (
	"fmt"

	"github.com/gin-gonic/gin"
	db "github.com/matoanbach/simple-bank/db/sqlc"
	"github.com/matoanbach/simple-bank/db/util"
	"github.com/matoanbach/simple-bank/token"
)

type Server struct {
	config     util.Config
	store      db.Store
	tokenMaker token.Maker
	router     *gin.Engine
}

func (server *Server) Serve() {
	server.router.Run()
}

func NewServer(config util.Config, store db.Store) (*Server, error) {
	// auth
	tokenMaker, err := token.NewJWTMaker(config.TokenSymmetric)
	if err != nil {
		return nil, fmt.Errorf("cannot create token maker: %w", err)
	}

	//
	server := &Server{
		config:     config,
		store:      store,
		tokenMaker: tokenMaker,
	}

	server.setupServer()
	return server, nil
}

func errorResponse(err error) gin.H {
	return gin.H{
		"error": err.Error(),
	}
}

func (server *Server) setupServer() {
	router := gin.Default()

	//routers
	router.POST("/users", server.createUser)
	router.POST("/users/login", server.loginUser)
	// router.POST("/tokens/renew_access", server.renewAccessToken)

	// authRoutes := router.Group("/").User(authMiddleWare(server.tokenMaker))
	router.POST("/accounts", server.createAccount)
	router.GET("/accounts/:id", server.getAccount)
	router.GET("/accounts", server.listAccount)
	router.POST("/transfers", server.createTransfer)

	server.router = router
}
