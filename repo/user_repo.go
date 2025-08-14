package repo

import (
	"context"
	"database/sql"

	db "github.com/devsirose/simplebank/db/sqlc"
)

type UserRepository interface {
	GetUserByUsername(ctx context.Context, username string)
}

type userRepository struct {
	conn    *sql.DB
	queries *db.Queries
}

func NewUserRepository(conn *sql.DB) UserRepository {
	return &userRepository{
		conn:    conn,
		queries: db.New(conn),
	}
}

func (r *userRepository) GetUserByUsername(ctx context.Context, username string) {

}
