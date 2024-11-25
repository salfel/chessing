mod action;
mod app;
mod board;
mod input;
mod piece;
mod state;
mod stats;

use std::io;

use app::App;

#[tokio::main]
async fn main() -> io::Result<()> {
    let mut terminal = ratatui::init();
    let app_result = App::new().await.run(&mut terminal).await;
    ratatui::restore();
    app_result
}
