use futures_util::SinkExt;
use ratatui::{
    prelude::{Buffer, Rect, Stylize},
    widgets::{Paragraph, StatefulWidget, Widget},
};
use tokio_tungstenite::tungstenite::Message;

use crate::{
    app::SocketStream,
    state::{State, Status},
};

#[derive(Clone, Copy)]
pub enum Action {
    Creating,
    Joining,
}

impl Action {
    pub async fn execute(&self, state: &mut State, socket: &mut SocketStream) {
        match self {
            Action::Creating => {
                socket
                    .writer
                    .send(Message::text("create game: test"))
                    .await
                    .expect("Failed to send message");
            }
            Action::Joining => socket
                .writer
                .send(Message::text(format!("join game: {}", state.input)))
                .await
                .expect("Failed to send message"),
        }
    }

    pub fn label(&self) -> String {
        match self {
            Action::Creating => "create game",
            Action::Joining => "join game",
        }
        .to_string()
    }

    pub fn keybind(&self) -> char {
        match self {
            Action::Creating => 'c',
            Action::Joining => 'j',
        }
    }
}

#[derive(Default)]
pub struct Actions {}

impl StatefulWidget for Actions {
    type State = State;

    fn render(self, area: Rect, buf: &mut Buffer, state: &mut Self::State) {
        let actions = match state.status {
            Status::Waiting => {
                let mut actions = vec![];
                if state.code.is_empty() {
                    actions.push(Action::Creating);
                    actions.push(Action::Joining);
                }

                actions
            }
            _ => Vec::new(),
        };

        let actions = actions
            .iter()
            .map(|action| format!("{} - {}", action.keybind(), action.label()))
            .collect::<Vec<String>>()
            .join("     ");

        Paragraph::new(actions).centered().bold().render(area, buf);
    }
}
