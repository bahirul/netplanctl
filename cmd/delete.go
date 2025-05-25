package cmd

import (
	"fmt"
	"slices"

	"github.com/bahirul/netplanctl/internal/netplan"
	"github.com/spf13/cobra"
)

var deleteCmd = &cobra.Command{
	Use:   "delete",
	Short: "delete a network interface configuration",
}

var deleteVlanCmd = &cobra.Command{
	Use:   "vlan",
	Short: "delete a VLAN interface configuration",
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) < 1 {
			fmt.Println("% command incomplete: delete vlan <interface>")
			return
		}

		ifaceName := args[0]
		config, err := netplan.LoadTemporaryOrRunConfig(temporaryNetplanFile, netplanFile)
		if err != nil {
			fmt.Println("% error:", err)
			return
		}

		if _, ok := config.Network.Vlans[ifaceName]; !ok {
			fmt.Println("% interface", ifaceName, "not found in the configuration.")
			return
		}

		delete(config.Network.Vlans, ifaceName)

		if err := netplan.SaveConfig(config, temporaryNetplanFile); err != nil {
			fmt.Println("% error saving configuration:", err)
			return
		}
	},
}

var deleteIpCmd = &cobra.Command{
	Use:   "ip",
	Short: "delete an IP address from an interface",
}

var deleteEthernetIpCmd = &cobra.Command{
	Use:   "ethernet",
	Short: "delete an IP address from an Ethernet interface",
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) < 2 {
			fmt.Println("% command incomplete: delete ethernet <interface> <ip>")
			return
		}

		ifaceName := args[0]
		ip := args[1]

		config, err := netplan.LoadTemporaryOrRunConfig(temporaryNetplanFile, netplanFile)
		if err != nil {
			fmt.Println("% error:", err)
			return
		}

		ethernet, ok := config.Network.Ethernets[ifaceName]
		if !ok {
			fmt.Println("% interface", ifaceName, "not found in the configuration.")
			return
		}

		found := false
		index := -1
		for i, addr := range ethernet.Addresses {
			if addr == ip {
				found = true
				index = i
				break
			}
		}
		if !found {
			fmt.Println("% IP address", ip, "not found on interface", ifaceName)
			return
		}

		// Remove the IP address from the slice
		ethernet.Addresses = slices.Delete(ethernet.Addresses, index, index+1)
		config.Network.Ethernets[ifaceName] = ethernet

		if err := netplan.SaveConfig(config, temporaryNetplanFile); err != nil {
			fmt.Println("% error saving configuration:", err)
			return
		}
	},
}

var deleteVlanIpCmd = &cobra.Command{
	Use:   "vlan",
	Short: "delete an IP address from a VLAN interface",
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) < 2 {
			fmt.Println("% command incomplete: delete vlan <interface> <ip>")
			return
		}

		ifaceName := args[0]
		ip := args[1]

		config, err := netplan.LoadTemporaryOrRunConfig(temporaryNetplanFile, netplanFile)
		if err != nil {
			fmt.Println("% error:", err)
			return
		}

		vlan, ok := config.Network.Vlans[ifaceName]
		if !ok {
			fmt.Println("% interface", ifaceName, "not found in the configuration.")
			return
		}

		found := false
		index := -1
		for i, addr := range vlan.Addresses {
			if addr == ip {
				found = true
				index = i
				break
			}
		}
		if !found {
			fmt.Println("% IP address", ip, "not found on interface", ifaceName)
			return
		}

		// Remove the IP address from the slice
		vlan.Addresses = slices.Delete(vlan.Addresses, index, index+1)
		config.Network.Vlans[ifaceName] = vlan

		if err := netplan.SaveConfig(config, temporaryNetplanFile); err != nil {
			fmt.Println("% error saving configuration:", err)
			return
		}
	},
}

func init() {
	deleteCmd.AddCommand(deleteVlanCmd)
	deleteCmd.AddCommand(deleteIpCmd)
	deleteIpCmd.AddCommand(deleteVlanIpCmd)
	deleteIpCmd.AddCommand(deleteEthernetIpCmd)
}
