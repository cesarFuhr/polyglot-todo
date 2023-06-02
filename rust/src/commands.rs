use std::time::Instant;

use crate::board::Board;
use crate::task::Task;

pub fn list(b: &Board) {
    println!("  {}", b.name);
    b.tasks.iter().enumerate().for_each(|(id, task)| {
        print!("{} ", id + 1);
        match task.done {
            true => print!("☑ "),
            false => print!("☐ "),
        }
        println!("{}", task.title);
    })
}

pub fn add(b: &mut Board, title: String) -> Result<(), String> {
    let t = Task::new(title);
    match t {
        Ok(task) => {
            b.insert_task(0, &task);
            Ok(())
        }
        Err(e) => Err(format!("adding a new task: {}", e)),
    }
}

pub fn del(b: &mut Board, position: usize) -> Result<(), String> {
    let pos = position - 1;

    if let None = b.get_task(pos) {
        return Err("fetching task: invalid position".to_string());
    }

    b.remove_task(pos);

    Ok(())
}

pub fn done(b: &mut Board, position: usize) -> Result<(), String> {
    let pos = position - 1;

    let task = b.get_task(pos);
    let mut t = match task {
        None => return Err("fetching task: invalid position".to_string()),
        Some(t) => t.to_owned()
    };

    t.done = !t.done;
    t.done_at = Instant::now();

    b.update_task(pos, &t)
}

pub fn update(b: &mut Board, position: usize, title: String) -> Result<(), String> {
    let pos: usize = position - 1;

    let task = b.get_task(pos);
    let mut t = match task {
        None => return Err("fetching task: invalid position".to_string()),
        Some(t) => t.to_owned()
    };

    t.title = title;

    b.update_task(pos, &t)
}
