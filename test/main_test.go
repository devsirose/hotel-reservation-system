package test

import (
	"os"
	"testing"

	mockdb "github.com/devsirose/simplebank/db/mock"
	_ "github.com/lib/pq"
)

var mockStore *mockdb.MockStore

func TestMain(m *testing.M) {

	os.Exit(m.Run())
}
