use crate::task::Task;
use std::time::Instant;

#[derive(Debug)]
struct Board {
    pub name: String,
    pub created_at: Instant,
    pub updated_at: Instant,
    pub tasks: Vec<Task>,
}

impl Board {
    pub fn new(name: String) -> Result<Self, String> {
        if name.is_empty() {
            return Err("name must not be empty".to_string());
        }

        Ok(Self {
            name,
            created_at: Instant::now(),
            updated_at: Instant::now(),
            tasks: vec![],
        })
    }

    pub fn insert_task(&mut self, position: usize, task: &Task) {
        if position >= self.tasks.len() {
            return self.tasks.push(task.to_owned());
        }
        self.tasks.insert(position, task.to_owned())
    }

    pub fn get_task(&mut self, position: usize) -> Option<&Task> {
        self.tasks.get(position)
    }

    pub fn update_task(&mut self, position: usize, task: &Task) -> Result<(), String> {
        if position >= self.tasks.len() {
            return Err("invalid board position".to_owned());
        }

        self.updated_at = Instant::now();
        self.tasks[position] = task.to_owned();

        Ok(())
    }

    pub fn remove_task(&mut self, position: usize) {
        if position >= self.tasks.len() {
            return;
        }
        self.tasks.remove(position);
    }
}

#[cfg(test)]
mod test {
    use crate::task::Task;

    #[test]
    fn new_success() {
        let board = super::Board::new("test".to_string());

        assert!(board.is_ok());
        if let Ok(t) = board {
            assert_eq!("test".to_string(), t.name);
        }
    }

    #[test]
    fn new_no_name() {
        let board = super::Board::new("".to_string());

        assert!(board.is_err());
        if let Err(e) = board {
            assert_eq!("name must not be empty".to_string(), e);
        }
    }

    #[test]
    fn insert_task() {
        let board = super::Board::new("test board".to_string());
        assert!(board.is_ok());

        let mut b = board.unwrap();
        let t = Task::new("task 1".to_string()).unwrap();

        b.insert_task(0, &t);

        assert_eq!(b.tasks.len(), 1);
        assert_eq!(b.tasks[0].title, "task 1");
    }

    #[test]
    fn insert_task_to_the_top() {
        let board = super::Board::new("test board".to_string());
        assert!(board.is_ok());

        let mut b = board.unwrap();

        let t = Task::new("task 1".to_string()).unwrap();
        b.insert_task(0, &t);

        let t = Task::new("task 2".to_string()).unwrap();
        b.insert_task(0, &t);

        assert_eq!(b.tasks.len(), 2);
        assert_eq!(b.tasks[0].title, "task 2");
        assert_eq!(b.tasks[1].title, "task 1");
    }

    #[test]
    fn insert_task_to_the_bottom() {
        let board = super::Board::new("test board".to_string());
        assert!(board.is_ok());

        let mut b = board.unwrap();

        let t = Task::new("task 1".to_string()).unwrap();
        b.insert_task(0, &t);

        let t = Task::new("task 2".to_string()).unwrap();
        b.insert_task(1, &t);

        assert_eq!(b.tasks.len(), 2);
        assert_eq!(b.tasks[0].title, "task 1");
        assert_eq!(b.tasks[1].title, "task 2");
    }

    #[test]
    fn insert_task_to_the_middle() {
        let board = super::Board::new("test board".to_string());
        assert!(board.is_ok());

        let mut b = board.unwrap();

        let t = Task::new("task 1".to_string()).unwrap();
        b.insert_task(0, &t);

        let t = Task::new("task 2".to_string()).unwrap();
        b.insert_task(1, &t);

        let t = Task::new("task 3".to_string()).unwrap();
        b.insert_task(1, &t);

        assert_eq!(b.tasks.len(), 3);
        assert_eq!(b.tasks[0].title, "task 1");
        assert_eq!(b.tasks[1].title, "task 3");
        assert_eq!(b.tasks[2].title, "task 2");
    }

    #[test]
    fn insert_task_to_the_bottom_with_a_high_postition() {
        let board = super::Board::new("test board".to_string());
        assert!(board.is_ok());

        let mut b = board.unwrap();

        let t = Task::new("task 1".to_string()).unwrap();
        b.insert_task(0, &t);

        let t = Task::new("task 2".to_string()).unwrap();
        b.insert_task(10, &t);

        assert_eq!(b.tasks.len(), 2);
        assert_eq!(b.tasks[0].title, "task 1");
        assert_eq!(b.tasks[1].title, "task 2");
    }

    #[test]
    fn get_task() {
        let board = super::Board::new("test board".to_string());
        assert!(board.is_ok());

        let mut b = board.unwrap();

        let t = Task::new("task 1".to_string()).unwrap();
        b.insert_task(0, &t);

        let opt = b.get_task(0);

        assert!(opt.is_some());
        assert_eq!(t.title, opt.unwrap().title);
    }

    #[test]
    fn get_task_invalid_position() {
        let board = super::Board::new("test board".to_string());
        assert!(board.is_ok());

        let mut b = board.unwrap();

        let t = Task::new("task 1".to_string()).unwrap();
        b.insert_task(0, &t);

        let opt = b.get_task(1);

        assert!(opt.is_none());
    }

    #[test]
    fn update_task() {
        let board = super::Board::new("test board".to_string());
        assert!(board.is_ok());

        let mut b = board.unwrap();

        let t = Task::new("task 1".to_string()).unwrap();
        b.insert_task(0, &t);

        let t = Task::new("task updated".to_string()).unwrap();
        let result = b.update_task(0, &t);

        assert!(result.is_ok());
        let updated = b.get_task(0).unwrap();
        assert_eq!(t.title, updated.title);
    }

    #[test]
    fn update_task_invalid_position() {
        let board = super::Board::new("test board".to_string());
        assert!(board.is_ok());

        let mut b = board.unwrap();

        let t = Task::new("task 1".to_string()).unwrap();
        b.insert_task(0, &t);

        let t = Task::new("task updated".to_string()).unwrap();
        let result = b.update_task(12, &t);

        assert!(result.is_err());
        assert_eq!("invalid board position", result.err().unwrap())
    }

    #[test]
    fn remove_task() {
        let board = super::Board::new("test board".to_string());
        assert!(board.is_ok());

        let mut b = board.unwrap();

        let t = Task::new("task 1".to_string()).unwrap();
        b.insert_task(0, &t);
        b.remove_task(0);

        assert!(b.get_task(0).is_none());
    }

    #[test]
    fn remove_task_invalid_position() {
        let board = super::Board::new("test board".to_string());
        assert!(board.is_ok());

        let mut b = board.unwrap();

        let t = Task::new("task 1".to_string()).unwrap();
        b.insert_task(0, &t);
        b.remove_task(10);

        assert!(b.get_task(0).is_some());
    }
}
