const ROWS: [u16; 8] = [8, 7, 6, 5, 4, 3, 2, 1];
const COLUMNS: [char; 8] = ['a', 'b', 'c', 'd', 'e', 'f', 'g', 'h'];

pub struct State {
    pub pieces: Option<Vec<Piece>>,
}

impl State {
    pub fn new() -> State {
        State { pieces: None }
    }
}

pub struct Piece {
    pub char: char,
    pub x: char,
    pub y: u16,
}
