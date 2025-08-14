package api

import (
	db "github.com/devsirose/simplebank/db/sqlc"
	"github.com/devsirose/simplebank/middleware"
	"github.com/gin-gonic/gin"
)

type Server struct {
	store  db.Store
	Router *gin.Engine
}

func NewServer(store db.Store) *Server {
	server := &Server{
		store: store,
	}

	router := gin.Default()
	router.Use(middleware.RecoveryWithLogger)

	router.POST("/api/v1/accounts", server.CreateAccount)
	router.GET("/api/v1/accounts/:id", server.GetAccountById)

	server.Router = router
	return server
}

func (server *Server) Start(address string) error {
	return server.Router.Run(address)
}

func (server *Server) Close() error {
	return server.Close()
}

func errorResponse(err error) gin.H {
	return gin.H{
		//custom err response here
		"error": err.Error(),
	}
}

func notFoundResponse(obj any) gin.H {
	return gin.H{
		"error": obj.(string) + " not found",
	}
}
