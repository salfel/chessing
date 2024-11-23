use ratatui::{
    prelude::{Buffer, Rect},
    style::Stylize,
    text::Span,
    widgets::{Block, StatefulWidget, Widget},
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

impl StatefulWidget for Board {
    type State = State;

    fn render(self, area: Rect, buf: &mut Buffer, state: &mut Self::State) {}

    //fn render(self, area: Rect, buf: &mut Buffer) {
    //    let mut color = BoardColor::White;
    //
    //    for y in 0..8 {
    //        for x in 1..9 {
    //            let area = Rect::new(area.x + x, area.y + y, 1, 1);
    //            let block = match color {
    //                BoardColor::Black => Block::new().on_black(),
    //                BoardColor::White => Block::new().on_white(),
    //            };
    //
    //            block.render(area, buf);
    //
    //            color = color.switch();
    //        }
    //
    //        let area = Rect::new(area.x, area.y + y, 1, 1);
    //        Span::from((8 - y).to_string()).render(area, buf);
    //
    //        color = color.switch();
    //    }
    //
    //    for x in 1..9 {
    //        let area = Rect::new(area.x + x, area.y + 8, 1, 1);
    //        let char = char::from_u32((x + 64) as u32).unwrap();
    //
    //        Span::from(char.to_string()).render(area, buf);
    //    }
    //}
}
