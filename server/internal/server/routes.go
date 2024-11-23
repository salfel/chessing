package server

import (
	"context"
	_ "embed"
	"encoding/json"
	"log"
	"net/http"
)

func (s *Server) RegisterRoutes() http.Handler {
	mux := http.NewServeMux()

	mux.HandleFunc("/hello", s.HelloWorldHandler)

	mux.HandleFunc("/game/create", s.handlePostRequest(s.CreateGameHandler))

	mux.HandleFunc("/auth/register", s.handlePostRequest(s.RegisterUser))

	return s.LoginMiddleware(mux)
}

func (s *Server) handlePostRequest(handlerFunc http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
			return
		}
		handlerFunc(w, r)
	}
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

func (s *Server) CreateGameHandler(w http.ResponseWriter, r *http.Request) {
	ctx := context.Background()

	username, _, _ := r.BasicAuth()

	game, err := s.queries.CreateGame(ctx)
	if err != nil {
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}

	// _, err := s.queries.create
}
