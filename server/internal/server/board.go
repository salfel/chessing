package server

import (
	"fmt"
)

type Field struct {
	x int
	y int
}

func (f *Field) String() string {
	return fmt.Sprintf("%c%d", byte('a'+f.x), f.y)
}

type Board struct {
	pieces map[Field]Piece
}

func newBoard() Board {
	pieces := make(map[Field]Piece)

	for i := range 8 {
		field := Field{x: i, y: 2}
		pieces[field] = Pawn{color: "white"}

		field.y = 7
		pieces[field] = Pawn{color: "black"}
	}

	pieces[Field{x: 0, y: 1}] = Rook{color: "white"}
	pieces[Field{x: 7, y: 1}] = Rook{color: "white"}

	pieces[Field{x: 0, y: 8}] = Rook{color: "black"}
	pieces[Field{x: 7, y: 8}] = Rook{color: "black"}

	pieces[Field{x: 1, y: 1}] = Knight{color: "white"}
	pieces[Field{x: 6, y: 1}] = Knight{color: "white"}

	pieces[Field{x: 1, y: 8}] = Knight{color: "black"}
	pieces[Field{x: 6, y: 8}] = Knight{color: "black"}

	pieces[Field{x: 2, y: 1}] = Bishop{color: "white"}
	pieces[Field{x: 5, y: 1}] = Bishop{color: "white"}

	pieces[Field{x: 2, y: 8}] = Bishop{color: "black"}
	pieces[Field{x: 5, y: 8}] = Bishop{color: "black"}

	pieces[Field{x: 3, y: 1}] = Queen{color: "white"}
	pieces[Field{x: 4, y: 1}] = King{color: "white"}

	pieces[Field{x: 3, y: 8}] = Queen{color: "black"}
	pieces[Field{x: 4, y: 8}] = King{color: "black"}

	return Board{
		pieces: pieces,
	}
}
