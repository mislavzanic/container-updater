{
  description = "A simple Go package";

  inputs.nixpkgs.url = "github:NixOS/nixpkgs/nixos-unstable";
  inputs.flake-utils.url = "github:numtide/flake-utils";

  outputs = { self, nixpkgs, flake-utils }:
    let

      lastModifiedDate = self.lastModifiedDate or self.lastModified or "19700101";
      version = builtins.substring 0 8 lastModifiedDate;
      supportedSystems = [ "x86_64-linux" "x86_64-darwin" "aarch64-linux" "aarch64-darwin" ];
      forAllSystems = nixpkgs.lib.genAttrs supportedSystems;
      nixpkgsFor = forAllSystems (system: import nixpkgs { inherit system; });

    in
    {
      packages = forAllSystems (system:
        let
          pkgs = nixpkgsFor.${system};
          src = ./.;
        in rec {
          default = pkgs.buildGoModule {
            pname = "container-updater";
            inherit version;
            src = ./.;
            vendorSha256 = "sha256-Pu78LI3fR0akbIyTzE3cWUD7FhpyEhcts1/qq8B+91M=";
          };

          docker = pkgs.dockerTools.buildLayeredImage {
            name = "mislavzanic/container-updater";
            tag = "dev";

            contents = [ default ];

            config = {
              Cmd = [ "${default}/bin/container-updater" ];
              WorkingDir = "${default}";
            };
          };

        });


      apps = forAllSystems (system:
        let
          pkgs = nixpkgsFor.${system};
          upload-script = pkgs.writeShellScriptBin "upload-image" ''
            set -eu
            nix build .#docker
            docker load < result
            docker push $USERNAME/container-updater:dev
          '';
        in {
          upload-script = flake-utils.lib.mkApp { drv = upload-script; };
        }
      );

      defaultPackage = forAllSystems (system: self.packages.${system}.gotsm);
    };
}
