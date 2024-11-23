package server

import (
	"chessing/internal/database"
	"database/sql"
	_ "embed"
	"encoding/json"
	"fmt"
	"log"
	"math/rand"
	"net/http"
)

func (s *Server) RegisterRoutes() http.Handler {
	mux := http.NewServeMux()

	mux.HandleFunc("/hello", s.HelloWorldHandler)

	mux.HandleFunc("/game/create", s.handlePostRequest(s.CreateGameHandler))
	mux.HandleFunc("/game/join", s.handlePostRequest(s.JoinGameHandler))

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
	ctx := r.Context()

	user := r.Context().Value(userContextKey).(database.User)

	random := rand.Intn(2)

	var createGameParams database.CreateGameParams
	var color string
	if random == 0 {
		createGameParams = database.CreateGameParams{Black: sql.NullInt64{Valid: false}, White: sql.NullInt64{Int64: user.ID, Valid: true}}
		color = "white"
	} else {
		createGameParams = database.CreateGameParams{Black: sql.NullInt64{Int64: user.ID, Valid: true}, White: sql.NullInt64{Valid: false}}
		color = "black"
	}

	err := s.queries.CreateGame(ctx, createGameParams)
	if err != nil {
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}

	res := map[string]string{"color": color}
	jsonRes, err := json.Marshal(res)
	if err != nil {
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}

	if _, err := w.Write(jsonRes); err != nil {
		fmt.Printf("Couldn't write output: %v", err.Error())
	}
}

func (s *Server) JoinGameHandler(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	var req struct {
		GameId int64 `json:"gameId"`
	}

	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}

	game, err := s.queries.GetGame(ctx, req.GameId)
	if err != nil {
		fmt.Println(err)
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}

	user := r.Context().Value(userContextKey).(database.User)

	joinGameParams := database.JoinGameParams{ID: game.ID, White: game.White, Black: game.Black}
	if game.White.Valid && game.Black.Valid {
		fmt.Println(err)
		http.Error(w, "Game is already full", http.StatusBadRequest)
		return
	} else if !game.White.Valid {
		joinGameParams.White = sql.NullInt64{Int64: user.ID, Valid: true}
	} else {
		joinGameParams.Black = sql.NullInt64{Int64: user.ID, Valid: true}
	}

	err = s.queries.JoinGame(ctx, joinGameParams)
	if err != nil {
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}

	res := map[string]string{"message": "Successfully joined the game"}
	jsonRes, err := json.Marshal(res)
	if err != nil {
		fmt.Println(err)
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	if _, err := w.Write(jsonRes); err != nil {
		log.Printf("Failed to write response: %v", err)
	}
}
