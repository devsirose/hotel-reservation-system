package test

import (
	"database/sql"
	"log"
	"os"
	"testing"

	db "github.com/devsirose/simplebank/db/sqlc"
	_ "github.com/lib/pq"
)

var testQueries *db.Queries

const (
	dbDriver = "postgres"
	dbSource = "postgres://root:mysecretpassword@localhost:5430/simple_bank?sslmode=disable"
)

func TestMain(m *testing.M) {
	conn, err := sql.Open(dbDriver, dbSource)
	if err != nil {
		log.Fatal("failed to connect to database", err)
	}

	testQueries = db.New(conn)

	os.Exit(m.Run())
}
