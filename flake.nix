{
  description =
    "A flake for my Go program to add pinyin tone marks to plain text files";
  inputs = {
    nixpkgs.url = "github:NixOS/nixpkgs/nixpkgs-unstable";
    flake-utils.url = "github:numtide/flake-utils";
  };

  outputs = { self, nixpkgs, flake-utils }:
    flake-utils.lib.eachDefaultSystem (system:
      let pkgs = nixpkgs.legacyPackages.${system};
      in {
        devShell = pkgs.mkShell {
          # The packages we need for this project
          buildInputs = with pkgs; [ go_1_22 go-tools gopls golangci-lint ];
        };
      });
}
