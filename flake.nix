{
  description = "Polyglot todo app";

  inputs.nixpkgs.url = "nixpkgs/nixos-unstable";
  inputs.flake-utils.url = "github:numtide/flake-utils";
  inputs.rust-overlay.url = "github:oxalica/rust-overlay";
  inputs.zig.url = "github:mitchellh/zig-overlay";

  outputs =
    {
      self,
      nixpkgs,
      flake-utils,
      rust-overlay,
      zig,
    }:
    # Add dependencies that are only needed for development
    flake-utils.lib.eachDefaultSystem (
      system:
      let
        overlays = [
          (import rust-overlay)
          zig.overlays.default
        ];
        pkgs = import nixpkgs {
          inherit system overlays;
        };
      in
      {
        devShells.default =
          let
            p = pkgs;
          in
          pkgs.mkShell {
            buildInputs = [
              p.act
              p.go
              p.gopls
              p.rust-bin.stable.latest.default
              p.rust-analyzer
              p.zig
              p.zls
              p.elixir
              p.elixir-ls
              p.hare
            ];
          };
      }
    );
}
