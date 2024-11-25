package board

type Board struct {
	Pieces []Piece
	Turn   string
}

func (b *Board) SwitchTurn() {
	if b.Turn == "white" {
		b.Turn = "black"
	} else if b.Turn == "black" {
		b.Turn = "white"
	}
}

func NewBoard() Board {
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
		Pieces: pieces,
		Turn:   "white",
	}
}
