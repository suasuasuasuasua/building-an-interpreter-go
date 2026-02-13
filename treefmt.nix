# treefmt.nix
{
  # Used to find the project root
  projectRootFile = "flake.nix";
  settings = {
    on-unmatched = "debug";
  };

  programs = {
    gofmt.enable = true;
    golangci-lint.enable = false;
    mdformat.enable = true;
    nixfmt.enable = true;
  };
}
