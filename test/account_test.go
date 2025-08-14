package test

import (
	"context"
	"testing"

	db "github.com/devsirose/simplebank/db/sqlc"
	"github.com/stretchr/testify/require"
)

type User struct {
	ID       int64  `json:"id"`
	Username string `json:"username"`
	Password string `json:"password"`
}

func TestCreateAccount(t *testing.T) {
	accountParams := db.CreateAccountParams{
		Owner:    "johndoe",
		Balance:  1000,
		Currency: "USD",
	}
	acc, err := testQueries.CreateAccount(context.Background(), accountParams)
	require.NoError(t, err)
	require.NotEmpty(t, acc)

	require.NotZero(t, acc.ID)
}
