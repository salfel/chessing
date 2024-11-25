package pieces

type Queen struct {
	Color    string
	Position Position
}

func NewQueen(color string, position Position) Queen {
	return Queen{Color: color, Position: position}
}

func (q Queen) GetColor() string {
	return q.Color
}

func (q Queen) GetName() string {
	return "queen"
}

func (q Queen) GetPosition() Position {
	return q.Position
}

func (q Queen) CanMove(position Position) bool {
	return true
}
