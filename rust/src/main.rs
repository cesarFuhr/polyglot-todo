use clap::Parser;
use std::{fs, io::BufReader};

mod board;
mod commands;
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
    add: bool,
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

    let mut board = load_board(".todo.json".to_string()).unwrap();

    if args.list {
        commands::list(&board);
        return;
    }

    let mut result: Result<(), String> = Err("no command was selected".to_string());
    if args.add {
        result = commands::add(&mut board, args.trailing.join(" "));
    }

    if args.del > 0 {
        result = commands::del(&mut board, args.del);
    }

    if args.done > 0 {
        result = commands::done(&mut board, args.done);
    }

    if args.update > 0 {
        result = commands::update(&mut board, args.update, args.trailing.join(" "));
    }

    match result {
        Ok(()) => save_board(".todo.json".to_string(), board).unwrap(),
        Err(e) => println!("{}", e),
    }
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

    let v = serde_json::from_reader(buf);
    match v {
        Ok(v) => Ok(v),
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
