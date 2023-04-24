{
  description = "Go 1.20 workspace";

  inputs.nixpkgs.url = "nixpkgs/nixos-22.11";
  inputs.flake-utils.url = "github:numtide/flake-utils";

  outputs = { self, nixpkgs, flake-utils }:
    # Add dependencies that are only needed for development
    flake-utils.lib.eachDefaultSystem
      (system:
        let
          pkgs = nixpkgs.legacyPackages.${system};
        in
        {
          devShells.default = let p = pkgs; in
            pkgs.mkShell {
              buildInputs =
                [
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
                ];
            };
        });
}

