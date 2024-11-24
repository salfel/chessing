use std::{
    io::{self, Stdout},
    sync::Arc,
    time::Duration,
};

use crossterm::event::{KeyCode, KeyEventKind};
use tokio::sync::Mutex;

use futures_util::{stream::StreamExt, SinkExt};
use ratatui::{
    crossterm::event::{Event, EventStream},
    prelude::CrosstermBackend,
    widgets::StatefulWidget,
    Frame, Terminal,
};
use tokio_tungstenite::{connect_async, tungstenite::protocol::Message};

use crate::{board::Board, state::State};

pub struct App {
    should_quit: bool,
    state: Arc<Mutex<State>>,
}

impl App {
    const FPS: f32 = 60.0;

    pub fn new() -> Self {
        App {
            should_quit: false,
            state: Arc::new(Mutex::new(State::new())),
        }
    }

    pub async fn run(
        mut self,
        terminal: &mut Terminal<CrosstermBackend<Stdout>>,
    ) -> io::Result<()> {
        let url = "ws://localhost:8000/game";
        let (ws_stream, _) = connect_async(url).await.expect("Failed to connect");
        let (mut write, read) = ws_stream.split();

        write
            .send(Message::text("create game: "))
            .await
            .expect("Failed to send message");

        let mut read = read.fuse();
        let period = Duration::from_secs_f32(1.0 / Self::FPS);
        let mut interval = tokio::time::interval(period);
        let mut events = EventStream::new();

        while !self.should_quit {
            tokio::select! {
                _ = interval.tick() => {
                    let mut state = self.state.lock().await;
                    terminal.draw(|frame| self.draw(frame, &mut state))?;
                },
                Some(Ok(event)) = events.next() => {self.handle_event(event); },
                Some(Ok(message)) = read.next() => {
                    self.on_message(message.to_string()).await;
                }
            }
        }

        Ok(())
    }

    fn draw(&self, frame: &mut Frame, state: &mut State) {
        Board::new().render(frame.area(), frame.buffer_mut(), state);
    }

    fn handle_event(&mut self, event: Event) {
        if let Event::Key(key) = event {
            if key.kind == KeyEventKind::Press {
                match key.code {
                    KeyCode::Char('q') => self.should_quit = true,
                    _ => {}
                }
            }
        }
    }

    async fn on_message(&mut self, message: String) {
        let mut state = self.state.lock().await;
        state.pieces.push(message);
    }
}
