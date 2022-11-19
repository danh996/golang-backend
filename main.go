package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"

	_ "github.com/lib/pq" // add this
)

const webPort = ":8008"

func main() {
	fmt.Println("School App is running")

	// connect to database
	connStr := "postgresql://<postgres>:<secret>@<127.0.0.1>/go_booking?sslmode=disable"
	// Connect to database
	_, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Connected to Postgres database")

	http.HandleFunc("/", getRoot)
	http.HandleFunc("/hello", getHello)

	err = http.ListenAndServe(webPort, nil)

	if err != nil {
		fmt.Println(err)
	}

	fmt.Println("HTTP Server is running on Port: ", webPort)
}
