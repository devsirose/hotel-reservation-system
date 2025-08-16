package util

import (
	db "github.com/devsirose/simplebank/db/sqlc"
	"github.com/devsirose/simplebank/model"
	"github.com/devsirose/simplebank/pb"
)

func AccountToDTO(acc db.Account) model.AccountDTO {
	return model.AccountDTO{
		Owner:    acc.Owner,
		Currency: acc.Currency,
		Balance:  acc.Balance,
	}
}

func CreateAccountGrpcRequestToParams(dto *pb.CreateAccountRequest) db.CreateAccountParams {
	return db.CreateAccountParams{
		Owner:    dto.Owner,
		Currency: dto.Currency,
		Balance:  dto.Balance,
	}
}

func DTOToCreateAccountGrpcResponse(dto model.AccountDTO) *pb.CreateAccountResponse {
	return &pb.CreateAccountResponse{
		Account: &pb.Account{
			Owner:    dto.Owner,
			Currency: dto.Currency,
			Balance:  dto.Balance,
		},
	}
}
