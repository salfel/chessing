use ratatui::{
    layout::{Constraint, Direction, Layout},
    prelude::{Buffer, Rect},
    style::Stylize,
    text::Span,
    widgets::{Block, Paragraph, StatefulWidget, Widget},
};

use crate::state::State;

pub struct Board {}

enum BoardColor {
    Black,
    White,
}

impl BoardColor {
    fn switch(self) -> BoardColor {
        match self {
            BoardColor::Black => BoardColor::White,
            BoardColor::White => BoardColor::Black,
        }
    }
}

impl Board {
    pub fn new() -> Board {
        Board {}
    }
}

impl Board {
    fn render_board<'a>(self, area: Rect, buf: &mut Buffer, state: &'a mut State) {
        let mut color = BoardColor::White;

        for y in 0..8 {
            for x in 0..8 {
                let area = Rect::new(area.x + x * 3 + 2, area.y + y, 3, 1);
                let block = match color {
                    BoardColor::Black => Block::new().on_black(),
                    BoardColor::White => Block::new().on_white(),
                };

                block.render(area, buf);

                color = color.switch();
            }

            let area = Rect::new(area.x, area.y + y, 1, 1);
            Span::from((8 - y).to_string()).render(area, buf);

            color = color.switch();
        }

        for x in 0..8 {
            let area = Rect::new(area.x + x * 3 + 2, area.y + 8, 3, 1);
            let char = char::from_u32((x + 65) as u32).unwrap();

            Span::from(format!(" {} ", char)).render(area, buf);
        }
    }

    fn render_waiting<'a>(self, area: Rect, buf: &mut Buffer, state: &'a mut State) {
        let layout = Layout::default()
            .direction(Direction::Vertical)
            .constraints(vec![Constraint::Length(3), Constraint::Length(3)])
            .split(area);

        Paragraph::new("Waiting for opponent")
            .centered()
            .render(layout[0], buf);
        Paragraph::new(format!("Code: {}", state.code))
            .centered()
            .render(layout[1], buf);
    }
}

impl StatefulWidget for Board {
    type State = State;

    fn render<'a>(self, area: Rect, buf: &mut Buffer, state: &'a mut Self::State) {
        if state.pieces.is_some() {
            self.render_board(area, buf, state);
        } else {
            self.render_waiting(area, buf, state);
        }
    }
}
