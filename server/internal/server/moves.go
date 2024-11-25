package server

import (
	. "chessing/internal/pieces"
	"fmt"
)

func parseMove(move string, color string) Piece {
	switch move[0] {
	case 'R':
		return NewRook(color, NewPosition(move[1:]))
	case 'N':
		return NewKnight(color, NewPosition(move[1:]))
	case 'B':
		return NewBishop(color, NewPosition(move[1:]))
	case 'Q':
		return NewQueen(color, NewPosition(move[1:]))
	case 'K':
		return NewKing(color, NewPosition(move[1:]))
	default:
		return NewPawn(color, NewPosition(move))
	}
}

func (s *Server) movePiece(message string, client *Client) {
	game := s.hub.clients[client]
	move := parseMove(message, game.Board.turn)

	found := false

	for _, piece := range game.Board.pieces {
		if piece.CanMove(move.GetPosition()) && move.GetColor() == piece.GetColor() {
			piece.Move(move.GetPosition())
			found = true
		}
	}

	if found {
		s.sendState(game)
		game.Board.switchTurn()
		game.send(fmt.Sprintf("turn: %s", game.Board.turn))
	}
}
