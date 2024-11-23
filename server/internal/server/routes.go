package server

import (
	_ "embed"
	"encoding/json"
	"log"
	"net/http"
)

func (s *Server) RegisterRoutes() http.Handler {
	mux := http.NewServeMux()

	mux.HandleFunc("/hello", s.HelloWorldHandler)

	mux.HandleFunc("/migration", s.MigrationHandler)

	return mux
}

func (s *Server) HelloWorldHandler(w http.ResponseWriter, r *http.Request) {
	res := map[string]string{"message": "Hello world"}
	jsonRes, err := json.Marshal(res)
	if err != nil {
		http.Error(w, "Failed to marshall response", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	if _, err := w.Write(jsonRes); err != nil {
		log.Printf("Failed to write response: %v", err)
	}
}

func (s *Server) MigrationHandler(w http.ResponseWriter, r *http.Request) {
	if err := s.queries.Migrate(); err != nil {
		http.Error(w, "Failed to migrate", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/text")
	if _, err := w.Write([]byte("success")); err != nil {
		log.Fatalf("Failed to write response: %v", err)
	}
}
