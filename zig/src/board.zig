const std = @import("std");
const task = @import("./task.zig");

pub const BoardCreateError = error{
    RequireName,
};

pub const Board = struct {
    name: []const u8,
    created_at: i64,
    updated_at: i64,
    tasks: std.ArrayList(task.Task),

    pub fn create(name: []const u8) BoardCreateError!Board {
        if (name.len == 0) {
            return error.RequireName;
        }

        return Board{
            .name = name,
            .created_at = std.time.timestamp(),
            .updated_at = std.time.timestamp(),
            .tasks = std.ArrayList(task.Task).init(std.heap.page_allocator),
        };
    }

    pub fn insertTask(self: *Board, position: usize, t: task.Task) !void {
        if ((self.tasks.items.len == 0) or (position >= self.tasks.items.len)) {
            return self.tasks.append(t);
        }
        return self.tasks.insert(position, t);
    }
};

test "create board - check for board name" {
    try std.testing.expectError(error.RequireName, Board.create(""));
}

test "create board - success" {
    const b = try Board.create("A name");
    try std.testing.expectEqualSlices(u8, "A name", b.name);
}

test "insert task - append to an empty list" {
    var b = try Board.create("name");
    try std.testing.expect(0 == b.tasks.items.len);

    const t = try task.Task.create("task name");
    try b.insertTask(0, t);

    try std.testing.expect(1 == b.tasks.items.len);
    try std.testing.expectEqualSlices(u8, "task name", b.tasks.items[0].title);
}
