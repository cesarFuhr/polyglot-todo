use serde::{Deserialize, Serialize};
use std::time::Instant;

#[derive(Clone, Deserialize, Serialize, Debug)]
pub struct Task {
    pub done: bool,
    pub title: String,
    #[serde(with = "serde_millis")]
    pub created_at: Instant,
    #[serde(with = "serde_millis")]
    pub done_at: Instant,
}

impl Task {
    pub fn new(title: String) -> Result<Self, String> {
        if title.is_empty() {
            return Err("title must not be empty".to_string());
        }

        Ok(Self {
            done: false,
            title,
            created_at: Instant::now(),
            done_at: Instant::now(),
        })
    }
}

#[cfg(test)]
mod tests {
    #[test]
    fn it_works() {
        let task = super::Task::new("test".to_string());

        if let Ok(t) = task {
            assert_eq!("test".to_string(), t.title);
        }
    }

    #[test]
    fn no_title() {
        let task = super::Task::new("".to_string());

        if let Err(e) = task {
            assert_eq!("title must not be empty".to_string(), e);
        }
    }
}
