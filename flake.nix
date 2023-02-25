{
  description = "This is a program representing rational numbers in Python";
  inputs.nixpkgs.url = "github:NixOS/nixpkgs/nixos-unstable";
  outputs = { self, nixpkgs }:
    let
      system = "x86_64-linux";
      pkgs = import nixpkgs {
        inherit system;
      };
    in
    {
      devShell.${system} =
        pkgs.mkShell
          {
            # The packages we need for this project
            buildInputs = with pkgs;
              [
                go
                go-tools
                gopls
              ];
          };
    };
}
