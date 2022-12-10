package main

import (
	"database/sql"
	"fmt"
	"log"

	"github.com/danh996/go-school/api"
	db "github.com/danh996/go-school/db/sqlc"
	_ "github.com/lib/pq"
)

const (
	dbDriver      = "postgres"
	dbSource      = "postgresql://postgres:secret@localhost:5432/simple_bank?sslmode=disable"
	serverAddress = "0.0.0.0:8089"
)

func main() {
	conn, err := sql.Open(dbDriver, dbSource)

	if err != nil {
		log.Fatal("can not connect to DB", err)
	}

	store := db.NewStore(conn)
	srv := api.NewServer(store)
	err = srv.Start(serverAddress)

	if err != nil {
		log.Fatal("can not start server", err)
	}

	fmt.Println("start server success")
}
