package server

import (
	. "chessing/internal/pieces"
)

type Board struct {
	pieces []Piece
	turn   string
}

func (b *Board) getPiece(position Position) (int, Piece, bool) {
	for i, piece := range b.pieces {
		if piece.GetPosition() == position {
			return i, piece, true
		}
	}

	return 0, nil, false
}

func (b *Board) removePiece(position Position) {
	i, _, ok := b.getPiece(position)
	if !ok {
		return
	}

	b.pieces = append(b.pieces[:i], b.pieces[i+1:]...)
}

func (b *Board) switchTurn() {
	if b.turn == "white" {
		b.turn = "black"
	} else if b.turn == "black" {
		b.turn = "white"
	}
}

func newBoard() Board {
	pieces := make([]Piece, 0, 32)

	for i := range 8 {
		position := NewPosition(string(byte('a'+i)) + "2")
		pieces = append(pieces, NewPawn("white", position))

		position = NewPosition(string(byte('a'+i)) + "7")
		pieces = append(pieces, NewPawn("black", position))
	}

	pieces = append(pieces, NewRook("white", NewPosition("a1")))
	pieces = append(pieces, NewRook("white", NewPosition("h1")))

	pieces = append(pieces, NewRook("black", NewPosition("a8")))
	pieces = append(pieces, NewRook("black", NewPosition("h8")))

	pieces = append(pieces, NewKnight("white", NewPosition("b1")))
	pieces = append(pieces, NewKnight("white", NewPosition("g1")))

	pieces = append(pieces, NewKnight("black", NewPosition("b8")))
	pieces = append(pieces, NewKnight("black", NewPosition("g8")))

	pieces = append(pieces, NewBishop("white", NewPosition("c1")))
	pieces = append(pieces, NewBishop("white", NewPosition("f1")))

	pieces = append(pieces, NewBishop("black", NewPosition("c8")))
	pieces = append(pieces, NewBishop("black", NewPosition("f8")))

	pieces = append(pieces, NewQueen("white", NewPosition("d1")))
	pieces = append(pieces, NewKing("white", NewPosition("e1")))

	pieces = append(pieces, NewQueen("black", NewPosition("d8")))
	pieces = append(pieces, NewKing("black", NewPosition("e8")))

	return Board{
		pieces: pieces,
		turn:   "white",
	}
}
