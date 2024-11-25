package pieces

type Pawn struct {
	Color    string
	Position Position
}

func NewPawn(color string, position Position) Pawn {
	return Pawn{Color: color, Position: position}
}

func (p Pawn) GetColor() string {
	return p.Color
}

func (p Pawn) GetName() string {
	return "pawn"
}

func (p Pawn) GetPosition() Position {
	return p.Position
}

func (p Pawn) CanMove(position Position) bool {
	return true
}
