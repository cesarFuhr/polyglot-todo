const std = @import("std");

pub const TaskCreateError = error{
    RequiredTitle,
};

pub const Task = struct {
    done: bool,
    title: []u8,
    created_at: i64,
    done_at: i64,

    pub fn create(title: []u8) TaskCreateError!Task {
        if (title.len == 0) {
            return error.RequiredTitle;
        }

        return Task{
            .done = false,
            .title = title,
            .created_at = std.time.timestamp(),
            .done_at = std.time.timestamp(),
        };
    }
};

test "check for task title" {
    try std.testing.expectError(error.RequiredTitle, Task.create(""));
}
