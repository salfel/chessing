package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"

	. "chessing/internal/server"
	_ "github.com/mattn/go-sqlite3"
)

func main() {
	db, err := sql.Open("sqlite3", "./database.sqlite")
	if err != nil {
		log.Fatalf("Couldn't open database: %v", err)
		return
	}

	server := NewServer(db)

	fmt.Println("Server started")
	err = server.ListenAndServe()
	if err != nil && err != http.ErrServerClosed {
		panic(fmt.Sprintf("http server error: %s", err))
	}
}
