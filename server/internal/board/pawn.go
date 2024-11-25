package board

type Pawn struct {
	Color    string
	Position Position
}

func NewPawn(color string, position Position) *Pawn {
	return &Pawn{Color: color, Position: position}
}

func (p *Pawn) GetColor() string {
	return p.Color
}

func (p *Pawn) GetName() string {
	return "pawn"
}

func (p *Pawn) GetPosition() Position {
	return p.Position
}

func (p *Pawn) CanMove(position Position) bool {
	if p.Position.x != position.x {
		return false
	}

	if p.Color == "white" {
		if p.Position.y == position.y-1 {
			return true
		} else if p.Position.y == position.y-2 && p.Position.y == 2 {
			return true
		}
	} else if p.Color == "black" {
		if p.Position.y == position.y+1 {
			return true
		} else if p.Position.y == position.y+2 && p.Position.y == 7 {
			return true
		}
	}

	return false
}

func (p *Pawn) Move(position Position) {
	p.Position = position
}
