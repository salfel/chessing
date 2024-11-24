use ratatui::{
    layout::{Constraint, Direction, Layout},
    widgets::{Paragraph, StatefulWidget, Widget},
};

use crate::state::State;

#[derive(Default)]
pub struct Stats {}

impl Stats {
    fn render_empty(
        area: ratatui::prelude::Rect,
        buf: &mut ratatui::prelude::Buffer,
        state: &State,
    ) {
        let layout = Layout::default()
            .direction(Direction::Vertical)
            .constraints(vec![Constraint::Length(2), Constraint::Length(2)])
            .split(area);

        Paragraph::new("Waiting for opponent").render(layout[0], buf);
        Paragraph::new(format!("Code: {}", state.code)).render(layout[1], buf);
    }

    fn render_stateful(
        area: ratatui::prelude::Rect,
        buf: &mut ratatui::prelude::Buffer,
        state: &State,
    ) {
        let layout = Layout::default()
            .direction(Direction::Vertical)
            .constraints(vec![Constraint::Length(2), Constraint::Length(2)])
            .split(area);

        Paragraph::new(format!(
            "Your color: {}",
            state.color.expect("Color not specified")
        ))
        .render(layout[0], buf);

        Paragraph::new(format!("{}'s turn", state.turn)).render(layout[1], buf);
    }
}

impl StatefulWidget for Stats {
    type State = State;

    fn render(
        self,
        area: ratatui::prelude::Rect,
        buf: &mut ratatui::prelude::Buffer,
        state: &mut Self::State,
    ) {
        if state.color.is_none() {
            Self::render_empty(area, buf, state);
        } else {
            Self::render_stateful(area, buf, state);
        }
    }
}
