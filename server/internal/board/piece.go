package board

import (
	"fmt"
	"strconv"
	"strings"
)

type Piece interface {
	GetColor() string
	GetName() string
	GetPosition() Position
	CanMove(Position, *Board) bool
	Move(Position)
}

func NewPosition(position string) Position {
	pos := strings.Split(position, "")

	x := int(pos[0][0]) - int('a')
	y, err := strconv.Atoi(pos[1])
	if err != nil {
		panic(fmt.Sprintf("could not convert to integer: %s", pos[1]))
	}

	return Position{x: x, y: y}
}

type Position struct {
	x int
	y int
}

func (f *Position) String() string {
	return fmt.Sprintf("%c%d", byte('a'+f.x), f.y)
}
