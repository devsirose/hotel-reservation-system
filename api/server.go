package api

import (
	db "github.com/devsirose/simplebank/db/sqlc"
	"github.com/devsirose/simplebank/middleware"
	"github.com/gin-gonic/gin"
)

type Server struct {
	store  db.Store
	router *gin.Engine
}

func NewServer(store db.Store) *Server {
	server := &Server{
		store: store,
	}

	router := gin.Default()
	router.Use(middleware.RecoveryWithLogger)

	router.POST("/api/v1/accounts", server.CreateAccount)
	router.GET("/api/v1/accounts/:id", server.GetAccountBy)

	server.router = router
	return server
}

func (server *Server) Start(address string) error {
	return server.router.Run(address)
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
