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

func (b *Board) FieldUsed(field Position) bool {
	for _, piece := range b.Pieces {
		if field == piece.GetPosition() {
			return true
		}
	}

	return false
}

func (b *Board) HasEmptyFields(fields []Position) bool {
	empty := true

	for _, field := range fields {
		for _, piece := range b.Pieces {
			if piece.GetPosition() == field {
				empty = false
			}
		}
	}

	return empty
}

func (b *Board) IsEmptyLine(current Position, destination Position) bool {
	if current.x == destination.x {
		min := min(current.y, destination.y)
		max := max(current.y, destination.y)

		fields := make([]Position, 0, max-min)

		for i := min; i <= max; i++ {
			field := Position{x: current.x, y: i}

			if b.FieldUsed(field) && field != current {
				fields = append(fields, field)
			}
		}

		return len(fields) == 0
	} else if current.y == destination.y {
		min := min(current.x, destination.x)
		max := max(current.x, destination.x)

		fields := make([]Position, 0, max-min)

		for i := min; i <= max; i++ {
			field := Position{x: i, y: current.y}

			if b.FieldUsed(field) && field != current {
				fields = append(fields, field)
			}
		}

		return len(fields) == 0
	}

	return false
}

func (b *Board) IsEmptyDiagonal(original Position, destination Position) bool {
	if Abs(original.x-destination.x) != Abs(original.y-destination.y) {
		return false
	}

	var xDiff, yDiff int

	if original.x < destination.x {
		xDiff = 1
	} else {
		xDiff = -1
	}

	if original.y < destination.y {
		yDiff = 1
	} else {
		yDiff = -1
	}

	fields := make([]Position, 0, Abs(original.x-destination.x))

	position := original

	for position != destination {
		position.x += xDiff
		position.y += yDiff

		if b.FieldUsed(position) {
			fields = append(fields, position)
		}
	}

	return len(fields) == 0
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

func Abs(x int) int {
	if x < 0 {
		return -x
	} else {
		return x
	}
}
