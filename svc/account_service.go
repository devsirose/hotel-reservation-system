package svc

import (
	"context"

	db "github.com/devsirose/simplebank/db/sqlc"
	"github.com/devsirose/simplebank/pb"
	"google.golang.org/protobuf/types/known/timestamppb"
)

type AccountService interface {
	CreateAccount(context.Context, *pb.CreateAccountRequest) (*pb.CreateAccountResponse, error)
	TransferTx(ctx context.Context, arg db.TransferTxParams) (db.TransferTxResult, error)
}

type AccountServiceImpl struct {
	store db.Store
	pb.UnimplementedAccountServiceServer
}

func NewAccountService(store db.Store) *AccountServiceImpl {
	return &AccountServiceImpl{store: store}
}

func (s AccountServiceImpl) CreateAccount(ctx context.Context, req *pb.CreateAccountRequest) (*pb.CreateAccountResponse, error) {
	acc, err := s.store.CreateAccount(ctx, db.CreateAccountParams{
		Owner:    req.Owner,
		Balance:  req.Balance,
		Currency: req.Currency,
	})
	if err != nil {
		return nil, err
	}

	pbAcc := &pb.Account{
		Owner:     acc.Owner,
		Balance:   acc.Balance,
		Currency:  acc.Currency,
		CreatedAt: timestamppb.New(acc.CreatedAt),
	}

	return &pb.CreateAccountResponse{
		Account: pbAcc,
	}, nil
}

func (s AccountServiceImpl) TransferTx(ctx context.Context, arg db.TransferTxParams) (db.TransferTxResult, error) {
	var result db.TransferTxResult
	err := s.store.ExecTx(ctx, func(q *db.Queries) error {
		var err error
		if result.Transfer, err = q.CreateTransfer(ctx, db.CreateTransferParams{
			FromAccountID: arg.FromAccountID,
			ToAccountID:   arg.ToAccountID,
			Amount:        arg.Amount,
		}); err != nil {
			return err
		}

		if result.FromEntry, err = q.CreateEntry(ctx, db.CreateEntryParams{
			AccountID: arg.FromAccountID,
			Amount:    -arg.Amount,
		}); err != nil {
			return err
		}

		if result.ToEntry, err = q.CreateEntry(ctx, db.CreateEntryParams{
			AccountID: arg.ToAccountID,
			Amount:    arg.Amount,
		}); err != nil {
			return err
		}
		//update account balance here
		return nil
	})
	if err != nil {
		return result, err
	}

	return result, nil
}
