use std::collections::HashMap;

pub struct State {
    pub pieces: Option<HashMap<Position, char>>,
    pub color: Option<Color>,
    pub code: String,
}

impl State {
    pub fn new() -> State {
        State {
            pieces: None,
            color: None,
            code: String::new(),
        }
    }

    pub fn parse_pieces(&mut self, pieces: &str) {
        let data = Self::from_json(pieces).unwrap();
        let mut pieces = HashMap::new();

        for (position, char) in data.into_iter() {
            let char = char
                .as_str()
                .expect("Expected a string as piece char")
                .chars()
                .next()
                .expect("Expected at least one char");

            pieces.insert(Position::new(position), char);
        }

        self.pieces = Some(pieces);
    }

    fn from_json(json: &str) -> Result<HashMap<String, serde_json::Value>, serde_json::Error> {
        serde_json::from_str(json)
    }

    pub fn get_piece(&self, x: char, y: u16) -> Option<String> {
        let position = format!("{}{}", x, y);

        self.pieces
            .as_ref()
            .expect("Hashmap not populated")
            .get(&Position::new(position))
            .map(|char| char.to_string())
    }
}

#[derive(Hash, Eq, PartialEq)]
pub struct Position {
    pub x: u16,
    pub y: u16,
}

impl Position {
    pub fn new(position: String) -> Position {
        let mut positions = position.chars();
        let x = positions.next().expect("x position not present");
        let y = positions.next().expect("y position not present");

        let x = (x as u16) - ('A' as u16) + 1;
        let y = y.to_digit(10).expect("y position is not a integer") as u16;

        Position { x, y }
    }
}

#[derive(Clone, Copy)]
pub enum Color {
    White,
    Black,
}

impl Color {
    pub fn new(color: &str) -> Color {
        match color {
            "white" => Color::White,
            "black" => Color::Black,
            _ => panic!("color shouldn't be anything else than black or white"),
        }
    }

    pub fn switch(self) -> Color {
        match self {
            Color::Black => Color::White,
            Color::White => Color::Black,
        }
    }
}
