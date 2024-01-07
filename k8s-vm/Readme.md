# K8s test VM

We use [NixOS](nixos.org) as basis for our virtual machines, which allows us to build reproducible VirtualBox applicances (or other types of images) from declarative configurations.

These images are configured to be as easy to use a possible, consequently they are totally unsafe and should not be exposed to unsafe networks.

The confguration includes:
- a user ("user") with `wheel` and empty password
- disabled firewall
- SSH server
- k3s
- git
- openssl
- Editors: vim and helix
- Zellij as terminal multiplexer

## Building and running Virtualbox OVA

Install [NixOS Generators](https://github.com/nix-community/nixos-generators)

Note: Since kvm is used during the build, make sure no other processes are using it (e.g. stop running VMs before building)



### Control Plane

The following command will build the image, save it to the Nix store and print the path (e.g. `/nix/store/10gvkacfikqrvz1ixfr6rzarh1vnna7m-NixOS-K8s-node/NixOS K8s node.ova`).

```console
nixos-generate -f virtualbox -c control-plane.nix
```

#### In VirtualBox 
- File -> Import Appliance -> Enter the image path (like `/nix/store/XXX/YYY.ova`)
- Set CPU and memory
- Confirm settings
- Wait until the OVA is imported
- Select the machine and go to Settings -> Network
- Select `Bridged Adapter`
- OK

- Start VM

Inside the vm, run 
```console
ip a s
```
to find the IP address

### Node

Insert the control plane vm ip address in the `serverAdr` in `node.nix`

Build the image
```console
nixos-generate -f virtualbox -c node.nix
```

Import the Appliance and start the machine.

If `serverAddr` is configured correctly and reachable, the node will join the cluster automatically.

## Troubleshooting

###  "too many learners" error when joining the cluster
https://github.com/k3s-io/k3s/issues/2306

reset k3s
https://github.com/k3s-io/k3s/discussions/7208

```console
sudo rm -rf /var/lib/rancher/k3s
sudo rm -rf /root/.kube
```
