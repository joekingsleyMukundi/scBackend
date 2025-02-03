package db

import (
	"database/sql"
	"log"
	"os"
	"testing"

	"github.com/joekingsleyMukundi/bank/util"
	_ "github.com/lib/pq"
)

const (
	dbdriver = "postgres"
	dbsource = "postgresql://root:secret@localhost:5432/simple_bank?sslmode=disable"
)

var testQueries *Queries
var testDb *sql.DB

func TestMain(m *testing.M) {
	config, err := util.LoadConfig("../..")
	if err != nil {
		log.Fatal("connot get config", err)
	}
	testDb, err = sql.Open(config.DBDriver, config.DBSource)
	if err != nil {
		log.Fatal("connot connect to db error: ", err)
	}
	testQueries = New(testDb)
	os.Exit(m.Run())
}
