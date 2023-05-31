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
