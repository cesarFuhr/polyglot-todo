const std = @import("std");
const testing = std.testing;
const task = @import("./task.zig");

pub const BoardCreateError = error{
    RequireName,
    OutOfMemory,
};

pub const BoardTaskError = error{
    InvalidPosition,
};

const serializableBoard = struct {
    name: []const u8,
    created_at: i64,
    updated_at: i64,
    tasks: []task.Task,
};

pub const Board = struct {
    name: []const u8,
    created_at: i64,
    updated_at: i64,
    tasks: std.ArrayList(task.Task),
    allocator: std.mem.Allocator,

    pub fn create(name: []const u8, allocator: std.mem.Allocator) BoardCreateError!Board {
        if (name.len == 0) {
            return error.RequireName;
        }

        const tasks = std.ArrayList(task.Task).init(allocator);
        const dupedName = try allocator.dupe(u8, name);

        return .{
            .name = dupedName,
            .created_at = std.time.timestamp(),
            .updated_at = std.time.timestamp(),
            .tasks = tasks,
            .allocator = allocator,
        };
    }

    pub fn decode(reader: anytype, allocator: std.mem.Allocator) anyerror!Board {
        // Convert it to a json reader.
        var jsonReader = std.json.reader(allocator, reader);
        defer jsonReader.deinit();

        const parsed = try std.json.parseFromTokenSource(serializableBoard, allocator, &jsonReader, .{
            .ignore_unknown_fields = true,
            .allocate = std.json.AllocWhen.alloc_always,
        });
        defer parsed.deinit();

        const board = parsed.value;
        const tasks = try allocator.dupe(task.Task, board.tasks);

        const dupedName = try allocator.dupe(u8, board.name);

        return .{
            .name = dupedName,
            .created_at = board.created_at,
            .updated_at = board.updated_at,
            .tasks = std.ArrayList(task.Task).fromOwnedSlice(allocator, tasks),
            .allocator = allocator,
        };
    }

    pub fn encode(self: *Board, writer: anytype) anyerror!void {
        var clonedTasks = try self.tasks.clone();
        defer clonedTasks.deinit();

        const taskSlice = try clonedTasks.toOwnedSlice();
        defer self.allocator.free(taskSlice);

        const serializable = serializableBoard{
            .name = self.name,
            .created_at = self.created_at,
            .updated_at = self.updated_at,
            .tasks = taskSlice,
        };

        return std.json.stringify(serializable, .{ .whitespace = .indent_2 }, writer);
    }

    pub fn deinit(self: *Board) void {
        self.tasks.deinit();
        self.allocator.free(self.name);
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
    try testing.expectError(error.RequireName, Board.create("", testing.allocator));
}

test "create board - success" {
    var b = try Board.create("A name", testing.allocator);
    defer b.deinit();
    try testing.expectEqualSlices(u8, "A name", b.name);
}

test "load board - success" {
    const s =
        \\    {
        \\  "name": "A name",
        \\  "created_at": 1735609291,
        \\  "updated_at": 1735609291,
        \\  "tasks": [
        \\    {
        \\      "done": false,
        \\      "title": "anything",
        \\      "created_at": 1735609291,
        \\      "done_at": 1735609291
        \\    }
        \\  ]
        \\}
    ;

    var byteStream = std.io.fixedBufferStream(s);

    var b = try Board.decode(byteStream.reader(), testing.allocator);
    defer b.deinit();

    try testing.expectEqualSlices(u8, "A name", b.name);
}

test "load board - fail missing field" {
    const s =
        \\{
        \\  "created_at": 1735609291,
        \\  "updated_at": 1735609291,
        \\  "tasks": [
        \\    {
        \\      "done": false,
        \\      "title": "anything",
        \\      "created_at": 1735609291,
        \\      "done_at": 1735609291
        \\    }
        \\  ]
        \\}
    ;

    var byteStream = std.io.fixedBufferStream(s);
    try testing.expectError(error.MissingField, Board.decode(byteStream.reader(), testing.allocator));
}

test "write board without tasks - success" {
    var b = try Board.create("A name", testing.allocator);
    defer b.deinit();
    b.created_at = 1735609291;
    b.updated_at = 1735609291;

    var list = std.ArrayList(u8).init(testing.allocator);
    try b.encode(list.writer());

    const expected =
        \\{
        \\  "name": "A name",
        \\  "created_at": 1735609291,
        \\  "updated_at": 1735609291,
        \\  "tasks": []
        \\}
    ;

    const actual = try list.toOwnedSlice();
    try testing.expectEqualStrings(expected, actual);
}

test "insert task - append to an empty list" {
    var b = try Board.create("name", testing.allocator);
    defer b.deinit();
    try testing.expect(0 == b.tasks.items.len);

    const t = try task.Task.create("task name");
    try b.insertTask(0, t);

    try testing.expect(1 == b.tasks.items.len);
    try testing.expectEqualSlices(u8, "task name", b.tasks.items[0].title);
}

test "insert task - insert two tasks to the end" {
    var b = try Board.create("name", testing.allocator);
    defer b.deinit();
    try testing.expect(0 == b.tasks.items.len);

    const t0 = try task.Task.create("task 0");
    try b.insertTask(10, t0);

    try testing.expect(1 == b.tasks.items.len);
    try testing.expectEqualSlices(u8, "task 0", b.tasks.items[0].title);

    const t1 = try task.Task.create("task 1");
    try b.insertTask(10, t1);

    try testing.expect(2 == b.tasks.items.len);
    try testing.expectEqualSlices(u8, "task 0", b.tasks.items[0].title);
    try testing.expectEqualSlices(u8, "task 1", b.tasks.items[1].title);
}

test "insert task - insert in the middle" {
    var b = try Board.create("name", testing.allocator);
    defer b.deinit();
    try testing.expect(0 == b.tasks.items.len);

    const t0 = try task.Task.create("task 0");
    try b.insertTask(10, t0);

    try testing.expect(1 == b.tasks.items.len);
    try testing.expectEqualSlices(u8, "task 0", b.tasks.items[0].title);

    const t1 = try task.Task.create("task 1");
    try b.insertTask(10, t1);

    try testing.expect(2 == b.tasks.items.len);
    try testing.expectEqualSlices(u8, "task 0", b.tasks.items[0].title);
    try testing.expectEqualSlices(u8, "task 1", b.tasks.items[1].title);

    const t2 = try task.Task.create("task 2");
    try b.insertTask(1, t2);

    try testing.expect(3 == b.tasks.items.len);
    try testing.expectEqualSlices(u8, "task 0", b.tasks.items[0].title);
    try testing.expectEqualSlices(u8, "task 2", b.tasks.items[1].title);
    try testing.expectEqualSlices(u8, "task 1", b.tasks.items[2].title);
}

test "get task - get the first task" {
    var b = try Board.create("name", testing.allocator);
    defer b.deinit();
    try testing.expect(0 == b.tasks.items.len);

    const expected = try task.Task.create("task 0");
    try b.insertTask(10, expected);

    const actual = try b.getTask(0);
    try testing.expectEqual(expected, actual);
}

test "get task - invalid position" {
    var b = try Board.create("name", testing.allocator);
    defer b.deinit();
    try testing.expect(0 == b.tasks.items.len);

    try testing.expectError(error.InvalidPosition, b.getTask(2));
}

test "delete task - success" {
    var b = try Board.create("name", testing.allocator);
    defer b.deinit();
    try testing.expect(0 == b.tasks.items.len);

    const expected = try task.Task.create("task 0");
    try b.insertTask(10, expected);

    b.deleteTask(0);

    try testing.expect(0 == b.tasks.items.len);
}

test "delete task - invalid position" {
    var b = try Board.create("name", testing.allocator);
    defer b.deinit();
    try testing.expect(0 == b.tasks.items.len);

    b.deleteTask(0);

    try testing.expect(0 == b.tasks.items.len);
}

test "update task - success" {
    var b = try Board.create("name", testing.allocator);
    defer b.deinit();
    try testing.expect(0 == b.tasks.items.len);

    const t1 = try task.Task.create("task 0");
    try b.insertTask(0, t1);

    const updated = try task.Task.create("task updated");
    try b.updateTask(0, updated);
}

test "update task - in the middle" {
    var b = try Board.create("name", testing.allocator);
    defer b.deinit();
    try testing.expect(0 == b.tasks.items.len);

    const t1 = try task.Task.create("task 0");
    try b.insertTask(0, t1);

    const t2 = try task.Task.create("task 1");
    try b.insertTask(1, t2);

    const t3 = try task.Task.create("task 2");
    try b.insertTask(2, t3);

    const updated = try task.Task.create("task updated");
    try b.updateTask(1, updated);

    try testing.expectEqual(updated, b.tasks.items[1]);
}

test "update task - invalid position" {
    var b = try Board.create("name", testing.allocator);
    defer b.deinit();
    try testing.expect(0 == b.tasks.items.len);

    const updated = try task.Task.create("task updated");
    try testing.expectError(error.InvalidPosition, b.updateTask(1, updated));
}
