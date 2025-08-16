package svc

import (
	"context"

	db "github.com/devsirose/simplebank/db/sqlc"
	"github.com/devsirose/simplebank/model"
	"github.com/devsirose/simplebank/util"
)

type AccountService interface {
	CreateAccount(context.Context, db.CreateAccountParams) (model.AccountDTO, error)
	TransferTx(ctx context.Context, arg db.TransferTxParams) (db.TransferTxResult, error)
}

type AccountServiceImpl struct {
	store db.Store
}

func NewAccountService(store db.Store) *AccountServiceImpl {
	return &AccountServiceImpl{store: store}
}

func (s AccountServiceImpl) CreateAccount(ctx context.Context, arg db.CreateAccountParams) (model.AccountDTO, error) {
	acc, err := s.store.CreateAccount(ctx, db.CreateAccountParams{
		Owner:    arg.Owner,
		Balance:  arg.Balance,
		Currency: arg.Currency,
	})
	if err != nil {
		return model.AccountDTO{}, err
	}
	// Convert db.Account to AccountDTO
	return util.AccountToDTO(acc), err
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
