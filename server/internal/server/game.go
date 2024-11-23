package server

import (
	crypto "crypto/rand"
	"encoding/hex"
	"math/rand"
	"time"
)

type Game struct {
	White *Client
	Black *Client
	Board Board
}

func newGame(client *Client) *Game {
	source := rand.NewSource(time.Now().UnixNano())
	random := rand.New(source)
	randomValue := random.Intn(2)

	game := Game{Board: newBoard()}

	if randomValue == 0 {
		game.White = client
	} else {
		game.Black = client
	}

	return &game
}

func (g *Game) getOpponent(client *Client) *Client {
	if client == g.Black {
		return g.White
	} else if client == g.White {
		return g.Black
	} else {
		panic("client is not in that game")
	}
}

func (s *Server) createGame(message string, client *Client) {
	randomBytes := make([]byte, 4)
	_, err := crypto.Read(randomBytes)
	if err != nil {
		return
	}

	token := hex.EncodeToString(randomBytes)

	s.games[token] = newGame(client)

	client.send <- []byte(token)
}

func (s *Server) joinGame(message string, client *Client) {
	game, ok := s.games[message]
	if !ok {
		client.send <- []byte("not found")
		return
	}

	if game.Black == nil {
		game.Black = client
	} else if game.White == nil {
		game.White = client
	} else {
		client.send <- []byte("game full")
		return
	}

	opponent := game.getOpponent(client)
	opponent.send <- []byte("somebody joined your game")

	client.send <- []byte("joined gamed")
}
