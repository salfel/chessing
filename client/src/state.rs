use std::{collections::HashMap, fmt::Display};

pub struct State {
    pub pieces: HashMap<Position, char>,
    pub color: Option<Color>,
    pub code: String,
    pub turn: Color,
    pub status: Status,
    pub should_quit: bool,
}

impl State {
    pub fn new() -> State {
        State {
            pieces: HashMap::new(),
            color: None,
            code: String::new(),
            turn: Color::White,
            status: Status::Waiting,
            should_quit: false,
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

        self.pieces = pieces;
    }

    fn from_json(json: &str) -> Result<HashMap<String, serde_json::Value>, serde_json::Error> {
        serde_json::from_str(json)
    }

    pub fn get_piece(&self, x: char, y: u16) -> Option<String> {
        let position = format!("{}{}", x, y);

        self.pieces
            .get(&Position::new(position))
            .map(|char| char.to_string())
    }
}

#[derive(PartialEq, Eq)]
pub enum Status {
    Waiting,
    Playing,
    Leaving,
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

impl Display for Color {
    fn fmt(&self, f: &mut std::fmt::Formatter<'_>) -> std::fmt::Result {
        let color_str = match self {
            Color::White => "white",
            Color::Black => "black",
        };
        write!(f, "{}", color_str)
    }
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
