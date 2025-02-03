package main

import (
	"database/sql"
	"log"

	"github.com/joekingsleyMukundi/bank/api"
	db "github.com/joekingsleyMukundi/bank/db/sqlc"
	"github.com/joekingsleyMukundi/bank/util"
	_ "github.com/lib/pq"
)

func main() {
	config, err := util.LoadConfig(".")
	if err != nil {
		log.Fatal("connot connect to db error: ", err)
	}
	conn, err := sql.Open(config.DBDriver, config.DBSource)
	if err != nil {
		log.Fatal("connot connect to db error: ", err)
	}
	store := db.NewStore(conn)
	server := api.NewServer(store)
	err = server.Start(config.ServerAddress)
	if err != nil {
		log.Fatal("cannot connet to server:", err)
	}
}
