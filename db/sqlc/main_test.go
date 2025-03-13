package db

import (
	"database/sql"
	"log"
	"os"
	"testing"

	_ "github.com/lib/pq" // This is a special import, because no function in this package is used. So a _ should be added.
)

const (
	dbDriver = "postgres"
	dbSource = "postgresql://root:secret@localhost:5432/Go-bank?sslmode=disable"
)

var testQueries *Queries

func TestMain(m *testing.M) {
	conn, err := sql.Open(dbDriver, dbSource)
	if err != nil {
		log.Fatal("cannot connect to db", err)
	}

	testQueries = New(conn)

	os.Exit(m.Run()) 
}