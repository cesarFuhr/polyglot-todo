use clap::Parser;
use std::{
    fs,
    io::{self, BufReader},
};

mod board;
mod task;

/// Simple todo app
#[derive(Parser, Debug)]
#[command(about, long_about = None, trailing_var_arg=true)]
struct Args {
    /// List tasks
    #[arg(short, group = "mutex", default_value_t = false)]
    list: bool,
    /// Adds new task
    #[arg(short, group = "mutex", group = "needs_input", default_value_t = false)]
    adds: bool,
    /// Marks a task as done
    #[arg(short, group = "mutex", default_value_t = 0)]
    done: usize,
    /// Deletes a task from the board
    #[arg(short = 'D', group = "mutex", default_value_t = 0)]
    del: usize,
    /// Update a given task
    #[arg(short, group = "mutex", group = "needs_input", default_value_t = 0)]
    update: usize,

    // Trailing arguments
    #[arg(requires = "needs_input", allow_hyphen_values = true)]
    trailing: Vec<String>,
}

fn main() {
    let args = Args::parse();

    let board = load_board(".todo.json".to_string()).unwrap();

    println!("Board {:?}", board);
    println!("Ok");

    println!("{:?}", save_board(".todo.json".to_string(), board));
}

fn load_board(path: String) -> Result<board::Board, String> {
    let result = fs::OpenOptions::new()
        .read(true)
        .write(true)
        .create(true)
        .open(path);
    let file = match result {
        Ok(f) => f,
        Err(e) => return Err(e.to_string()),
    };
    let buf = BufReader::new(file);

    let result = serde_json::from_reader(buf);
    match result {
        Ok(v) => v,
        Err(e) => {
            if e.is_eof() {
                return board::Board::new("TODO".to_string());
            }

            Err(e.to_string())
        }
    }
}

fn save_board(path: String, b: board::Board) -> Result<(), String> {
    let result = fs::OpenOptions::new().write(true).truncate(true).open(path);
    let file = match result {
        Ok(f) => f,
        Err(e) => return Err(e.to_string()),
    };

    let result = serde_json::to_writer_pretty(file, &b);
    match result {
        Ok(()) => Ok(()),
        Err(e) => Err(e.to_string()),
    }
}
