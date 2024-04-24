pub const tasks = @import("task.zig");
pub const board = @import("board.zig");

test {
    @import("std").testing.refAllDecls(@This());
}
