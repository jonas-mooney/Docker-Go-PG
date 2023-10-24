package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"

	_ "github.com/lib/pq"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		// Open a connection to the PostgreSQL database
		connStr := "user=jouser dbname=jodb sslmode=disable"
		db, err := sql.Open("postgres", connStr)
		if err != nil {
			log.Fatal(err)
		}
		defer db.Close()

		// Create a test table and insert data
		_, err = db.Exec("CREATE TABLE IF NOT EXISTS test (id serial primary key, data text);")
		if err != nil {
			log.Fatal(err)
		}
		_, err = db.Exec("INSERT INTO test (data) VALUES ('Hello, PostgreSQL!');")
		if err != nil {
			log.Fatal(err)
		}

		// Write a response
		w.WriteHeader(http.StatusOK)
		fmt.Fprintln(w, "Data written to PostgreSQL")
	})

	http.ListenAndServe(":8080", nil)
}
