package board

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

func (k *Knight) CanMove(position Position) bool {
	return false
}

func (k *Knight) Move(position Position) {
	k.Position = position
}
