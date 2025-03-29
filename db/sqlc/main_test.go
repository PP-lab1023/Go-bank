package db

import (
	"database/sql"
	"log"
	"os"
	"testing"

	"github.com/PP-lab1023/Go-bank/util"
	_ "github.com/lib/pq" // This is a special import, because no function in this package is used. So a _ should be added.
)

var testQueries *Queries
var testDB *sql.DB

func TestMain(m *testing.M) {
	var err error 

	config, err := util.LoadConfig("../..")
	if err != nil {
		log.Fatal("cannot connect to db", err)
	}
	
	testDB, err = sql.Open(config.DBDriver, config.DBSource)
	if err != nil {
		log.Fatal("cannot connect to db", err)
	}

	testQueries = New(testDB)

	os.Exit(m.Run()) 
}