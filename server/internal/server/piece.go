package server

type Color string

const BLACK Color = Color("white")
const WHITE Color = Color("black")

type Piece struct {
	color   Color
	variant rune
}

var BLACKKING = Piece{
	color:   BLACK,
	variant: '\u2654',
}

var BLACKQUEEN = Piece{
	color:   BLACK,
	variant: '\u2655',
}

var BLACKROOK = Piece{
	color:   BLACK,
	variant: '\u2656',
}

var BLACKBISHOP = Piece{
	color:   BLACK,
	variant: '\u2657',
}

var BLACKKNIGHT = Piece{
	color:   BLACK,
	variant: '\u2658',
}

var BLACKPAWN = Piece{
	color:   BLACK,
	variant: '\u2659',
}

var WHITEKING = Piece{
	color:   WHITE,
	variant: '\u265A',
}

var WHITEQUEEN = Piece{
	color:   WHITE,
	variant: '\u265B',
}

var WHITEROOK = Piece{
	color:   WHITE,
	variant: '\u265C',
}

var WHITEBISHOP = Piece{
	color:   WHITE,
	variant: '\u265D',
}

var WHITEKNIGHT = Piece{
	color:   WHITE,
	variant: '\u265E',
}

var WHITEPAWN = Piece{
	color:   WHITE,
	variant: '\u265F',
}
