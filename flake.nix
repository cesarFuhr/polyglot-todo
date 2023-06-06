{
  description = "Go 1.20 workspace";

  inputs.nixpkgs.url = "nixpkgs/nixos-22.11";
  inputs.flake-utils.url = "github:numtide/flake-utils";
  inputs.rust-overlay.url = "github:oxalica/rust-overlay";

  outputs = { self, nixpkgs, flake-utils, rust-overlay }:
    # Add dependencies that are only needed for development
    flake-utils.lib.eachDefaultSystem
      (system:
        let
          overlays = [ (import rust-overlay) ];
          pkgs = import nixpkgs {
            inherit system overlays;
          };
        in
        {
          devShells.default = let p = pkgs; in
            pkgs.mkShell {
              buildInputs =
                [
                  p.act
                  p.go_1_20
                  p.gopls
                  p.gotools
                  p.go-tools
                  p.go-outline
                  p.gocode
                  p.gopkgs
                  p.gocode-gomod
                  p.godef
                  p.golint
                  p.go-mockery
                  p.rust-bin.stable.latest.default
                  p.rust-analyzer
                  p.zig
                  p.zls
                ];
            };
        });
}

