const std = @import("std");
const task = @import("./task.zig");

pub const BoardCreateError = error{
    RequireName,
};

pub const BoardTaskError = error{
    InvalidPosition,
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
            .tasks = std.ArrayList(task.Task).init(std.heap.HeapAllocator),
        };
    }

    pub fn deinit(self: *Board) void {
        self.tasks.deinit();
    }

    pub fn insertTask(self: *Board, position: usize, t: task.Task) !void {
        if ((self.tasks.items.len == 0) or (position >= self.tasks.items.len)) {
            return self.tasks.append(t);
        }
        return self.tasks.insert(position, t);
    }

    pub fn getTask(self: *Board, position: usize) BoardTaskError!task.Task {
        if (position >= self.tasks.items.len) {
            return error.InvalidPosition;
        }
        return self.tasks.items[position];
    }

    pub fn deleteTask(self: *Board, position: usize) void {
        if (position >= self.tasks.items.len) {
            return;
        }
        _ = self.tasks.swapRemove(position);
    }

    pub fn updateTask(self: *Board, position: usize, t: task.Task) !void {
        if (position >= self.tasks.items.len) {
            return error.InvalidPosition;
        }
        _ = self.tasks.swapRemove(position);
        try self.tasks.insert(position, t);
    }
};

test "create board - check for board name" {
    try std.testing.expectError(error.RequireName, Board.create(""));
}

test "create board - success" {
    var b = try Board.create("A name");
    defer b.deinit();
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

test "insert task - insert two tasks to the end" {
    var b = try Board.create("name");
    try std.testing.expect(0 == b.tasks.items.len);

    const t0 = try task.Task.create("task 0");
    try b.insertTask(10, t0);

    try std.testing.expect(1 == b.tasks.items.len);
    try std.testing.expectEqualSlices(u8, "task 0", b.tasks.items[0].title);

    const t1 = try task.Task.create("task 1");
    try b.insertTask(10, t1);

    try std.testing.expect(2 == b.tasks.items.len);
    try std.testing.expectEqualSlices(u8, "task 0", b.tasks.items[0].title);
    try std.testing.expectEqualSlices(u8, "task 1", b.tasks.items[1].title);
}

test "insert task - insert in the middle" {
    var b = try Board.create("name");
    try std.testing.expect(0 == b.tasks.items.len);

    const t0 = try task.Task.create("task 0");
    try b.insertTask(10, t0);

    try std.testing.expect(1 == b.tasks.items.len);
    try std.testing.expectEqualSlices(u8, "task 0", b.tasks.items[0].title);

    const t1 = try task.Task.create("task 1");
    try b.insertTask(10, t1);

    try std.testing.expect(2 == b.tasks.items.len);
    try std.testing.expectEqualSlices(u8, "task 0", b.tasks.items[0].title);
    try std.testing.expectEqualSlices(u8, "task 1", b.tasks.items[1].title);

    const t2 = try task.Task.create("task 2");
    try b.insertTask(1, t2);

    try std.testing.expect(3 == b.tasks.items.len);
    try std.testing.expectEqualSlices(u8, "task 0", b.tasks.items[0].title);
    try std.testing.expectEqualSlices(u8, "task 2", b.tasks.items[1].title);
    try std.testing.expectEqualSlices(u8, "task 1", b.tasks.items[2].title);
}

test "get task - get the first task" {
    var b = try Board.create("name");
    try std.testing.expect(0 == b.tasks.items.len);

    const expected = try task.Task.create("task 0");
    try b.insertTask(10, expected);

    const actual = try b.getTask(0);
    try std.testing.expectEqual(expected, actual);
}

test "get task - invalid position" {
    var b = try Board.create("name");
    try std.testing.expect(0 == b.tasks.items.len);

    try std.testing.expectError(error.InvalidPosition, b.getTask(2));
}

test "delete task - success" {
    var b = try Board.create("name");
    try std.testing.expect(0 == b.tasks.items.len);

    const expected = try task.Task.create("task 0");
    try b.insertTask(10, expected);

    b.deleteTask(0);

    try std.testing.expect(0 == b.tasks.items.len);
}

test "delete task - invalid position" {
    var b = try Board.create("name");
    try std.testing.expect(0 == b.tasks.items.len);

    b.deleteTask(0);

    try std.testing.expect(0 == b.tasks.items.len);
}

test "update task - success" {
    var b = try Board.create("name");
    try std.testing.expect(0 == b.tasks.items.len);

    const t1 = try task.Task.create("task 0");
    try b.insertTask(0, t1);

    const updated = try task.Task.create("task updated");
    try b.updateTask(0, updated);
}

test "update task - in the middle" {
    var b = try Board.create("name");
    try std.testing.expect(0 == b.tasks.items.len);

    const t1 = try task.Task.create("task 0");
    try b.insertTask(0, t1);

    const t2 = try task.Task.create("task 1");
    try b.insertTask(1, t2);

    const t3 = try task.Task.create("task 2");
    try b.insertTask(2, t3);

    const updated = try task.Task.create("task updated");
    try b.updateTask(1, updated);

    try std.testing.expectEqual(updated, b.tasks.items[1]);
}

test "update task - invalid position" {
    var b = try Board.create("name");
    try std.testing.expect(0 == b.tasks.items.len);

    const updated = try task.Task.create("task updated");
    try std.testing.expectError(error.InvalidPosition, b.updateTask(1, updated));
}
