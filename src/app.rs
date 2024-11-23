use std::{
    io::{self, Stdout},
    time::Duration,
};

use ratatui::{
    crossterm::event::{self, Event, KeyCode, KeyEvent, KeyEventKind},
    prelude::{Buffer as TBuffer, CrosstermBackend, Rect},
    widgets::Widget,
    Frame, Terminal,
};

use crate::board::Board;

pub struct App {
    exit: bool,
}

impl App {
    pub fn new() -> Self {
        App { exit: false }
    }

    pub fn run(&mut self, terminal: &mut Terminal<CrosstermBackend<Stdout>>) -> io::Result<()> {
        while !self.exit {
            terminal.draw(|frame| self.render_frame(frame))?;
            self.handle_events()?;
        }

        Ok(())
    }

    fn render_frame(&mut self, frame: &mut Frame) {
        self.render(frame.area(), frame.buffer_mut());
    }

    fn handle_events(&mut self) -> io::Result<()> {
        if event::poll(Duration::from_millis(10))? {
            match event::read()? {
                Event::Key(event) if event.kind == KeyEventKind::Press => self.handle_keys(event),
                _ => {}
            }
        }

        Ok(())
    }

    fn handle_keys(&mut self, event: KeyEvent) {
        match event.code {
            KeyCode::Char('q') => self.exit = true,
            _ => {}
        }
    }
}

impl Widget for &App {
    fn render(self, area: Rect, buf: &mut TBuffer) {
        Board::new().render(area, buf);
    }
}
