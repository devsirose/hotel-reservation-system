package main

import (
	"database/sql"
	"os"

	"github.com/devsirose/simplebank/api"
	"github.com/devsirose/simplebank/config"
	db "github.com/devsirose/simplebank/db/sqlc"
	"github.com/devsirose/simplebank/logger"
	_ "github.com/lib/pq" // import this package to run init() function in package
	"go.uber.org/zap"
)

func main() {
	logger.Init("dev")
	cfg, err := config.LoadConfig(".")
	if err != nil {
		logger.Log.Error(err.Error())
		os.Exit(1)
	}

	//run application
	conn, err := sql.Open(cfg.DbDriver, cfg.DbSource)
	if err != nil {
		logger.Log.Error("Failed to connect to database", zap.Error(err))
		os.Exit(1)
	}
	server := api.NewServer(db.NewStore(conn))
	if err := server.Start(cfg.ServerHost + ":" + cfg.ServerPort); err != nil {
		logger.Log.Error("Failed to start server", zap.Error(err))
		os.Exit(1)
	}

}
