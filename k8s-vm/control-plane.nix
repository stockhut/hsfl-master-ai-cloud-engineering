{ config, pkgs, ... }:
let
  name = "NixOS K8s controle plane";
in
{
  imports =
    [
      ./common.nix
    ];

  networking.hostName = "controlplane";

  virtualbox = {
    vmDerivationName = name;
    vmFileName = "${name}.ova";
    vmName = name;
    params = {
      cpus = 2;
    };
  };
  services.k3s = {
    enable = true;
    role = "server";
    token = "1234";
    clusterInit = true;
  };
}
