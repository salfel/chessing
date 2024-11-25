package pieces

type King struct {
	Color    string
	Position Position
}

func NewKing(color string, position Position) King {
	return King{Color: color, Position: position}
}

func (k King) GetColor() string {
	return k.Color
}

func (k King) GetName() string {
	return "king"
}

func (k King) GetPosition() Position {
	return k.Position
}

func (k King) CanMove(position Position) bool {
	return true
}
