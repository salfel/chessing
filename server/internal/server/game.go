package server

import (
	crypto "crypto/rand"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"math/rand"
	"time"

	. "chessing/internal/board"
)

type Game struct {
	White *Client
	Black *Client
	Board Board
}

func (g *Game) send(message string) {
	g.White.send <- message
	g.Black.send <- message
}

func (c *Client) newGame() *Game {
	source := rand.NewSource(time.Now().UnixNano())
	random := rand.New(source)
	randomValue := random.Intn(2)

	game := Game{Board: NewBoard()}

	if randomValue == 0 {
		game.White = c
	} else {
		game.Black = c
	}

	return &game
}

func (c *Client) leaveGame(game *Game) {
	var opponent *Client

	if c == game.White {
		opponent = game.Black
		game.White = nil
	} else if c == game.Black {
		opponent = game.White
		game.Black = nil
	} else {
		return
	}

	if opponent != nil {
		opponent.send <- "opponent left game:"
	}
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

	game := client.newGame()
	s.games[token] = game
	s.hub.clients[client] = game

	client.send <- fmt.Sprintf("code: %s", token)
}

func (s *Server) joinGame(message string, client *Client) {
	game, ok := s.games[message]
	if !ok {
		client.send <- "not found"
		return
	}

	if game.Black == nil {
		game.Black = client
	} else if game.White == nil {
		game.White = client
	} else {
		client.send <- "game full"
		return
	}

	s.hub.clients[client] = game

	game.White.send <- "color: white"
	game.Black.send <- "color: black"
	game.send("starting:")

	s.sendState(game)
}

func (s *Server) sendState(game *Game) {
	pieces := map[string]map[string]string{}

	for _, piece := range game.Board.Pieces {
		position := piece.GetPosition()
		pieces[position.String()] = map[string]string{
			"color":   piece.GetColor(),
			"variant": piece.GetName(),
		}
	}

	jsonPieces, err := json.Marshal(pieces)
	if err != nil {
		fmt.Println(err)
		return
	}

	game.send(fmt.Sprintf("state: %s", jsonPieces))
}
