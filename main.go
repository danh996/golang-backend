package main

import (
	"database/sql"
	"fmt"
	"log"

	"github.com/danh996/go-school/api"
	db "github.com/danh996/go-school/db/sqlc"
	"github.com/danh996/go-school/util"
	_ "github.com/lib/pq"
)

func main() {
	config, err := util.LoadConfig(".")
	if err != nil {
		log.Fatal("can not load config", err)
	}
	conn, err := sql.Open(config.DBDriver, config.DBSource)

	if err != nil {
		log.Fatal("can not connect to DB", err)
	}

	store := db.NewStore(conn)
	srv := api.NewServer(store)
	err = srv.Start(config.ServerAddress)

	if err != nil {
		log.Fatal("can not start server", err)
	}

	fmt.Println("start server success")
}
