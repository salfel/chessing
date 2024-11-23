package server

import (
	"net/http"
)

func (s *Server) RegisterRoutes() http.Handler {
	mux := http.NewServeMux()

	mux.HandleFunc("/game", func(w http.ResponseWriter, r *http.Request) {
		serveWs(s.hub, w, r)
	})

	return mux
}
