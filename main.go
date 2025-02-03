package main

import (
	"database/sql"
	"log"

	"github.com/joekingsleyMukundi/bank/api"
	db "github.com/joekingsleyMukundi/bank/db/sqlc"
	_ "github.com/lib/pq"
)

const (
	dbdriver      = "postgres"
	dbsource      = "postgresql://root:secret@localhost:5432/simple_bank?sslmode=disable"
	serverAddress = "0.0.0.0:8080"
)

func main() {
	conn, err := sql.Open(dbdriver, dbsource)
	if err != nil {
		log.Fatal("connot connect to db error: ", err)
	}
	store := db.NewStore(conn)
	server := api.NewServer(store)
	err = server.Start(serverAddress)
	if err != nil {
		log.Fatal("cannot connet to server:", err)
	}
}
