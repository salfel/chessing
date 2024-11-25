use std::cmp::max;

use ratatui::{
    layout::{Alignment, Constraint, Flex, Layout, Rect},
    widgets::{Block, Clear, Paragraph, StatefulWidget, Widget},
};

use crate::state::State;

pub struct Input {
    label: String,
}

impl Input {
    pub fn new(label: String) -> Input {
        Input { label }
    }
}

impl StatefulWidget for Input {
    type State = State;

    fn render(
        self,
        area: ratatui::prelude::Rect,
        buf: &mut ratatui::prelude::Buffer,
        state: &mut Self::State,
    ) {
        let vertical = Constraint::Length(3);
        let horizontal = Constraint::Max(max(20, state.input.len() as u16));

        let area = center(area, horizontal, vertical);

        Clear.render(area, buf);
        Paragraph::new(state.input.clone())
            .block(
                Block::bordered()
                    .title(self.label)
                    .title_alignment(Alignment::Center),
            )
            .render(area, buf);
    }
}

fn center(area: Rect, horizontal: Constraint, vertical: Constraint) -> Rect {
    let [area] = Layout::vertical([vertical]).flex(Flex::Center).areas(area);
    let [area] = Layout::horizontal([horizontal])
        .flex(Flex::Center)
        .areas(area);

    area
}
