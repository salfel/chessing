package server

import (
	"net/http"
	"strings"
)

func (s *Server) RegisterRoutes() http.Handler {
	mux := http.NewServeMux()

	mux.HandleFunc("/game", func(w http.ResponseWriter, r *http.Request) {
		serveWs(s.hub, w, r)
	})

	return mux
}

type WebsocketRouter struct {
	client  *Client
	message string
}

type WebsocketHandlerFunc = func(string, *Client)

func (s *Server) RouteWebsockets(message Message) {
	r := WebsocketRouter{client: message.client, message: message.content}

	r.HandleFunc("create game:", s.createGame)
}

func (r *WebsocketRouter) HandleFunc(path string, handler WebsocketHandlerFunc) {
	if strings.HasPrefix(r.message, path) {
		handler(strings.TrimPrefix(r.message, path), r.client)
	}
}

func (s *Server) createGame(message string, client *Client) {
	client.send <- []byte(message)
}
