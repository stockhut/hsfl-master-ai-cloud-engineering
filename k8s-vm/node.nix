{ config, pkgs, ... }:
let
  name = "NixOS K8s node";
in
{
  imports =
    [
      ./common.nix
    ];

  environment.etc = {
    "nixos/configuration.nix".source = ./node.nix;
    "nixos/common.nix".source = ./common.nix;
  };

  networking.hostName = "node";

  virtualbox = {
    vmDerivationName = name;
    vmFileName = "${name}.ova";
    vmName = name;
  };

  services.k3s = {
    enable = true;
    role = "server";
    token = "1234";
    serverAddr = "https://192.168.178.34:6443";
  };
}
