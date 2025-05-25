package cmd

import (
	"fmt"
	"os"
	"sort"
	"strings"

	"github.com/bahirul/netplanctl/internal/netplan"
	"github.com/bahirul/netplanctl/internal/status"
	"github.com/sergi/go-diff/diffmatchpatch"
	"github.com/spf13/cobra"
)

var showCmd = &cobra.Command{
	Use:   "show",
	Short: "show the current network configuration",
}

var showEthernetCmd = &cobra.Command{
	Use:   "ethernets",
	Short: "show ethernets configuration",
	Run: func(cmd *cobra.Command, args []string) {
		config, err := netplan.LoadConfig(netplanFile)

		// sort the keys of the Ethernets map
		keys := make([]string, 0, len(config.Network.Ethernets))
		for k := range config.Network.Ethernets {
			keys = append(keys, k)
		}
		sort.Strings(keys)

		if err != nil {
			fmt.Println("% error:", err)
			os.Exit(1)
		}

		fmt.Printf("%-15s %-20s %-20s %-15s %-5s\n", "Interface", "IP Address", "Mac Address", "Admin/Status", "MTU")
		fmt.Println(strings.Repeat("-", 79))

		// no ethernet interfaces
		if len(config.Network.Ethernets) == 0 {
			fmt.Println("% no ethernet interfaces found.")
			return
		}

		for _, key := range keys {
			iface := config.Network.Ethernets[key]
			ipAddr := ""
			stateLink := ""
			mtu := iface.MTU

			// find mtu from config file first
			if mtu == "" {
				mtu, err = status.GetInterfaceMTU(key)
				if err != nil {
					mtu = "N/A"
				}
			}

			operState, carrier, err := status.GetInterfaceStatus(key)
			if err != nil {
				stateLink = "N/A"
			} else {
				stateLink = fmt.Sprintf("%s/%s", operState, carrier)
			}

			if len(iface.Addresses) > 0 {
				ipAddr = iface.Addresses[0]
			}

			if len(iface.Addresses) > 1 {
				fmt.Printf("%-15s %-20s %-20s %-15s %-5s\n", key, ipAddr, iface.Match.MacAddress, stateLink, mtu)

				for _, ip := range iface.Addresses[1:] {
					fmt.Printf("%-15s %-20s %-20s %-15s %-5s\n", "", ip, "", "", "")
				}
			} else {
				fmt.Printf("%-15s %-20s %-20s %-15s %-5s\n", key, ipAddr, iface.Match.MacAddress, stateLink, mtu)
			}
		}
	},
}

var showVlanCmd = &cobra.Command{
	Use:   "vlans",
	Short: "show the current vlans configuration",
	Run: func(cmd *cobra.Command, args []string) {
		config, err := netplan.LoadConfig(netplanFile)

		if err != nil {
			fmt.Println("% error:", err)
			os.Exit(1)
		}

		fmt.Printf("%-15s %-20s %-20s %-20s\n", "Name", "Port", "VLAN ID", "IP Address")
		fmt.Println(strings.Repeat("-", 78))

		if len(config.Network.Vlans) == 0 {
			fmt.Println("% no vlans found.")
			return
		}

		for name, vlan := range config.Network.Vlans {
			ipAddr := ""
			if len(vlan.Addresses) > 0 {
				ipAddr = vlan.Addresses[0]
			}
			fmt.Printf("%-15s %-20s %-20d %-20s\n", name, vlan.Link, vlan.Id, ipAddr)

			if len(vlan.Addresses) > 1 {
				for _, ip := range vlan.Addresses[1:] {
					fmt.Printf("%-15s %-20s %-20s %-20s\n", "", "", "", ip)
				}
			}
		}
	},
}

var showUncommittedCmd = &cobra.Command{
	Use:   "uncommitted",
	Short: "show uncommitted changes",
	Run: func(cmd *cobra.Command, args []string) {
		tempConfig, err := netplan.RawConfig(temporaryNetplanFile)
		if err != nil {
			fmt.Println("% error:", err)
			return
		}

		currentConfig, err := netplan.RawConfig(netplanFile)
		if err != nil {
			fmt.Println("% error:", err)
			return
		}

		dmp := diffmatchpatch.New()
		diff := dmp.DiffMain(currentConfig, tempConfig, false)

		if currentConfig == tempConfig || len(diff) == 0 {
			fmt.Println("% no uncommitted changes found.")
			return
		}

		fmt.Println(dmp.DiffPrettyText(diff))
	},
}

var showVersionCmd = &cobra.Command{
	Use:   "version",
	Short: "show the version of netplanctl",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println(version)
	},
}

func init() {
	showCmd.AddCommand(showEthernetCmd)
	showCmd.AddCommand(showVlanCmd)
	showCmd.AddCommand(showUncommittedCmd)
	showCmd.AddCommand(showVersionCmd)
}
