package server

import (
	"fmt"
	"strconv"
	"strings"
)

func newField(position string) Field {
	pos := strings.Split(position, "")

	x := int(pos[0][0]) - int('a')
	y, err := strconv.Atoi(pos[1])
	if err != nil {
		panic(fmt.Sprintf("could not convert to integer: %s", pos[1]))
	}

	return Field{x: x, y: y}
}

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
		field = newField(string(byte('a'+i)) + "2")
		pieces[field] = Pawn{color: "white"}

		field.y = 7
		pieces[field] = Pawn{color: "black"}
	}

	pieces[newField("a1")] = Rook{color: "white"}
	pieces[newField("h1")] = Rook{color: "white"}

	pieces[newField("a8")] = Rook{color: "black"}
	pieces[newField("h8")] = Rook{color: "black"}

	pieces[newField("b1")] = Knight{color: "white"}
	pieces[newField("g1")] = Knight{color: "white"}

	pieces[newField("b8")] = Knight{color: "black"}
	pieces[newField("g8")] = Knight{color: "black"}

	pieces[newField("c1")] = Bishop{color: "white"}
	pieces[newField("f1")] = Bishop{color: "white"}

	pieces[newField("c8")] = Bishop{color: "black"}
	pieces[newField("f8")] = Bishop{color: "black"}

	pieces[newField("d1")] = Queen{color: "white"}
	pieces[newField("e1")] = King{color: "white"}

	pieces[newField("d8")] = Queen{color: "black"}
	pieces[newField("e8")] = King{color: "black"}

	return Board{
		pieces: pieces,
	}
}
