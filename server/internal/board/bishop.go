package board

type Bishop struct {
	Color    string
	Position Position
}

func NewBishop(color string, position Position) *Bishop {
	return &Bishop{Color: color, Position: position}
}

func (b *Bishop) GetColor() string {
	return b.Color
}

func (b *Bishop) GetName() string {
	return "bishop"
}

func (b *Bishop) GetPosition() Position {
	return b.Position
}

func (b *Bishop) CanMove(position Position, board *Board) bool {
	return false
}

func (b *Bishop) Move(position Position) {
	b.Position = position
}
