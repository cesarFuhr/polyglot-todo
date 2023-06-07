const std = @import("std");

pub const TaskCreateError = error{
    RequiredTitle,
};

pub const Task = struct {
    done: bool,
    title: []const u8,
    created_at: i64,
    done_at: i64,

    pub fn create(title: []const u8) TaskCreateError!Task {
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

test "success" {
    const t = try Task.create("A title");
    try std.testing.expectEqualSlices(u8, "A title", t.title);
}
