package gapi

import (
	"context"

	"github.com/devsirose/simplebank/pb"
	"github.com/devsirose/simplebank/svc"
	"github.com/devsirose/simplebank/util"
)

type AccountGrpcServer interface {
	CreateAccount(context.Context, *pb.CreateAccountRequest) (*pb.CreateAccountResponse, error)
}

type AccountGrpcServerImpl struct {
	svc svc.AccountService
	pb.UnimplementedAccountServiceServer
}

func NewAccountGRPCServer(s svc.AccountService) *AccountGrpcServerImpl {
	return &AccountGrpcServerImpl{
		svc: s,
	}
}

func (s *AccountGrpcServerImpl) CreateAccount(ctx context.Context, req *pb.CreateAccountRequest) (*pb.CreateAccountResponse, error) {
	accDto, err := s.svc.CreateAccount(ctx, util.CreateAccountGrpcRequestToParams(req))
	if err != nil {
		return nil, err
	}
	return util.DTOToCreateAccountGrpcResponse(accDto), nil
}
