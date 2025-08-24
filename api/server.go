package api

import (
	"github.com/gin-gonic/gin"
	db "github.com/matoanbach/simple_bank/db/sqlc"
)

type Server struct {
	store  *db.Store
	router *gin.Engine
}

func (server *Server) Serve() {
	server.router.Run()
}

func NewServer(store *db.Store) (server *Server, err error) {
	// auth

	//
	server = &Server{store: store}

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

	// //routers
	// router.POST("/users", server.createUser)
	// router.POST("/users/login", server.loginUser)
	// router.POST("/tokens/renew_access", server.renewAccessToken)

	// authRoutes := router.Group("/").User(authMiddleWare(server.tokenMaker))
	router.POST("/accounts", server.createAccount)
	router.GET("/accounts/:id", server.getAccount)
	router.GET("/accounts", server.listAccount)
	// router.POST("/transfer", server.createTransfer)

	server.router = router
}
