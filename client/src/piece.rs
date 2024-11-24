pub struct Piece {
    color: Color,
    variant: Variant,
}

impl Piece {
    pub fn new(color: &str, variant: &str) -> Piece {
        Piece {
            color: Color::new(color),
            variant: Variant::new(variant),
        }
    }

    pub fn to_char(&self) -> char {
        match (&self.color, &self.variant) {
            (Color::Black, Variant::Rook) => '♖',
            (Color::Black, Variant::Bishop) => '♗',
            (Color::Black, Variant::Queen) => '♕',
            (Color::Black, Variant::King) => '♔',
            (Color::Black, Variant::Pawn) => '♙',
            (Color::Black, Variant::Knight) => '♘',
            (Color::White, Variant::Pawn) => '♟',
            (Color::White, Variant::Knight) => '♞',
            (Color::White, Variant::Rook) => '♜',
            (Color::White, Variant::Bishop) => '♝',
            (Color::White, Variant::Queen) => '♛',
            (Color::White, Variant::King) => '♚',
        }
    }
}

enum Variant {
    Pawn,
    Rook,
    Knight,
    Bishop,
    Queen,
    King,
}

impl Variant {
    pub fn new(variant: &str) -> Variant {
        match variant {
            "pawn" => Variant::Pawn,
            "rook" => Variant::Rook,
            "knight" => Variant::Knight,
            "bishop" => Variant::Bishop,
            "queen" => Variant::Queen,
            "king" => Variant::King,
            _ => panic!("no valid variant"),
        }
    }
}

enum Color {
    Black,
    White,
}

impl Color {
    pub fn new(color: &str) -> Color {
        match color {
            "black" => Color::Black,
            "white" => Color::White,
            _ => panic!("{} is not a supported color", color),
        }
    }
}
