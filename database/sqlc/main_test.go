package database

import (
	"database/sql"
	_ "github.com/lib/pq"
	"os"
	"testing"
)

const (
	dbDriver = "postgres"
	dbSource = "postgresql://postgres@localhost:5432/practice?sslmode=disable"
)

var testQueries *Queries

func TestMain(m *testing.M) {
	conn, err := sql.Open(dbDriver, dbSource)
	if err != nil {
		panic(err)
	}
	testQueries = New(conn)

	os.Exit(m.Run())
}
