package main

import (
	"database/sql"
	"log"

	"github.com/matoanbach/simple-bank/api"
	db "github.com/matoanbach/simple-bank/db/sqlc"
)

const (
	DBDriver = "postgres"
	DBSource = "postgresql://root:secret@localhost:49466/simple_bank?sslmode=disable"
)

func main() {
	conn, err := sql.Open(DBDriver, DBSource)
	if err != nil {
		log.Fatal("unable to connec to DB", err)
	}
	store := db.NewStore(conn)
	// runServer(store)Serve
	server, _ := api.NewServer(store)
	server.Serve()
}

func runServer(store *db.Store) {
	server, err := api.NewServer(store)
	if err != nil {
		log.Fatal("cannot create the server")
	}

	server.Serve()
}
