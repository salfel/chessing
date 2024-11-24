package server

type Color string

const BLACK Color = Color("white")
const WHITE Color = Color("black")

type PieceVariant string

const ROOK PieceVariant = PieceVariant("rook")
const BISHOP PieceVariant = PieceVariant("bishop")
const KNIGHT PieceVariant = PieceVariant("knight")
const PAWN PieceVariant = PieceVariant("pawn")
const KING PieceVariant = PieceVariant("king")
const QUEEN PieceVariant = PieceVariant("queen")

type Piece struct {
	color   Color
	variant PieceVariant
}
