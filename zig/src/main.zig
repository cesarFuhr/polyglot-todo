const std = @import("std");
const clap = @import("clap");

const task = @import("./task.zig");
const board = @import("./board.zig");

pub fn main() !void {
    const params = comptime clap.parseParamsComptime(
        \\-h, --help             Display this help and exit.
        \\-l, --list
        \\-a, --add <str>...     
        \\-d, --done <usize>
        \\-D, --del <usize>
        \\-u, --update <string>...
        \\
    );

    var gpa = std.heap.GeneralPurposeAllocator(.{}){};
    const allocator = gpa.allocator();
    defer {
        const deinit_status = gpa.deinit();
        if (deinit_status == .leak) @panic("we had a leak");
    }

    var diag = clap.Diagnostic{};
    var res = clap.parse(clap.Help, &params, clap.parsers.default, .{
        .diagnostic = &diag,
        .allocator = allocator,
    }) catch |err| {
        // Report useful error and exit
        diag.report(std.io.getStdErr().writer(), err) catch {};
        return err;
    };
    defer res.deinit();

    const fileName = "todo.json";
    // Try to open the file.
    const openFlags = std.fs.File.OpenFlags{
        .mode = .read_write,
        .lock = .exclusive,
    };
    var file = std.fs.cwd().openFile(fileName, openFlags) catch |err| switch (err) {
        // If the file wasn't found, create one.
        error.FileNotFound => try std.fs.cwd().createFile(fileName, .{ .read = true }),
        else => {
            return err;
        },
    };
    defer file.close();

    var b = try loadOrCreate(file, allocator);
    defer b.deinit();

    if (res.args.help != 0)
        return clap.usage(std.io.getStdErr().writer(), clap.Help, &params);
    for (res.args.add) |a| {
        const t = try task.Task.create(a);
        try b.insertTask(0, t);
    }

    const stdout = std.io.getStdOut().writer();

    _ = try stdout.print("Board: {s}\n", .{b.name});
    _ = try stdout.write("Tasks:\n");
    for (b.tasks.items) |t| {
        _ = try stdout.print("- {s}\n", .{t.title});
    }

    // Write into the file.
    try file.seekTo(0);
    try b.encode(file.writer());
}

fn loadOrCreate(file: std.fs.File, allocator: std.mem.Allocator) !board.Board {
    const fileMeta = try file.metadata();
    // If there is something in the file try to load.
    if (fileMeta.size() > 0) {
        var buffered = std.io.bufferedReader(file.reader());
        return board.Board.decode(buffered.reader(), allocator);
    }

    // Otherwise create a new one.
    return board.Board.create("todo", allocator);
}
