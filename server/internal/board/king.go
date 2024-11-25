package board

type King struct {
	Color    string
	Position Position
}

func NewKing(color string, position Position) *King {
	return &King{Color: color, Position: position}
}

func (k *King) GetColor() string {
	return k.Color
}

func (k *King) GetName() string {
	return "king"
}

func (k *King) GetPosition() Position {
	return k.Position
}

func (k *King) CanMove(position Position, board *Board) bool {
	if board.FieldUsed(position) {
		return false
	}

	xDiff := Abs(k.Position.x - position.x)
	yDiff := Abs(k.Position.y - position.y)

	if (xDiff == 1 && yDiff <= 1) || (yDiff == 1 && xDiff <= 1) {
		return true
	}

	return false
}

func (k *King) Move(position Position) {
	k.Position = position
}
