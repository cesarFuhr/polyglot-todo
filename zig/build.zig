const std = @import("std");

pub fn build(b: *std.Build) !void {
    const target = b.standardTargetOptions(.{});
    const optimize = b.standardOptimizeOption(.{});

    const exe = b.addExecutable(.{
        .name = "todo",
        .root_source_file = .{ .path = "src/main.zig" },
        .target = target,
        .optimize = optimize,
    });

    const clap = b.dependency("clap", .{
        .target = target,
        .optimize = optimize,
    });
    // clap has exported itself as clap
    // now you are re-exporting clep
    // as a module in your project with the name clap
    exe.addModule("clap", clap.module("clap"));

    b.installArtifact(exe);

    const test_step = b.step("test", "Run unit tests");
    const unit_tests = b.addTest(.{
        .root_source_file = .{ .path = "src/tests.zig" },
        .target = target,
    });

    const run_unit_tests = b.addRunArtifact(unit_tests);
    test_step.dependOn(&run_unit_tests.step);

    //const run_exe = b.addRunArtifact(exe);
    //const run_step = b.step("run", "Run the application");
    //run_step.dependOn(&run_exe.step);
}
