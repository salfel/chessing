package server

import (
	. "chessing/internal/board"
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
	move := parseMove(message, game.Board.Turn)

	if game.Board.Turn == "white" && game.White != client || game.Board.Turn == "black" && game.Black != client {
		return
	}

	found := false

	for _, piece := range game.Board.Pieces {
		if move.GetColor() == piece.GetColor() && piece.GetName() == move.GetName() && piece.CanMove(move.GetPosition(), &game.Board) {
			piece.Move(move.GetPosition())
			found = true
		}
	}

	if found {
		s.sendState(game)
		game.Board.SwitchTurn()
		game.send(fmt.Sprintf("turn: %s", game.Board.Turn))
	}
}
