package cmd

import (
	"fmt"

	"github.com/bahirul/netplanctl/internal/netplan"
	"github.com/spf13/cobra"
)

var startupCmd = &cobra.Command{
	Use:   "startup",
	Short: "startup the network interfaces",
}

var ethernetStartupCmd = &cobra.Command{
	Use:   "ethernet",
	Short: "startup an ethernet interface",
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) < 1 {
			fmt.Println("% command incomplete: startup ethernet <interface>")
			return
		}

		ifaceName := args[0]
		config, err := netplan.LoadTemporaryOrRunConfig(temporaryNetplanFile, netplanFile)
		if err != nil {
			fmt.Println("% error:", err)
			return
		}

		iface, ok := config.Network.Ethernets[ifaceName]
		if !ok {
			fmt.Println("% interface", ifaceName, "not found in the configuration.")
			return
		}

		// remove ActivationMode if it exists
		if iface.ActivationMode == "off" {
			iface.ActivationMode = ""
		}

		config.Network.Ethernets[ifaceName] = iface
		if err := netplan.SaveConfig(config, temporaryNetplanFile); err != nil {
			fmt.Println("% error saving configuration:", err)
			return
		}
	},
}

var vlanStartupCmd = &cobra.Command{
	Use:   "vlan",
	Short: "startup a VLAN interface",
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) < 1 {
			fmt.Println("% command incomplete: startup vlan <interface>")
			return
		}

		ifaceName := args[0]
		config, err := netplan.LoadTemporaryOrRunConfig(temporaryNetplanFile, netplanFile)
		if err != nil {
			fmt.Println("% error:", err)
			return
		}

		iface, ok := config.Network.Vlans[ifaceName]
		if !ok {
			fmt.Println("% interface", ifaceName, "not found in the configuration.")
			return
		}

		// remove ActivationMode if it exists
		if iface.ActivationMode == "off" {
			iface.ActivationMode = ""
		}

		config.Network.Vlans[ifaceName] = iface
		if err := netplan.SaveConfig(config, temporaryNetplanFile); err != nil {
			fmt.Println("% error saving configuration:", err)
			return
		}
	},
}

func init() {
	startupCmd.AddCommand(ethernetStartupCmd)
	startupCmd.AddCommand(vlanStartupCmd)
}
