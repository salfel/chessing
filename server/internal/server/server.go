package server

import (
	"database/sql"
	"fmt"
	"net/http"

	"chessing/internal/database"
)

type Server struct {
	port    int
	db      *sql.DB
	queries *database.Queries
	games   map[string]Game
	hub     *Hub
}

func NewServer(db *sql.DB) *http.Server {
	newServer := Server{
		port:    8000,
		db:      db,
		queries: database.New(db),
		games:   map[string]Game{},
	}

	hub := newHub(&newServer)
	newServer.hub = hub

	go hub.run()

	server := &http.Server{
		Addr:    fmt.Sprintf(":%d", newServer.port),
		Handler: newServer.RegisterRoutes(),
	}

	return server
}
