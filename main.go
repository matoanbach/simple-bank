package main

import (
	"database/sql"
	"log"

	"github.com/matoanbach/simple-bank/api"
	db "github.com/matoanbach/simple-bank/db/sqlc"
	"github.com/matoanbach/simple-bank/db/util"
)

const (
	DBDriver = "postgres"
	DBSource = "postgresql://root:secret@localhost:49466/simple_bank?sslmode=disable"
)

func main() {
	config, err := util.LoadConfig(".")
	if err != nil {
		log.Fatal("cannot load config")
	}

	conn, err := sql.Open(DBDriver, DBSource)
	if err != nil {
		log.Fatal("unable to connec to DB", err)
	}
	store := db.NewStore(conn)
	// runServer(store)Serve
	server, _ := api.NewServer(config, store)
	server.Serve()
}
