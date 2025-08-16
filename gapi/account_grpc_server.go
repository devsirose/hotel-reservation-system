package gapi

import (
	"github.com/devsirose/simplebank/pb"
	"github.com/devsirose/simplebank/svc"
)

type AccountServiceGrpcServer struct {
	svc svc.AccountService
	pb.UnimplementedAccountServiceServer
}
