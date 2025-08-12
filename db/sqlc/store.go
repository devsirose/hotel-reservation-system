package db

import (
	"context"
	"database/sql"
	"fmt"
)

type Store struct {
	db      *sql.DB  // db connection manager
	Queries *Queries // query interface
} // EntityManager

func NewStore(db *sql.DB) *Store {
	return &Store{
		db:      db,
		Queries: New(db),
	}
}

func (store *Store) ExecTx(ctx context.Context, fn func(*Queries) error) error {
	tx, err := store.db.BeginTx(ctx, nil) // lay mot connection tu pool luu vao tx -> mo 1 commit
	if err != nil {
		return err
	}
	q := New(tx) // build cac cau query tu tx
	err = fn(q)  // do lambda function (truyen q de lambda function consume)
	if err != nil {
		if rbErr := tx.Rollback(); rbErr != nil {
			return fmt.Errorf("tx rollback failed: %w", rbErr)
		}
		return err
	}
	return tx.Commit()
}
