package board

type Rook struct {
	Color    string
	Position Position
}

func NewRook(color string, position Position) *Rook {
	return &Rook{Color: color, Position: position}
}

func (r *Rook) GetColor() string {
	return r.Color
}

func (r *Rook) GetName() string {
	return "rook"
}

func (r *Rook) GetPosition() Position {
	return r.Position
}

func (r *Rook) CanMove(position Position, board *Board) bool {
	return board.IsEmptyLine(r.Position, position)
}

func (r *Rook) Move(position Position) {
	r.Position = position
}
