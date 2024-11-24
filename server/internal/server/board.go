package server

import "fmt"

type Field struct {
	x int
	y int
}

func (f *Field) String() string {
	return fmt.Sprintf("%c%d", byte('A'+f.x), f.y)
}

type Board struct {
	pieces map[Field]Piece
}

func newBoard() Board {
	pieces := make(map[Field]Piece)

	for i := range 8 {
		field := Field{x: i, y: 2}
		pieces[field] = WHITEPAWN

		field.y = 7
		pieces[field] = BLACKPAWN
	}

	pieces[Field{x: 0, y: 1}] = WHITEROOK
	pieces[Field{x: 7, y: 1}] = WHITEROOK

	pieces[Field{x: 0, y: 8}] = BLACKROOK
	pieces[Field{x: 7, y: 8}] = BLACKROOK

	pieces[Field{x: 1, y: 1}] = WHITEKNIGHT
	pieces[Field{x: 6, y: 1}] = WHITEKNIGHT

	pieces[Field{x: 1, y: 8}] = BLACKKNIGHT
	pieces[Field{x: 6, y: 8}] = BLACKKNIGHT

	pieces[Field{x: 2, y: 1}] = WHITEBISHOP
	pieces[Field{x: 5, y: 1}] = WHITEBISHOP

	pieces[Field{x: 2, y: 8}] = BLACKBISHOP
	pieces[Field{x: 5, y: 8}] = BLACKBISHOP

	pieces[Field{x: 3, y: 1}] = WHITEQUEEN
	pieces[Field{x: 4, y: 1}] = WHITEKING

	pieces[Field{x: 3, y: 8}] = BLACKQUEEN
	pieces[Field{x: 4, y: 8}] = BLACKKING

	return Board{
		pieces: pieces,
	}
}
