const ROWS: [u16; 8] = [8, 7, 6, 5, 4, 3, 2, 1];
const COLUMNS: [char; 8] = ['a', 'b', 'c', 'd', 'e', 'f', 'g', 'h'];

pub struct State {
    pub pieces: Vec<Piece>,
}

impl State {
    pub fn new() -> State {
        let mut pieces = Vec::new();

        for y in ROWS.iter().take(2) {
            for x in COLUMNS {}
        }

        State { pieces }
    }
}

pub struct Piece {
    pub color: PieceColor,
    pub r#type: PieceType,
    x: char,
    y: u16,
}

pub enum PieceColor {
    White,
    Black,
}

pub enum PieceType {
    Pawn,
    Bishop,
    Knight,
    Queen,
    King,
    Rook,
}

impl Piece {
    pub fn new(x: char, y: u16) -> Piece {
        match (x, y) {
            (_, 2) => {
                return Piece {
                    color: PieceColor::White,
                    r#type: PieceType::Pawn,
                    x,
                    y,
                }
            }
            (_, 7) => {
                return Piece {
                    color: PieceColor::Black,
                    r#type: PieceType::Pawn,
                    x,
                    y,
                }
            }

            _ => panic!("no piece matched"),
        }
    }
}
