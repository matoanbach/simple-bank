package db

import (
	"database/sql"
	"log"
	"os"
	"testing"

	_ "github.com/lib/pq"
)

var testQueries *Queries
var testDB *sql.DB

const (
	DBDriver = "postgres"
	DBSource = "postgresql://root:secret@localhost:49466/simple_bank?sslmode=disable"
)

func TestMain(m *testing.M) {
	var err error
	testDB, err = sql.Open(DBDriver, DBSource)
	if err != nil {
		log.Fatal("unable to connec to DB", err)
	}
	testQueries = New(testDB)

	os.Exit(m.Run())
}
