# netplanctl

A command line tool to manage netplan configuration files built with Go.
To feel more like router world, it uses a simple command line interface to apply network configurations.

## Installation

- Clone this repository.
- Update `.env` file with your desired configuration.
- Run `go build` to compile the binary.
- Move the binary to a directory in your PATH, e.g., `/usr/local/bin`.
- Make sure the binary is executable: `chmod +x /usr/local/bin/netplanctl`.

## Supported Features

- [x] Display interfaces: ethernets and vlans
- [x] Display uncommitted changes
- [x] Set ip address: ethernet and vlan
- [x] Restore and backup configuration
- [x] Apply configuration
- [x] Create vlan interface
- [x] Delete interface: vlan
- [x] Delete ip address: ethernet and vlan
- [x] Shutdown interface: ethernet and vlan
- [x] Startup interface: ethernet and vlan
- [x] Frr console login alias for vtysh
