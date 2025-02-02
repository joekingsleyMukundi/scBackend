package db

import (
	"database/sql"
	"log"
	"os"
	"testing"

	_ "github.com/lib/pq"
)

const (
	dbdriver = "postgres"
	dbsource = "postgresql://root:secret@localhost:5432/simple_bank?sslmode=disable"
)

var testQueries *Queries
var testDb *sql.DB

func TestMain(m *testing.M) {
	var err error
	testDb, err = sql.Open(dbdriver, dbsource)
	if err != nil {
		log.Fatal("connot connect to db error: ", err)
	}
	testQueries = New(testDb)
	os.Exit(m.Run())
}
