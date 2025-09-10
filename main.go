package main

import (
	"context"
	"database/sql"
	"fmt"
	"net"
	"net/http"
	"os"

	"github.com/devsirose/simplebank/api"
	"github.com/devsirose/simplebank/config"
	db "github.com/devsirose/simplebank/db/sqlc"
	"github.com/devsirose/simplebank/gapi"
	"github.com/devsirose/simplebank/logger"
	"github.com/devsirose/simplebank/pb"
	"github.com/devsirose/simplebank/svc"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	_ "github.com/lib/pq" // import this package to run init() function in package
	"go.uber.org/zap"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"google.golang.org/protobuf/encoding/protojson"
)

func main() {
	logger.Init("dev")
	cfg, err := config.LoadConfig(".")
	if err != nil {
		logger.Log.Error(err.Error())
		os.Exit(1)
	}
	runGinServer(cfg)
}
func runGrpcServer(cfg config.Config) {
	grpcServer := grpc.NewServer()
	//run application
	conn, err := sql.Open(cfg.DbDriver, cfg.DbSource)
	if err != nil {
		logger.Log.Error("Failed to connect to database", zap.Error(err))
		os.Exit(1)
	}
	accSvc := svc.NewAccountService(db.NewStore(conn))
	accSvcServer := gapi.NewAccountGRPCServer(accSvc)
	pb.RegisterAccountServiceServer(grpcServer, accSvcServer)
	reflection.Register(grpcServer) // client can explore rpc(s) on server & how to call them

	listener, err := net.Listen("tcp", cfg.ServerHost+":"+cfg.GRPCServerPort)
	if err != nil {
		logger.Log.Error(fmt.Sprintf("Failed to listen on port %s", cfg.GRPCServerPort), zap.Error(err))
	}
	logger.Log.Info(fmt.Sprintf("GRPC server used port %s", cfg.GRPCServerPort))

	err = grpcServer.Serve(listener)
	if err != nil {
		logger.Log.Error(fmt.Sprintf("Failed to start gRPC server: %s", err.Error()))
	}
}
func runGatewayServer(cfg config.Config) {
	//run application
	conn, err := sql.Open(cfg.DbDriver, cfg.DbSource)
	if err != nil {
		logger.Log.Error("Failed to connect to database", zap.Error(err))
		os.Exit(1)
	}

	accSvc := svc.NewAccountService(db.NewStore(conn))
	accSvcServer := gapi.NewAccountGRPCServer(accSvc)

	grpcMux := runtime.NewServeMux(
		runtime.WithMarshalerOption(runtime.MIMEWildcard, &runtime.JSONPb{
			MarshalOptions: protojson.MarshalOptions{
				UseProtoNames: true,
			},
			UnmarshalOptions: protojson.UnmarshalOptions{
				DiscardUnknown: true,
			},
		}),
	)
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	err = pb.RegisterAccountServiceHandlerServer(ctx, grpcMux, accSvcServer)
	if err != nil {
		logger.Log.Error("Failed to register gRPC gateway handler", zap.Error(err))
	}

	httpMux := http.NewServeMux()
	httpMux.Handle("/", grpcMux)

	listener, err := net.Listen("tcp", cfg.ServerHost+":"+cfg.HTTPServerPort)
	if err != nil {
		logger.Log.Error(fmt.Sprintf("Failed to listen on port %s", cfg.HTTPServerPort), zap.Error(err))
	}
	logger.Log.Info(fmt.Sprintf("Started HTTP gateway at %s", cfg.HTTPServerPort))

	err = http.Serve(listener, httpMux)
	if err != nil {
		logger.Log.Error(fmt.Sprintf("Failed to start gRPC server: %s", err.Error()))
	}
}

func runGinServer(cfg config.Config) {
	//run application
	conn, err := sql.Open(cfg.DbDriver, cfg.DbSource)
	if err != nil {
		logger.Log.Error("Failed to connect to database", zap.Error(err))
		os.Exit(1)
	}
	server := api.NewServer(db.NewStore(conn))
	if err := server.Start(cfg.ServerHost + ":" + cfg.HTTPServerPort); err != nil {
		logger.Log.Error("Failed to start server", zap.Error(err))
		os.Exit(1)
	}
}
