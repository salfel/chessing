package server

type Hub struct {
	server *Server

	clients map[*Client]*Game

	broadcast  chan Message
	register   chan *Client
	unregister chan *Client
}

type Message struct {
	content string
	client  *Client
}

func newHub(server *Server) *Hub {
	return &Hub{
		broadcast:  make(chan Message),
		register:   make(chan *Client),
		unregister: make(chan *Client),
		clients:    make(map[*Client]*Game),
		server:     server,
	}
}

func (h *Hub) run() {
	for {
		select {
		case client := <-h.register:
			h.clients[client] = nil
		case client := <-h.unregister:
			if _, ok := h.clients[client]; ok {
				client.leaveGame(h.clients[client])

				delete(h.clients, client)
				close(client.send)
			}
		case message := <-h.broadcast:
			h.server.RouteWebsockets(message)
		}
	}
}
