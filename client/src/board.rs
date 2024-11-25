use ratatui::{
    prelude::{Buffer, Rect},
    style::Stylize,
    text::Span,
    widgets::{Block, StatefulWidget, Widget},
};

use crate::state::{Color, State, Status};

#[derive(Default)]
pub struct Board {}

impl StatefulWidget for Board {
    type State = State;

    fn render<'a>(self, area: Rect, buf: &mut Buffer, state: &'a mut Self::State) {
        let mut color = state.color.unwrap_or(Color::White);

        for y in 0..8 {
            for x in 0..8 {
                let mut area = Rect::new(area.x + x * 3 + 2, area.y + y, 3, 1);
                let block = match color {
                    Color::Black => Block::new().on_black(),
                    Color::White => Block::new().on_white(),
                };

                block.render(area, buf);

                if state.status != Status::Waiting {
                    let number = match state.color {
                        Some(Color::White) => 8 - y,
                        Some(Color::Black) => y + 1,
                        None => panic!("color is none"),
                    };

                    let char = char::from_u32(x as u32 + 97).expect("Isn't a valid char");
                    match state.get_piece(char, number) {
                        Some(char) => {
                            area.x += 1;
                            area.width = 1;
                            Span::from(char).render(area, buf);
                        }
                        None => {}
                    }
                }

                color = color.switch();
            }

            if state.status != Status::Waiting {
                let number = match state.color {
                    Some(Color::White) => 8 - y,
                    Some(Color::Black) => y + 1,
                    None => panic!("color is none"),
                };

                let area = Rect::new(area.x, area.y + y, 1, 1);
                Span::from(number.to_string()).render(area, buf);
            }

            color = color.switch();
        }

        if state.status != Status::Waiting {
            for x in 0..8 {
                let area = Rect::new(area.x + x * 3 + 2, area.y + 8, 3, 1);
                let char = char::from_u32((x + 65) as u32).unwrap();

                Span::from(format!(" {} ", char)).render(area, buf);
            }
        }
    }
}
