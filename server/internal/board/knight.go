package board

import (
	"math"
)

type Knight struct {
	Color    string
	Position Position
}

func NewKnight(color string, position Position) *Knight {
	return &Knight{Color: color, Position: position}
}

func (k *Knight) GetColor() string {
	return k.Color
}

func (k *Knight) GetName() string {
	return "knight"
}

func (k *Knight) GetPosition() Position {
	return k.Position
}

func (k *Knight) CanMove(position Position, board *Board) bool {
	xDiff := int(math.Abs(float64(k.Position.x - position.x)))
	yDiff := int(math.Abs(float64(k.Position.y - position.y)))

	if (xDiff == 2 && yDiff == 1) || (xDiff == 1 && yDiff == 2) && !board.FieldUsed(position) {
		return true
	}

	return false
}

func (k *Knight) Move(position Position) {
	k.Position = position
}
