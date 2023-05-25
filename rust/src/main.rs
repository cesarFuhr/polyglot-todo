use clap::Parser;

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

    println!("==> {:?}", args.trailing)
}
