package server

type Piece interface {
	getColor() string
	getName() string
}

type Pawn struct {
	color string
}

func (p Pawn) getColor() string {
	return p.color
}

func (p Pawn) getName() string {
	return "pawn"
}

type Rook struct {
	color string
}

func (r Rook) getColor() string {
	return r.color
}

func (r Rook) getName() string {
	return "rook"
}

type Bishop struct {
	color string
}

func (b Bishop) getColor() string {
	return b.color
}

func (b Bishop) getName() string {
	return "bishop"
}

type Knight struct {
	color string
}

func (k Knight) getColor() string {
	return k.color
}

func (k Knight) getName() string {
	return "knight"
}

type King struct {
	color string
}

func (k King) getColor() string {
	return k.color
}

func (k King) getName() string {
	return "king"
}

type Queen struct {
	color string
}

func (q Queen) getColor() string {
	return q.color
}

func (q Queen) getName() string {
	return "queen"
}
