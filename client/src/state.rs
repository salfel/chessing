const ROWS: [u16; 8] = [8, 7, 6, 5, 4, 3, 2, 1];
const COLUMNS: [char; 8] = ['a', 'b', 'c', 'd', 'e', 'f', 'g', 'h'];

pub struct State {
    pub pieces: Vec<String>,
}

impl State {
    pub fn new() -> State {
        let pieces = Vec::new();

        State { pieces }
    }
}

pub struct Piece {
    pub color: PieceColor,
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
