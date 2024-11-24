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
		pieces[field] = Piece{color: WHITE, variant: PAWN}

		field.y = 7
		pieces[field] = Piece{color: BLACK, variant: PAWN}
	}

	pieces[Field{x: 0, y: 1}] = Piece{color: WHITE, variant: ROOK}
	pieces[Field{x: 7, y: 1}] = Piece{color: WHITE, variant: ROOK}

	pieces[Field{x: 0, y: 8}] = Piece{color: BLACK, variant: ROOK}
	pieces[Field{x: 7, y: 8}] = Piece{color: BLACK, variant: ROOK}

	pieces[Field{x: 1, y: 1}] = Piece{color: WHITE, variant: KNIGHT}
	pieces[Field{x: 6, y: 1}] = Piece{color: WHITE, variant: KNIGHT}

	pieces[Field{x: 1, y: 8}] = Piece{color: BLACK, variant: KNIGHT}
	pieces[Field{x: 6, y: 8}] = Piece{color: BLACK, variant: KNIGHT}

	pieces[Field{x: 2, y: 1}] = Piece{color: WHITE, variant: BISHOP}
	pieces[Field{x: 5, y: 1}] = Piece{color: WHITE, variant: BISHOP}

	pieces[Field{x: 2, y: 8}] = Piece{color: BLACK, variant: BISHOP}
	pieces[Field{x: 5, y: 8}] = Piece{color: BLACK, variant: BISHOP}

	pieces[Field{x: 3, y: 1}] = Piece{color: WHITE, variant: QUEEN}
	pieces[Field{x: 4, y: 1}] = Piece{color: WHITE, variant: KING}

	pieces[Field{x: 3, y: 8}] = Piece{color: BLACK, variant: QUEEN}
	pieces[Field{x: 4, y: 8}] = Piece{color: BLACK, variant: KING}

	return Board{
		pieces: pieces,
	}
}
