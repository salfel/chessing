package server

import (
	. "chessing/internal/pieces"
	"fmt"
)

func parseMove(move string, color string) (Piece, Position) {
	var piece Piece

	fmt.Println(move)

	switch move[0] {
	case 'R':
		piece = Rook{Color: color}
	case 'N':
		piece = Knight{Color: color}
	case 'B':
		piece = Bishop{Color: color}
	case 'Q':
		piece = Queen{Color: color}
	case 'K':
		piece = King{Color: color}
	default:
		piece = Pawn{Color: color}
		return piece, NewPosition(move)
	}

	position := NewPosition(move[1:])

	return piece, position
}

func (s *Server) movePiece(message string, client *Client) {
	piece, field := parseMove(message, "white")

	client.send <- fmt.Sprint(piece.GetName(), piece, field)
}
