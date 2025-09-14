package main

import (
	"database/sql"
	"log"
	"os"

	"github.com/devsirose/hotel-reservation/api"
	"github.com/devsirose/hotel-reservation/config"
	db "github.com/devsirose/hotel-reservation/db/sqlc"
	"github.com/devsirose/hotel-reservation/logger"
	_ "github.com/lib/pq"
	"go.uber.org/zap"
)

func main() {
	logger.Init("dev")
	cfg, err := config.LoadConfig(".")
	if err != nil {
		logger.Log.Error(err.Error())
		os.Exit(1)
	}

	dbSQL, err := sql.Open(cfg.DbDriver, cfg.DbSource)
	if err != nil {
		logger.Log.Error("Failed to connect to database", zap.Error(err))
		os.Exit(1)
	}
	defer dbSQL.Close()

	if err := dbSQL.Ping(); err != nil {
		logger.Log.Error("Failed to ping database", zap.Error(err))
		os.Exit(1)
	}

	logger.Log.Info("Database connected successfully",
		zap.String("driver", cfg.DbDriver),
	)

	// Create db store
	store := db.NewStore(dbSQL)

	// Create API server
	server := api.NewServer(store, dbSQL)

	serverAddr := cfg.ServerHost + ":" + cfg.HTTPServerPort
	logger.Log.Info("Starting Hotel Reservation System",
		zap.String("address", serverAddr),
		zap.String("app_name", cfg.AppName),
		zap.String("environment", "development"),
	)
	
	if err := server.Start(serverAddr); err != nil {
		log.Fatal("Failed to start server:", err)
	}
}