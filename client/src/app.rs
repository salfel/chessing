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
    layout::{Constraint, Direction, Layout},
    prelude::CrosstermBackend,
    widgets::StatefulWidget,
    Frame, Terminal,
};
use tokio_tungstenite::{connect_async, tungstenite::protocol::Message};

use crate::{
    board::Board,
    state::{Color, State},
    stats::Stats,
};

type SocketWriter = futures_util::stream::SplitSink<
    tokio_tungstenite::WebSocketStream<tokio_tungstenite::MaybeTlsStream<tokio::net::TcpStream>>,
    Message,
>;
type SocketReader = futures_util::stream::Fuse<
    futures_util::stream::SplitStream<
        tokio_tungstenite::WebSocketStream<
            tokio_tungstenite::MaybeTlsStream<tokio::net::TcpStream>,
        >,
    >,
>;

#[allow(dead_code)]
struct SocketStream {
    writer: SocketWriter,
    reader: SocketReader,
}

pub struct App {
    should_quit: bool,
    state: Arc<Mutex<State>>,
    stream: SocketStream,
}

impl App {
    const FPS: f32 = 60.0;

    pub async fn new() -> Self {
        let stream = Self::init_socket().await;

        App {
            should_quit: false,
            state: Arc::new(Mutex::new(State::new())),
            stream,
        }
    }

    pub async fn run(
        mut self,
        terminal: &mut Terminal<CrosstermBackend<Stdout>>,
    ) -> io::Result<()> {
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
                Some(Ok(message)) = self.stream.reader.next() => {
                    self.on_message(message.to_string()).await;
                }
            }
        }

        Ok(())
    }

    async fn init_socket() -> SocketStream {
        let url = "ws://localhost:8000/game";
        let (ws_stream, _) = connect_async(url).await.expect("Failed to connect");
        let (mut write, read) = ws_stream.split();

        write
            .send(Message::text("create game:  test"))
            .await
            .expect("Failed to send message");

        SocketStream {
            writer: write,
            reader: read.fuse(),
        }
    }

    fn draw(&self, frame: &mut Frame, state: &mut State) {
        let layout = Layout::default()
            .direction(Direction::Horizontal)
            .constraints(vec![Constraint::Length(30), Constraint::Min(1)])
            .split(frame.area());

        Board::default().render(layout[0], frame.buffer_mut(), state);
        Stats::default().render(layout[1], frame.buffer_mut(), state);
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

        let splitted: Vec<String> = message.splitn(2, ':').map(String::from).collect();
        let command = splitted.get(0).unwrap();
        let details = splitted.get(1).map_or("", |value| value).trim();

        match command.as_str() {
            "state" => state.parse_pieces(details),
            "color" => state.color = Some(Color::new(details)),
            "code" => state.code = details.to_string(),
            _ => println!("{}", message),
        }
    }
}
